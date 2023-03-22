package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	cl "github.com/osmosis-labs/osmosis/v15/x/concentrated-liquidity"
	cltypes "github.com/osmosis-labs/osmosis/v15/x/concentrated-liquidity/types"
)

const (
	expectedPoolId     uint64 = 950
	addressPrefix             = "osmo"
	clientHomePath            = "/root/.osmosisd-local"
	consensusFee              = "1500uosmo"
	denom0                    = "uosmo"
	denom1                    = "uion"
	accountNamePrefix         = "lo-test"
	numPositions              = 1_000
	minAmountDeposited        = int64(1_000_000)
	randSeed                  = 1
	maxAmountDeposited        = 1_00_000_000
)

var (
	defaultAccountName = fmt.Sprintf("%s%d", accountNamePrefix, 1)
	exponentAtPriceOne = sdk.NewInt(-12)
	defaultMinAmount   = sdk.ZeroInt()
)

func main() {
	ctx := context.Background()

	// Create a Cosmos igniteClient instance
	igniteClient, err := cosmosclient.New(
		ctx,
		cosmosclient.WithAddressPrefix(addressPrefix),
		cosmosclient.WithKeyringBackend(cosmosaccount.KeyringTest),
		cosmosclient.WithHome(clientHomePath),
	)
	if err != nil {
		log.Fatal(err)
	}
	igniteClient.Factory = igniteClient.Factory.WithGas(300000).WithGasAdjustment(1.3).WithFees(consensusFee)

	statusResp, err := igniteClient.Status(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to: ", "chain-id", statusResp.NodeInfo.Network, "height", statusResp.SyncInfo.LatestBlockHeight)

	// pool current tick is set at 94
	// thus we populate positions in between -406 ~ 594 (across 1000 ticks)
	// positions would be
	// minTick := int64(-406)
	// maxTick := int64(594)

	// step2.
	minTick, maxTick := cl.GetMinAndMaxTicksFromExponentAtPriceOne(exponentAtPriceOne)
	log.Println(minTick, " ", maxTick)
	rand.Seed(randSeed)

	lowerTick := int64(-161999999100027)
	upperTick := minTick + (minTick+maxTick)/2
	var wg sync.WaitGroup
	for upperTick < maxTick {
		wg.Add(1)
		go func(lowerTick, upperTick int64) {
			defer wg.Done()
			var (
				// lowerTick = rand.Int63n(maxTick-minTick+1) + minTick
				// lowerTick <= upperTick <= maxTick
				// upperTick = maxTick - rand.Int63n(int64(math.Abs(float64(maxTick-lowerTick))))

				tokenDesired0 = sdk.NewCoin(denom0, sdk.NewInt(rand.Int63n(maxAmountDeposited)))
				tokenDesired1 = sdk.NewCoin(denom1, sdk.NewInt(rand.Int63n(maxAmountDeposited)))
			)

			// log every 1000 ticks
			// if lowerTick%1000 == 0 {
			// 	log.Println("creating tick of ", lowerTick)
			// }

			accountName := "my-key"
			log.Println("creating position: pool id", expectedPoolId, "accountName", accountName, "lowerTick", lowerTick, "upperTick", upperTick, "token0Desired", tokenDesired0, "tokenDesired1", tokenDesired1, "defaultMinAmount", defaultMinAmount)
			maxRetries := 20
			for i := 0; i < maxRetries; i++ {
				amt0, amt1, liquidity, err := createPosition(igniteClient, expectedPoolId, accountName, lowerTick, upperTick, tokenDesired0, tokenDesired1, defaultMinAmount, defaultMinAmount)
				if err == nil {
					break
				}
				log.Println("created position: amt0", amt0, "amt1", amt1, "liquidity", liquidity)
				time.Sleep(8 * time.Second)
			}
			// time.Sleep(time.Second * 8)
		}(lowerTick, upperTick)
		lowerTick++
		upperTick++

	}
	wg.Wait()
}

// func createPool(igniteClient cosmosclient.Client, accountName string) uint64 {
// 	msg := &model.MsgCreateConcentratedPool{
// 		Sender:                    getAccountAddressFromKeyring(igniteClient, accountName),
// 		Denom1:                    denom0,
// 		Denom0:                    denom1,
// 		TickSpacing:               1,
// 		PrecisionFactorAtPriceOne: exponentAtPriceOne,
// 		SwapFee:                   sdk.ZeroDec(),
// 	}
// 	txResp, err := igniteClient.BroadcastTx(accountName, msg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	resp := model.MsgCreateConcentratedPoolResponse{}
// 	if err := txResp.Decode(&resp); err != nil {
// 		log.Fatal(err)
// 	}
// 	return resp.PoolID
// }

func createPosition(client cosmosclient.Client, poolId uint64, senderKeyringAccountName string, lowerTick int64, upperTick int64, tokenDesired0, tokenDesired1 sdk.Coin, tokenMinAmount0, tokenMinAmount1 sdk.Int) (amountCreated0, amountCreated1 sdk.Int, liquidityCreated sdk.Dec, err error) {
	msg := &cltypes.MsgCreatePosition{
		PoolId:          poolId,
		Sender:          getAccountAddressFromKeyring(client, senderKeyringAccountName),
		LowerTick:       lowerTick,
		UpperTick:       upperTick,
		TokenDesired0:   tokenDesired0,
		TokenDesired1:   tokenDesired1,
		TokenMinAmount0: tokenMinAmount0,
		TokenMinAmount1: tokenMinAmount1,
	}
	txResp, err := client.BroadcastTx(senderKeyringAccountName, msg)
	if err != nil {
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, err
	}
	resp := cltypes.MsgCreatePositionResponse{}
	if err := txResp.Decode(&resp); err != nil {
		log.Fatal(err)
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, err
	}
	return resp.Amount0, resp.Amount1, resp.LiquidityCreated, nil
}

func getAccountAddressFromKeyring(igniteClient cosmosclient.Client, accountName string) string {
	account, err := igniteClient.Account(accountName)
	if err != nil {
		log.Fatal(fmt.Errorf("did not fimf account with name (%s) in the keyring: %w", accountName, err))
	}

	address := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}
	return address
}