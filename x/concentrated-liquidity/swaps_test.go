package concentrated_liquidity_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/osmosis-labs/osmosis/osmomath"
	"github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/internal/math"
	"github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/types"
	cltypes "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/types"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v14/x/poolmanager/types"
)

var _ = suite.TestingSuite(nil)

func (s *KeeperTestSuite) TestCalcAndSwapOutAmtGivenIn() {

	feeAdditiveTolerance := osmomath.ErrTolerance{
		// smallest dec * 10 = 10^-17
		AdditiveTolerance: sdk.SmallestDec().Mul(sdk.NewDec(10)),
		RoundingDir:       osmomath.RoundUp, // we want the actual fee to be higher the expected while withing tolerance.
	}

	tests := map[string]struct {
		positionAmount0                   sdk.Int
		positionAmount1                   sdk.Int
		addPositions                      func(ctx sdk.Context, poolId uint64)
		tokenIn                           sdk.Coin
		tokenOutDenom                     string
		priceLimit                        sdk.Dec
		expectedTokenIn                   sdk.Coin
		expectedTokenOut                  sdk.Coin
		expectedTick                      sdk.Int
		expectedSqrtPrice                 sdk.Dec
		newLowerPrice                     sdk.Dec
		newUpperPrice                     sdk.Dec
		poolLiqAmount0                    sdk.Int
		poolLiqAmount1                    sdk.Int
		swapFee                           sdk.Dec
		expectedFeeGrowthAccumulatorValue sdk.Dec
		expectErr                         bool
	}{
		//  One price range
		//
		//          5000
		//  4545 -----|----- 5500
		"single position within one tick: usdc -> eth (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(5004),
			swapFee:       sdk.ZeroDec(),
			// params
			// liquidity: 		 1517882343.751510418088349649
			// sqrtPriceNext:    70.738348247484497717 which is 5003.9139127823931095409 https://www.wolframalpha.com/input?i=70.710678118654752440+%2B+42000000+%2F+1517882343.751510418088349649
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  41999999.9999 rounded up https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2870.738349405152439867+-+70.710678118654752440%29
			// expectedTokenOut: 8396.71424216 rounded down https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2870.738348247484497717+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+70.738348247484497717%29
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(8396)),
			expectedTick:     sdk.NewInt(310040),
		},
		"fee 1 - single position within one tick: usdc -> eth (1% fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(5004),
			swapFee:       sdk.MustNewDecFromStr("0.01"),

			// params
			// liquidity: 		 1517882343.751510418088349649
			// sqrtPriceNext:    70.738071546196200264 which is 5003.9139127814610432508
			// expectedTokenIn:  41999999.9999 rounded up
			// expectedTokenOut: 8312
			// expectedFeeGrowthAccumulatorValue: 0.000276701288297452
			expectedTokenIn:                   sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			expectedTokenOut:                  sdk.NewCoin("eth", sdk.NewInt(8312)),
			expectedTick:                      sdk.NewInt(310039),
			expectedFeeGrowthAccumulatorValue: sdk.MustNewDecFromStr("0.000276701288297452"),
		},
		"single position within one tick: eth -> usdc (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4993),
			swapFee:       sdk.ZeroDec(),
			// params
			// liquidity: 		 1517882343.751510418088349649
			// sqrtPriceNext:    70.6666639108571443311 which is 4993.7773882900395488 https://www.wolframalpha.com/input?i=%28%281517882343.751510418088349649%29%29+%2F+%28%28%281517882343.751510418088349649%29+%2F+%2870.710678118654752440%29%29+%2B+%2813370%29%29
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  13370.00000 rounded up https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2870.710678118654752440+-+70.6666639108571443311+%29%29+%2F+%2870.6666639108571443311+*+70.710678118654752440%29
			// expectedTokenOut: 66808388.8901 rounded down https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2870.710678118654752440+-+70.6666639108571443311%29
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(13370)),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(66808388)),
			expectedTick:     sdk.NewInt(309938),
		},
		//  Two equal price ranges
		//
		//          5000
		//  4545 -----|----- 5500
		//  4545 -----|----- 5500
		"two positions within one tick: usdc -> eth (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// add second position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(5002),
			swapFee:       sdk.ZeroDec(),
			// // params are calculates by utilizing scripts from scripts/cl/main.py
			// liquidity: 		 3035764687.503020836176699298
			// sqrtPriceNext:    70.724513183069625078 which is 5001.956764982189191089 https://www.wolframalpha.com/input?i=70.710678118654752440%2B%2842000000+%2F+3035764687.503020836176699298%29
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  41999999.999 rounded up https://www.wolframalpha.com/input?i=3035764687.503020836176699298+*+%2870.724513183069625078+-+70.710678118654752440%29
			// expectedTokenOut: 8398.3567 rounded down https://www.wolframalpha.com/input?i=%283035764687.503020836176699298+*+%2870.724513183069625078+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+70.724513183069625078%29
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(8398)),
			expectedTick:     sdk.NewInt(310020),
			// two positions with same liquidity entered
			poolLiqAmount0: sdk.NewInt(1000000).MulRaw(2),
			poolLiqAmount1: sdk.NewInt(5000000000).MulRaw(2),
		},
		"two positions within one tick: eth -> usdc (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// add second position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4996),
			swapFee:       sdk.ZeroDec(),
			// params
			// liquidity: 		 3035764687.503020836176699298
			// sqrtPriceNext:    70.688664163408836319 which is 4996.88724120720067710 https://www.wolframalpha.com/input?i=%28%283035764687.503020836176699298%29%29+%2F+%28%28%283035764687.503020836176699298%29+%2F+%2870.710678118654752440%29%29+%2B+%2813370%29%29
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  13370.0000 rounded up https://www.wolframalpha.com/input?i=%283035764687.503020836176699298+*+%2870.710678118654752440+-+70.688664163408836319+%29%29+%2F+%2870.688664163408836319+*+70.710678118654752440%29
			// expectedTokenOut: 66829187.9678 rounded down https://www.wolframalpha.com/input?i=3035764687.503020836176699298+*+%2870.710678118654752440+-+70.688664163408836319%29
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(13370)),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(66829187)),
			expectedTick:     sdk.NewInt(309969),
			// two positions with same liquidity entered
			poolLiqAmount0: sdk.NewInt(1000000).MulRaw(2),
			poolLiqAmount1: sdk.NewInt(5000000000).MulRaw(2),
		},
		"fee 2 - two positions within one tick: eth -> usdc (3% fee) ": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// add second position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4996),
			swapFee:       sdk.MustNewDecFromStr("0.03"),
			// params
			// liquidity: 		 3035764687.503020836176699298
			// sqrtPriceNext:    70.689324382628080101 which is 4996.98058167241679801 https://www.wolframalpha.com/input?i=%28%283035764687.503020836176699298%29%29+%2F+%28%28%283035764687.503020836176699298%29+%2F+%2870.710678118654752440%29%29+%2B+%2813370+*+%281+-+0.03%29%29%29
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  13370.0000 rounded up https://www.wolframalpha.com/input?i=%283035764687.503020836176699298+*+%2870.710678118654752440+-+70.688664163408836319+%29%29+%2F+%2870.688664163408836319+*+70.710678118654752440%29
			// expectedTokenOut: 64824917.7760 rounded down https://www.wolframalpha.com/input?i=3035764687.503020836176699298+*+%2870.710678118654752440+-+70.689324382628080101%29
			// expectedFeeGrowthAccumulatorValue: 0.000000132124865162 https://www.wolframalpha.com/input?i=%2813370+*+0.03%29+%2F+3035764687.503020836176699298
			expectedTokenIn:                   sdk.NewCoin("eth", sdk.NewInt(13370)),
			expectedTokenOut:                  sdk.NewCoin("usdc", sdk.NewInt(64824917)),
			expectedTick:                      sdk.NewInt(309970),
			expectedFeeGrowthAccumulatorValue: sdk.MustNewDecFromStr("0.000000132124865162"),
			// two positions with same liquidity entered
			poolLiqAmount0: sdk.NewInt(1000000).MulRaw(2),
			poolLiqAmount1: sdk.NewInt(5000000000).MulRaw(2),
		},
		//  Consecutive price ranges
		//
		//          5000
		//  4545 -----|----- 5500
		//             5500 ----------- 6250
		//
		"two positions with consecutive price ranges: usdc -> eth (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
				// params
				// liquidity (1st):  1517882343.751510418088349649
				// sqrtPriceNext:    74.161984870956629487 which is 5500
				// sqrtPriceCurrent: 70.710678118654752440 which is 5000
				// expectedTokenIn:  5238677582.189386755771808942932776 rounded up https://www.wolframalpha.com/input?i=5.238677582189386755771808942932776425143606503+%C3%97+10%5E9&assumption=%22ClashPrefs%22+-%3E+%7B%22Math%22%7D
				// expectedTokenOut: 998976.6183474263883566299269 rounded down https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2874.161984870956629487+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+74.161984870956629487%29

				// create second position parameters
				newLowerPrice := sdk.NewDec(5500)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 315000
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(6250)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 322500
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
				// params
				// liquidity (2nd):  1197767444.955508123222985080
				// sqrtPriceNext:    78.137149196772377272 which is 6105.41408459866616274 https://www.wolframalpha.com/input?i=74.161984870956629487+%2B+4763454462.135+%2F+1197767444.955508123222985080
				// sqrtPriceCurrent: 74.161984870956629487 which is 5500
				// expectedTokenIn:  4761322417.810 rounded up https://www.wolframalpha.com/input?i=1197767444.955508123222985080+*+%2878.137149196772377272+-+74.161984870956629487%29
				// expectedTokenOut: 821653.452 rounded down https://www.wolframalpha.com/input?i=%281197767444.955508123222985080+*+%2878.137149196772377272+-+74.161984870956629487+%29%29+%2F+%2874.161984870956629487+*+78.137149196772377272%29
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(6106),
			swapFee:       sdk.ZeroDec(),
			// expectedTokenIn:  5238677582.189386755771808942932776 + 4761322417.810613244228191057067224 = 10000000000 usdc
			// expectedTokenOut: 998976.6183474263883566299269 + 821653.4522259 = 1820630.070 round down = 1.820630 eth
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(1820630)),
			expectedTick:     sdk.NewInt(321055),
			newLowerPrice:    sdk.NewDec(5500),
			newUpperPrice:    sdk.NewDec(6250),
		},
		//  Consecutive price ranges
		//
		//                     5000
		//             4545 -----|----- 5500
		//  4000 ----------- 4545
		//
		"two positions with consecutive price ranges: eth -> usdc (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
				// params
				// liquidity (1st):  1517882343.751510418088349649
				// sqrtPriceNext:    67.416615162732695594 which is 4545
				// sqrtPriceCurrent: 70.710678118654752440 which is 5000
				// expectedTokenIn:  1048861.292545921016650926872369 rounded up https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2870.710678118654752440+-+67.416615162732695594%29%29+%2F+%2867.416615162732695594+*+70.710678118654752440%29
				// expectedTokenOut: 5000000000.00000000000000 rounded down https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2870.710678118654752440-+67.416615162732695594%29

				// create second position parameters
				newLowerPrice := sdk.NewDec(4000)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 300000
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(4545)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 305450
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
				// params
				// liquidity (2nd):  1198735489.597250295669959398
				// sqrtPriceNext:    63.993486606491127478 which is 4095.1663280551593186 https://www.wolframalpha.com/input?i=%28%281198735489.597250295669959398%29%29+%2F+%28%28%281198735489.597250295669959398%29+%2F+%2867.416615162732695594%29%29+%2B+%28951138.707454078983349073127631%29%29
				// sqrtPriceCurrent: 67.416615162732695594 which is 4545
				// expectedTokenIn:  951138.707454078983338 rounded up https://www.wolframalpha.com/input?i=%281198735489.597250295669959398+*+%2867.416615162732695594+-+63.993486606491127478%29%29+%2F+%2863.993486606491127478+*+67.416615162732695594%29
				// expectedTokenOut: 4103425685.82056469999 rounded down https://www.wolframalpha.com/input?i=1198735489.597250295669959398+*+%2867.416615162732695594-+63.993486606491127478%29
				// expectedTick:     83179.3 rounded down https://www.wolframalpha.com/input?key=&i2d=true&i=Log%5B1.0001%2C4094.962290419%5D
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(2000000)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4094),
			swapFee:       sdk.ZeroDec(),
			// expectedTokenIn:  1048861.292545921016650926872369 + 951138.707454078983349073127631 = 2000000 eth
			// expectedTokenOut: 5000000000.000 + 4103425685.8205646999916265193598043375713541686 = 9103425685.8205646999916265193598043375713541686 round down = 9103.425685 usdc
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(2000000)),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(9103425685)),
			expectedTick:     sdk.NewInt(300952),
			newLowerPrice:    sdk.NewDec(4000),
			newUpperPrice:    sdk.NewDec(4545),
		},
		//  Consecutive price ranges
		//
		//                     5000
		//             4545 -----|----- 5500
		//  4000 ----------- 4545
		//
		// Ticks:
		// position   1:    305450, 315000,
		// posisition 2:    300000, 305450
		// current tick: 310000
		"fee 3 - two positions with consecutive price ranges: eth -> usdc (5% fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
				// params computed with sage scripts in scripts/cl/main.py
				// liquidity (1st):  1517882343.751510418088349649
				// sqrtPriceNext:    67.416615162732695594 which is 4545
				// sqrtPriceCurrent: 70.710678118654752440 which is 5000
				// expectedTokenIn:  1048861.292545921016650960
				// expectedTokenOut: 4999999999.99999999999999999970
				// expectedFeeGrowthAccumulatorValue: 0.000034550151296760

				// create second position parameters
				newLowerPrice := sdk.NewDec(4000)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 300000
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(4545)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 305450
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
				// params computed with sage scripts in scripts/cl/main.py
				// liquidity (2nd):  1198735489.597250295669959398
				// sqrtPriceNext:    64.3278909344373169748576312422 which is 4138.07755207286274968064829159
				// sqrtPriceCurrent: 67.416615162732695594 which is 4545
				// expectedTokenIn:  898695.642826782932516526784010 = 2000000 - 1048861.292545921016650960
				// expectedTokenOut: 3702563350.03654978405015422548
				// expectedFeeGrowthAccumulatorValue: 0.0000374851520884196734228699332666
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(2000000)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4094),
			swapFee:       sdk.MustNewDecFromStr("0.05"),
			// expectedTokenIn: 1101304.35717321706748347321599 + 898695.642826782932516526784010 = 2000000 eth
			// expectedTokenOut: 4999999999.99999999999999999970 + 3702563350.03654978405015422548 = 8702563350.03654978405015422518 round down = 8702.563350 usdc
			// expectedFeeGrowthAccumulatorValue:   0.000034550151296760 + 0.0000374851520884196734228699332666 = 0.0000720353033851796734228699332666
			expectedTokenIn:                   sdk.NewCoin("eth", sdk.NewInt(2000000)),
			expectedTokenOut:                  sdk.NewCoin("usdc", sdk.NewInt(8702563350)),
			expectedFeeGrowthAccumulatorValue: sdk.MustNewDecFromStr("0.000072035303385179"),
			expectedTick:                      sdk.NewInt(301381),
			newLowerPrice:                     sdk.NewDec(4000),
			newUpperPrice:                     sdk.NewDec(4545),
		},
		//  Partially overlapping price ranges

		//          5000
		//  4545 -----|----- 5500
		//        5001 ----------- 6250
		//
		// Ticks
		// position 1: 305450, 315000
		// position 2: 310010, 322500
		// current tick: 310000
		"two positions with partially overlapping price ranges: usdc -> eth (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
				// params
				// liquidity (1st):  1517882343.751510418088349649
				// sqrtPriceNext:    74.161984870956629487 which is 5500
				// sqrtPriceCurrent: 70.710678118654752440 which is 5000
				// expectedTokenIn:  5238677582.189386755771808942932776 rounded up https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2874.161984870956629487+-+70.710678118654752440%29
				// expectedTokenOut: 998976.6183474263883566299269692777 rounded down https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2874.161984870956629487+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+74.161984870956629487%29

				// create second position parameters
				newLowerPrice := sdk.NewDec(5001)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 310010
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(6250)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 322500
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
				// params
				// liquidity (2nd):  670416088.605668727039250938
				// sqrtPriceNext:    77.819789638253848946 which is 6055.9196593420811141 https://www.wolframalpha.com/input?i=70.717748832948578243+%2B+4761322417.810613244228191057067224+%2F+670416088.605668727039250938
				// sqrtPriceCurrent: 70.717748832948578243 which is 5001
				// expectedTokenIn:  4761322417.8106132444 rounded up https://www.wolframalpha.com/input?i=670416088.605668727039250938+*+%2877.819789638253848946+-+70.717748832948578243%29
				// expectedTokenOut: 865185.25913637514045 rounded down https://www.wolframalpha.com/input?i=%28670416088.605668727039250938+*+%2877.819789638253848946+-+70.717748832948578243+%29%29+%2F+%2870.717748832948578243+*+77.819789638253848946%29
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(6056),
			swapFee:       sdk.ZeroDec(),
			// expectedTokenIn:  5238677582.189386755771808942932776 + 4761322417.8106132444 = 10000000000.0000 = 10000.00 usdc
			// expectedTokenOut: 998976.6183474263883566299269692777 + 865185.2591363751404579873403641 = 1864161.877 round down = 1.864161 eth
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(1864161)),
			expectedTick:     sdk.NewInt(320560),
			newLowerPrice:    sdk.NewDec(5001),
			newUpperPrice:    sdk.NewDec(6250),
		},
		// Partially overlapping price ranges with fee

		//          5000
		//  4545 -----|----- 5500
		//        5001 ----------- 5843
		//
		// Ticks
		// position 1: 305450, 315000
		// position 2: 310010, 322500
		// current tick: 310000
		"fee 4 - two positions with partially overlapping price ranges: usdc -> eth (10% fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
				// params are calculates by utilizing scripts from scripts/cl/main.py
				// liquidity (1st):  1517882343.751510418088349649
				// sqrtPriceNext:    74.161984870956629487 which is 5500
				// sqrtPriceCurrent: 70.710678118654752440 which is 5000
				// expectedTokenIn:  5762545340.40832543134898983723
				// expectedTokenOut: 998976.618347426388356629926971

				// create second position parameters
				newLowerPrice := sdk.NewDec(5001)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 310010
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(6250)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 322500
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)

				pool, err := s.App.ConcentratedLiquidityKeeper.GetPoolById(ctx, 1)
				fmt.Println(pool.GetCurrentSqrtPrice().Mul(pool.GetCurrentSqrtPrice()).String())

				// params
				// liquidity (2nd):  670416088.605668727039250938
				// sqrtPriceNext:    76.4063193467682254976579845167 which is 5837.925636120329
				// sqrtPriceCurrent: 70.717748832948578243 which is 5001
				// expectedTokenIn:  4237454659.59167456865101016277 = 10000000000 - 5762545340.40832543134898983723
				// expectedTokenOut: 705813.347855134472186382130036

				//////////////////////////////////////////////////
				// 1. Only position 1
				// * tick 310000 to 310010
				// * price range: 5000 to 5001
				// * sqrt price: 70.710678118654752440 to 70.717748832948578243

				// liquidity:  1517882343.751510418088349649 (1st)

				// expectedTokenIn (no fee): 10732512.384309615746632728158
				// expectedTokenOut (with fee): 11805763.622740577321296000974
				// expectedTokenOut: 2146.28785880640879265591374059
				// expectedFeeGrowthAccumulatorValue: 0.000707071429382580300000000000073

				// expectedRemainingTokenIn = 10000000000 - 11805763.622740577321296000974 = 9988194236.37725942267870399903

				//////////////////////////////////////////////////
				// 2. Both position 1 and 2
				// * tick 310000 to 315000
				// * price range: 5001 to 5500
				// * sqrt price: 70.717748832948578243 to 74.161984870956629487

				// liquidity: 1517882343.751510418088349649 (1st) + 670416088.605668727039250938 (2nd) = 2188298432.35717914512760058700

				// expectedTokenIn (no fee): 7537016322.64112022429423919467
				// expectedTokenIn (with fee): 8290717954.9052322467236631141
				// expectedTokenOut: 1437108.91592757237716789250871
				// expectedFeeGrowthAccumulatorValue: 0.344423603800805124400000000000

				// expectedRemainingTokenIn = 9988194236.37725942267870399903 - 8290717954.9052322467236631141 = 1697476281.47202717595504088493

				//////////////////////////////////////////////////
				// 3. Only position 2
				// * tick 315000 to 322500
				// * price range: 5500 to 5843
				// * sqrt price: 74.161984870956629487 to 76.4422024931482315166509926684

				// liquidity: 670416088.605668727039250938 (2nd)

				// remaining token in (with fee): 1697476281.47202717595504088493
				// expectedTokenOut: 269488.274305469529889078712213
				// expectedFeeGrowthAccumulatorValue: 0.253197426243519613677553835191

			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(6056),
			swapFee:       sdk.MustNewDecFromStr("0.1"),
			// expectedTokenIn:  5762545340.40832543134898983723 + 4237454659.59167456865101016277 = 10000000000.0000 = 10000.00 usdc
			// expectedTokenOut: 2146.28785880640879265591374059 + "1437108.91592757237716789250871 + 269488.274305469529889078712213 = 1708743.47809184831584962713466 eth
			// expectedFeeGrowthAccumulatorValue: 0.000707071429382580300000000000073 + 0.344423603800805124400000000000 + 0.253197426243519613677553835191 = 0.598328101473707318377553835191
			expectedTokenIn:                   sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			expectedTokenOut:                  sdk.NewCoin("eth", sdk.NewInt(1708743)),
			expectedFeeGrowthAccumulatorValue: sdk.MustNewDecFromStr("0.598328101473707318"),
			expectedTick:                      sdk.NewInt(318432),
			newLowerPrice:                     sdk.NewDec(5001),
			newUpperPrice:                     sdk.NewDec(6250),
		},
		"two positions with partially overlapping price ranges, not utilizing full liquidity of second position: usdc -> eth (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
				// params
				// liquidity (1st):  1517882343.751510418088349649
				// sqrtPriceNext:    74.161984870956629487 which is 5500
				// sqrtPriceCurrent: 70.710678118654752440 which is 5000
				// expectedTokenIn:  5238677582.189386755771808942932776 rounded up https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2874.161984870956629487+-+70.710678118654752440%29
				// expectedTokenOut: 998976.61834742638835662992696 rounded down https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2874.161984870956629487+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+74.161984870956629487%29

				// create second position parameters
				newLowerPrice := sdk.NewDec(5001)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 310010
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(6250)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 322500
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
				// params
				// liquidity (2nd):  670416088.605668727039250938
				// sqrtPriceNext:    75.582373165866231044 which is 5712.695133384 https://www.wolframalpha.com/input?i=70.717748832948578243+%2B+3261322417.810613244228191057067224+%2F+670416088.605668727039250938
				// sqrtPriceCurrent: 70.717748832948578243 which is 5001
				// expectedTokenIn:  3261322417.8106132442 rounded up https://www.wolframalpha.com/input?i=670416088.605668727039250938+*+%2875.582373165866231044+-+70.717748832948578243%29
				// expectedTokenOut: 610161.47679708043791 rounded down https://www.wolframalpha.com/input?i=%28670416088.605668727039250938+*+%2875.582373165866231044+-+70.717748832948578243+%29%29+%2F+%2870.717748832948578243+*+75.582373165866231044%29
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(8500000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(6056),
			swapFee:       sdk.ZeroDec(),
			// expectedTokenIn:  5238677582.189386755771808942932776 + 3261322417.810613244228191057067224 = 8500000000.000 = 8500.00 usdc
			// expectedTokenOut: 998976.61834742638835662992696 + 610161.47679708043791 = 1609138.09 round down = 1.609138 eth
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(8500000000)),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(1609138)),
			expectedTick:     sdk.NewInt(317127),
			newLowerPrice:    sdk.NewDec(5001),
			newUpperPrice:    sdk.NewDec(6250),
		},
		//  Partially overlapping price ranges
		//
		//                5000
		//        4545 -----|----- 5500
		//  4000 ----------- 4999
		//
		"two positions with partially overlapping price ranges: eth -> usdc (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
				// params
				// liquidity (1st):  1517882343.751510418088349649
				// sqrtPriceNext:    67.416615162732695594 which is 4545
				// sqrtPriceCurrent: 70.710678118654752440 which is 5000
				// expectedTokenIn:  1048861.292545921016650926872369076 rounded up https://www.wolframalpha.com/input?key=&i=%281517882343.751510418088349649+*+%2870.710678118654752440+-+67.416615162732695594%29%29+%2F+%2867.416615162732695594+*+70.710678118654752440%29
				// expectedTokenOut: 5000000000.000 rounded down https://www.wolframalpha.com/input?key=&i=1517882343.751510418088349649+*+%2870.710678118654752440-+67.416615162732695594%29

				// create second position parameters
				newLowerPrice := sdk.NewDec(4000)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 300000
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(4999)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 309990
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
				// params
				// liquidity (2nd):  670416215.718827443660400594
				// sqrtPriceNext:    64.257941776684699569 which is 4129.083081375800804213 https://www.wolframalpha.com/input?i=%28%28670416215.718827443660400594%29%29+%2F+%28%28%28670416215.718827443660400594%29+%2F+%2870.703606697254136612%29%29+%2B+%28951138.707454078983349%29%29
				// sqrtPriceCurrent: 70.703606697254136612 which is 4999.00
				// expectedTokenIn:  951138.70745407898329 rounded up https://www.wolframalpha.com/input?i=%28670416215.718827443660400594+*+%2870.703606697254136612+-+64.257941776684699569%29%29+%2F+%2864.257941776684699569+*+70.703606697254136612%29
				// expectedTokenOut: 4321278283.8397584645 rounded down https://www.wolframalpha.com/input?i=670416215.718827443660400594+*+%2870.703606697254136612-+64.257941776684699569%29
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(2000000)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4128),
			swapFee:       sdk.ZeroDec(),
			// expectedTokenIn:  1048861.292545921016650926872369076 + 951138.70745407898329 = 2000000 eth
			// expectedTokenOut: 5000000000.000 + 4321278283.839758464593299720838190090442803542 = 9321278283.8397584645932997208 round down = 9321.278283 usdc
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(2000000)),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(9321278283)),
			expectedTick:     sdk.NewInt(301291),
			newLowerPrice:    sdk.NewDec(4000),
			newUpperPrice:    sdk.NewDec(4999),
		},
		"two positions with partially overlapping price ranges, not utilizing full liquidity of second position: eth -> usdc (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
				// params
				// liquidity (1st):  1517882343.751510418088349649
				// sqrtPriceNext:    67.416615162732695594 which is 4545
				// sqrtPriceCurrent: 70.710678118654752440 which is 5000
				// expectedTokenIn:  1048861.292545921016650926872369076 rounded up https://www.wolframalpha.com/input?key=&i=%281517882343.751510418088349649+*+%2870.710678118654752440+-+67.416615162732695594%29%29+%2F+%2867.416615162732695594+*+70.710678118654752440%29
				// expectedTokenOut: 5000000000.000 rounded down https://www.wolframalpha.com/input?key=&i=1517882343.751510418088349649+*+%2870.710678118654752440-+67.416615162732695594%29

				// create second position parameters
				newLowerPrice := sdk.NewDec(4000)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 300000
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(4999)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 309990
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
				// params
				// liquidity (2nd):  670416215.718827443660400594
				// sqrtPriceNext:    65.513813187509027302 which is 4292.059718367831736 https://www.wolframalpha.com/input?i=%28%28670416215.718827443660400594%29%29+%2F+%28%28%28670416215.718827443660400594%29+%2F+%2870.703606697254136612%29%29+%2B+%28751138.70745407898334907%29%29
				// sqrtPriceCurrent: 70.703606697254136612 which is 4999.00
				// expectedTokenIn:  751138.70745407898 rounded up https://www.wolframalpha.com/input?key=&i=%28670416215.718827443660400594+*+%2870.703606697254136612+-+65.513813187509027302%29%29+%2F+%2865.513813187509027302+*+70.703606697254136612%29
				// expectedTokenOut: 3479321725.1654478001 rounded down https://www.wolframalpha.com/input?key=&i=670416215.718827443660400594+*+%2870.703606697254136612-+65.513813187509027302%29
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(1800000)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4128),
			swapFee:       sdk.ZeroDec(),
			// expectedTokenIn:  1048861.292545921016650926872369076 + 751138.70745407898334907 = 1.800000 eth
			// expectedTokenOut: 5000000000.000 + 3479321725.1654478001068768736 = 8479321725.1654478001068768736 round down = 8479.321725 usdc
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(1800000)),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(8479321725)),
			expectedTick:     sdk.NewInt(302921),
			newLowerPrice:    sdk.NewDec(4000),
			newUpperPrice:    sdk.NewDec(4999),
		},
		//  Sequential price ranges with a gap
		//
		//          5000
		//  4545 -----|----- 5500
		//              5501 ----------- 6250
		//
		"two sequential positions with a gap (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
				// params
				// liquidity (1st):  1517882343.751510418088349649
				// sqrtPriceNext:    74.161984870956629487 which is 5500
				// sqrtPriceCurrent: 70.710678118654752440 which is 5000
				// expectedTokenIn:  5238677582.1893867557718089429327 rounded up https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2874.161984870956629487+-+70.710678118654752440%29
				// expectedTokenOut: 998976.61834742638835 rounded down https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2874.161984870956629487+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+74.161984870956629487%29

				// create second position parameters
				newLowerPrice := sdk.NewDec(5501)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 315010
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(6250)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 322500
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
				// params
				// liquidity (2nd):  1199528406.187413669220037261
				// sqrtPriceNext:    78.138055170339538272 which is 6105.5556658030254493528 https://www.wolframalpha.com/input?i=74.168726563154635303++%2B++4761322417.8106132442281910570673+%2F+1199528406.187413669220037261
				// sqrtPriceCurrent: 74.168726563154635303 which is 5501
				// expectedTokenIn:  4761322417.810613244281820035563194 rounded up https://www.wolframalpha.com/input?i=1199528406.187413669220037261+*+%2878.138055170339538272+-+74.168726563154635303%29
				// expectedTokenOut: 821569.240826953837970 rounded down https://www.wolframalpha.com/input?i=%281199528406.187413669220037261+*+%2878.138055170339538272+-+74.168726563154635303+%29%29+%2F+%2874.168726563154635303+*+78.138055170339538272%29
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(6106),
			swapFee:       sdk.ZeroDec(),
			// expectedTokenIn:  5238677582.1893867557718089429327 + 4761322417.810613244281820035563194 = 10000000000 usdc
			// expectedTokenOut: 998976.61834742638835 + 821569.240826953837970 = 1820545.85917438022632 round down = 1.820545 eth
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(1820545)),
			expectedTick:     sdk.NewInt(321056),
			newLowerPrice:    sdk.NewDec(5501),
			newUpperPrice:    sdk.NewDec(6250),
		},
		// Slippage protection doesn't cause a failure but interrupts early.
		"single position within one tick, trade completes but slippage protection interrupts trade early: eth -> usdc (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4994),
			swapFee:       sdk.ZeroDec(),
			// params
			// liquidity: 		 1517882343.751510418088349649
			// sqrtPriceNext:    70.668238976219012614 which is 4994 https://www.wolframalpha.com/input?i=70.710678118654752440+%2B+42000000+%2F+1517882343.751510418088349649
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  12891.26207649936510 rounded up https://www.wolframalpha.com/input?key=&i=%281517882343.751510418088349649+*+%2870.710678118654752440+-+70.668238976219012614+%29%29+%2F+%2870.710678118654752440+*+70.668238976219012614%29
			// expectedTokenOut: 64417624.98716495170 rounded down https://www.wolframalpha.com/input?key=&i=1517882343.751510418088349649+*+%2870.710678118654752440+-+70.668238976219012614%29
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(12891)),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(64417624)),
			expectedTick:     sdk.NewInt(309941),
		},
		"single position within one tick, trade does not complete due to lack of liquidity: usdc -> eth (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(5300000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(6000),
			swapFee:       sdk.ZeroDec(),
			expectErr:     true,
		},
		"single position within one tick, trade does not complete due to lack of liquidity: eth -> usdc (zero fee)": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(1100000)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4000),
			swapFee:       sdk.ZeroDec(),
			expectErr:     true,
		},
	}

	for name, test := range tests {
		test := test
		s.Run(name, func() {
			s.Setup()
			s.FundAcc(s.TestAccs[0], sdk.NewCoins(sdk.NewCoin("eth", sdk.NewInt(10000000000000)), sdk.NewCoin("usdc", sdk.NewInt(1000000000000))))
			s.FundAcc(s.TestAccs[1], sdk.NewCoins(sdk.NewCoin("eth", sdk.NewInt(10000000000000)), sdk.NewCoin("usdc", sdk.NewInt(1000000000000))))

			// Create default CL pool
			pool := s.PrepareCustomConcentratedPool(s.TestAccs[0], ETH, USDC, DefaultTickSpacing, DefaultExponentAtPriceOne, sdk.ZeroDec())

			// add positions
			test.addPositions(s.Ctx, pool.GetId())

			poolBeforeCalc, err := s.App.ConcentratedLiquidityKeeper.GetPoolById(s.Ctx, pool.GetId())
			s.Require().NoError(err)

			// perform calc
			_, tokenIn, tokenOut, updatedTick, updatedLiquidity, _, err := s.App.ConcentratedLiquidityKeeper.CalcOutAmtGivenInInternal(
				s.Ctx,
				test.tokenIn, test.tokenOutDenom,
				test.swapFee, test.priceLimit, pool.GetId())
			if test.expectErr {
				s.Require().Error(err)
			} else {
				// writeCtx()

				s.Require().NoError(err)

				// check that tokenIn, tokenOut, tick, and sqrtPrice from CalcOut are all what we expected
				s.Require().Equal(test.expectedTick.String(), updatedTick.String())
				s.Require().Equal(test.expectedTokenIn.String(), tokenIn.String())
				s.Require().Equal(test.expectedTokenOut.String(), tokenOut.String())

				if test.newLowerPrice.IsNil() && test.newUpperPrice.IsNil() {
					test.newLowerPrice = DefaultLowerPrice
					test.newUpperPrice = DefaultUpperPrice
				}

				newLowerTick, err := math.PriceToTick(test.newLowerPrice, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)
				newUpperTick, err := math.PriceToTick(test.newUpperPrice, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)

				lowerSqrtPrice, err := math.TickToSqrtPrice(newLowerTick, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)
				upperSqrtPrice, err := math.TickToSqrtPrice(newUpperTick, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)

				if test.poolLiqAmount0.IsNil() && test.poolLiqAmount1.IsNil() {
					test.poolLiqAmount0 = DefaultAmt0
					test.poolLiqAmount1 = DefaultAmt1
				}

				// check that liquidity is what we expected
				expectedLiquidity := math.GetLiquidityFromAmounts(DefaultCurrSqrtPrice, lowerSqrtPrice, upperSqrtPrice, test.poolLiqAmount0, test.poolLiqAmount1)
				s.Require().Equal(expectedLiquidity.String(), updatedLiquidity.String())

				// check that the pool has not been modified after performing calc
				poolAfterCalc, err := s.App.ConcentratedLiquidityKeeper.GetPoolById(s.Ctx, pool.GetId())
				s.Require().NoError(err)

				s.Require().Equal(poolBeforeCalc.GetCurrentSqrtPrice(), poolAfterCalc.GetCurrentSqrtPrice())
				s.Require().Equal(poolBeforeCalc.GetCurrentTick(), poolAfterCalc.GetCurrentTick())
				s.Require().Equal(poolBeforeCalc.GetTotalShares(), poolAfterCalc.GetTotalShares())
				s.Require().Equal(poolBeforeCalc.GetLiquidity(), poolAfterCalc.GetLiquidity())
				s.Require().Equal(poolBeforeCalc.GetTickSpacing(), poolAfterCalc.GetTickSpacing())
			}

			// perform swap
			// TODO: Add sqrtPrice check
			tokenIn, tokenOut, updatedTick, updatedLiquidity, _, err = s.App.ConcentratedLiquidityKeeper.SwapOutAmtGivenIn(
				s.Ctx,
				test.tokenIn, test.tokenOutDenom,
				test.swapFee, test.priceLimit, pool.GetId())
			if test.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				s.Require().Equal(test.expectedTokenIn.String(), tokenIn.String())
				s.Require().Equal(test.expectedTokenOut.String(), tokenOut.String())
				s.Require().Equal(test.expectedTick.String(), updatedTick.String())

				if test.newLowerPrice.IsNil() && test.newUpperPrice.IsNil() {
					test.newLowerPrice = DefaultLowerPrice
					test.newUpperPrice = DefaultUpperPrice
				}

				newLowerTick, err := math.PriceToTick(test.newLowerPrice, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)
				newUpperTick, err := math.PriceToTick(test.newUpperPrice, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)

				lowerSqrtPrice, err := math.TickToSqrtPrice(newLowerTick, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)
				upperSqrtPrice, err := math.TickToSqrtPrice(newUpperTick, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)

				if test.poolLiqAmount0.IsNil() && test.poolLiqAmount1.IsNil() {
					test.poolLiqAmount0 = DefaultAmt0
					test.poolLiqAmount1 = DefaultAmt1
				}

				expectedLiquidity := math.GetLiquidityFromAmounts(DefaultCurrSqrtPrice, lowerSqrtPrice, upperSqrtPrice, test.poolLiqAmount0, test.poolLiqAmount1)
				s.Require().Equal(expectedLiquidity.String(), updatedLiquidity.String())

				// Validate fee accumulator updates
				feeGrowthAccumulator, err := s.App.ConcentratedLiquidityKeeper.GetFeeAccumulator(s.Ctx, pool.GetId())
				s.Require().NoError(err)
				feeGrowthAccumulatorValue := feeGrowthAccumulator.GetValue().AmountOf(test.tokenIn.Denom)
				if test.swapFee.IsZero() {
					s.Require().Equal(feeGrowthAccumulatorValue, sdk.ZeroDec())
				} else {
					// s.Require().Equal(test.expectedFeeGrowthAccumulatorValue, feeGrowthAccumulatorValue)
					// We must not undercharge the fee growth accumulator, but we can overcharge it
					// by a small amount. This assert validates this.
					// The epsilon is 10^-17.
					isExpectedLTEActual := feeAdditiveTolerance.CompareBigDec(
						osmomath.BigDecFromSDKDec(test.expectedFeeGrowthAccumulatorValue),
						osmomath.BigDecFromSDKDec(feeGrowthAccumulatorValue)) == 0

					s.Require().True(isExpectedLTEActual, "expected (%s) <= actual (%s)", test.expectedFeeGrowthAccumulatorValue, feeGrowthAccumulatorValue)
				}
			}
		})

	}
}

func (s *KeeperTestSuite) TestCalcAndSwapInAmtGivenOut() {
	tests := map[string]struct {
		positionAmount0   sdk.Int
		positionAmount1   sdk.Int
		addPositions      func(ctx sdk.Context, poolId uint64)
		tokenOut          sdk.Coin
		tokenInDenom      string
		priceLimit        sdk.Dec
		expectedTokenIn   sdk.Coin
		expectedTokenOut  sdk.Coin
		expectedTick      sdk.Int
		expectedSqrtPrice sdk.Dec
		newLowerPrice     sdk.Dec
		newUpperPrice     sdk.Dec
		poolLiqAmount0    sdk.Int
		poolLiqAmount1    sdk.Int
		expectErr         bool
	}{
		//  One price range
		//
		//          5000
		//  4545 -----|----- 5500
		"single position within one tick: usdc -> eth": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			tokenInDenom:     "eth",
			priceLimit:       sdk.NewDec(5004),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(8396)),
			expectedTick:     sdk.NewInt(310040),
		},
		"single position within one tick: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenInDenom:     "usdc",
			priceLimit:       sdk.NewDec(4993),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(13370)),
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(66808388)),
			expectedTick:     sdk.NewInt(309938),
		},
		//  Two equal price ranges
		//
		//          5000
		//  4545 -----|----- 5500
		//  4545 -----|----- 5500
		"two positions within one tick: usdc -> eth": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// add second position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			tokenInDenom:     "eth",
			priceLimit:       sdk.NewDec(5002),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(8398)),
			expectedTick:     sdk.NewInt(310020),
			// two positions with same liquidity entered
			poolLiqAmount0: sdk.NewInt(1000000).MulRaw(2),
			poolLiqAmount1: sdk.NewInt(5000000000).MulRaw(2),
		},
		"two positions within one tick: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// add second position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenInDenom:     "usdc",
			priceLimit:       sdk.NewDec(4996),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(13370)),
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(66829187)),
			expectedTick:     sdk.NewInt(309969),
			// two positions with same liquidity entered
			poolLiqAmount0: sdk.NewInt(1000000).MulRaw(2),
			poolLiqAmount1: sdk.NewInt(5000000000).MulRaw(2),
		},
		//  Consecutive price ranges
		//
		//          5000
		//  4545 -----|----- 5500
		//             5500 ----------- 6250
		//
		"two positions with consecutive price ranges: usdc -> eth": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// create second position parameters
				newLowerPrice := sdk.NewDec(5500)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 315000
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(6250)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 322500
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			tokenInDenom:     "eth",
			priceLimit:       sdk.NewDec(6106),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(1820630)),
			expectedTick:     sdk.NewInt(321055),
			newLowerPrice:    sdk.NewDec(5500),
			newUpperPrice:    sdk.NewDec(6250),
		},
		//  Consecutive price ranges
		//
		//                     5000
		//             4545 -----|----- 5500
		//  4000 ----------- 4545
		//
		"two positions with consecutive price ranges: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// create second position parameters
				newLowerPrice := sdk.NewDec(4000)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 300000
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(4545)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 305450
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("eth", sdk.NewInt(2000000)),
			tokenInDenom:     "usdc",
			priceLimit:       sdk.NewDec(4094),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(2000000)),
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(9103425685)),
			expectedTick:     sdk.NewInt(300952),
			newLowerPrice:    sdk.NewDec(4000),
			newUpperPrice:    sdk.NewDec(4545),
		},
		//  Partially overlapping price ranges
		//
		//          5000
		//  4545 -----|----- 5500
		//        5001 ----------- 6250
		//
		"two positions with partially overlapping price ranges: usdc -> eth": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// create second position parameters
				newLowerPrice := sdk.NewDec(5001)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 310010
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(6250)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 322500
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			tokenInDenom:     "eth",
			priceLimit:       sdk.NewDec(6056),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(1864161)),
			expectedTick:     sdk.NewInt(320560),
			newLowerPrice:    sdk.NewDec(5001),
			newUpperPrice:    sdk.NewDec(6250),
		},
		"two positions with partially overlapping price ranges, not utilizing full liquidity of second position: usdc -> eth": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// create second position parameters
				newLowerPrice := sdk.NewDec(5001)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 310010
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(6250)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 322500
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("usdc", sdk.NewInt(8500000000)),
			tokenInDenom:     "eth",
			priceLimit:       sdk.NewDec(6056),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(8500000000)),
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(1609138)),
			expectedTick:     sdk.NewInt(317127),
			newLowerPrice:    sdk.NewDec(5001),
			newUpperPrice:    sdk.NewDec(6250),
		},
		//  Partially overlapping price ranges
		//
		//                5000
		//        4545 -----|----- 5500
		//  4000 ----------- 4999
		//
		"two positions with partially overlapping price ranges: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// create second position parameters
				newLowerPrice := sdk.NewDec(4000)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 300000
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(4999)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 309990
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("eth", sdk.NewInt(2000000)),
			tokenInDenom:     "usdc",
			priceLimit:       sdk.NewDec(4128),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(2000000)),
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(9321278283)),
			expectedTick:     sdk.NewInt(301291),
			newLowerPrice:    sdk.NewDec(4000),
			newUpperPrice:    sdk.NewDec(4999),
		},
		"two positions with partially overlapping price ranges, not utilizing full liquidity of second position: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// create second position parameters
				newLowerPrice := sdk.NewDec(4000)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 300000
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(4999)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 309990
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("eth", sdk.NewInt(1800000)),
			tokenInDenom:     "usdc",
			priceLimit:       sdk.NewDec(4128),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(1800000)),
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(8479321725)),
			expectedTick:     sdk.NewInt(302921),
			newLowerPrice:    sdk.NewDec(4000),
			newUpperPrice:    sdk.NewDec(4999),
		},
		//  Sequential price ranges with a gap
		//
		//          5000
		//  4545 -----|----- 5500
		//              5501 ----------- 6250
		//
		"two sequential positions with a gap": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)

				// create second position parameters
				newLowerPrice := sdk.NewDec(5501)
				newLowerTick, err := math.PriceToTick(newLowerPrice, DefaultExponentAtPriceOne) // 315010
				s.Require().NoError(err)
				newUpperPrice := sdk.NewDec(6250)
				newUpperTick, err := math.PriceToTick(newUpperPrice, DefaultExponentAtPriceOne) // 322500
				s.Require().NoError(err)

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			tokenInDenom:     "eth",
			priceLimit:       sdk.NewDec(6106),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(1820545)),
			expectedTick:     sdk.NewInt(321056),
			newLowerPrice:    sdk.NewDec(5501),
			newUpperPrice:    sdk.NewDec(6250),
		},
		// Slippage protection doesn't cause a failure but interrupts early.
		"single position within one tick, trade completes but slippage protection interrupts trade early: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenOut:         sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenInDenom:     "usdc",
			priceLimit:       sdk.NewDec(4994),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(12891)),
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(64417624)),
			expectedTick:     sdk.NewInt(309941),
		},
		"single position within one tick, trade does not complete due to lack of liquidity: usdc -> eth": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenOut:     sdk.NewCoin("usdc", sdk.NewInt(5300000000)),
			tokenInDenom: "eth",
			priceLimit:   sdk.NewDec(6000),
			expectErr:    true,
		},
		"single position within one tick, trade does not complete due to lack of liquidity: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], DefaultAmt0, DefaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), DefaultLowerTick, DefaultUpperTick)
				s.Require().NoError(err)
			},
			tokenOut:     sdk.NewCoin("eth", sdk.NewInt(1100000)),
			tokenInDenom: "usdc",
			priceLimit:   sdk.NewDec(4000),
			expectErr:    true,
		},
	}

	for name, test := range tests {
		s.Run(name, func() {
			s.Setup()
			s.FundAcc(s.TestAccs[0], sdk.NewCoins(sdk.NewCoin("eth", sdk.NewInt(10000000000000)), sdk.NewCoin("usdc", sdk.NewInt(1000000000000))))
			s.FundAcc(s.TestAccs[1], sdk.NewCoins(sdk.NewCoin("eth", sdk.NewInt(10000000000000)), sdk.NewCoin("usdc", sdk.NewInt(1000000000000))))

			// Create default CL pool
			pool := s.PrepareConcentratedPool()

			// add positions
			test.addPositions(s.Ctx, pool.GetId())

			poolBeforeCalc, err := s.App.ConcentratedLiquidityKeeper.GetPoolById(s.Ctx, pool.GetId())
			s.Require().NoError(err)

			// perform calc
			_, tokenIn, tokenOut, updatedTick, updatedLiquidity, _, err := s.App.ConcentratedLiquidityKeeper.CalcInAmtGivenOutInternal(
				s.Ctx,
				test.tokenOut, test.tokenInDenom,
				DefaultZeroSwapFee, test.priceLimit, pool.GetId())
			if test.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// check that tokenIn, tokenOut, tick, and sqrtPrice from CalcOut are all what we expected
				s.Require().Equal(test.expectedTokenOut.String(), tokenOut.String())
				s.Require().Equal(test.expectedTokenIn.String(), tokenIn.String())
				s.Require().Equal(test.expectedTick.String(), updatedTick.String())

				if test.newLowerPrice.IsNil() && test.newUpperPrice.IsNil() {
					test.newLowerPrice = DefaultLowerPrice
					test.newUpperPrice = DefaultUpperPrice
				}

				newLowerTick, err := math.PriceToTick(test.newLowerPrice, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)
				newUpperTick, err := math.PriceToTick(test.newUpperPrice, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)

				lowerSqrtPrice, err := math.TickToSqrtPrice(newLowerTick, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)
				upperSqrtPrice, err := math.TickToSqrtPrice(newUpperTick, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)

				if test.poolLiqAmount0.IsNil() && test.poolLiqAmount1.IsNil() {
					test.poolLiqAmount0 = DefaultAmt0
					test.poolLiqAmount1 = DefaultAmt1
				}

				// check that liquidity is what we expected
				expectedLiquidity := math.GetLiquidityFromAmounts(DefaultCurrSqrtPrice, lowerSqrtPrice, upperSqrtPrice, test.poolLiqAmount0, test.poolLiqAmount1)
				s.Require().Equal(expectedLiquidity.String(), updatedLiquidity.String())

				// check that the pool has not been modified after performing calc
				poolAfterCalc, err := s.App.ConcentratedLiquidityKeeper.GetPoolById(s.Ctx, pool.GetId())
				s.Require().NoError(err)

				s.Require().Equal(poolBeforeCalc.GetCurrentSqrtPrice(), poolAfterCalc.GetCurrentSqrtPrice())
				s.Require().Equal(poolBeforeCalc.GetCurrentTick(), poolAfterCalc.GetCurrentTick())
				s.Require().Equal(poolBeforeCalc.GetTotalShares(), poolAfterCalc.GetTotalShares())
				s.Require().Equal(poolBeforeCalc.GetLiquidity(), poolAfterCalc.GetLiquidity())
				s.Require().Equal(poolBeforeCalc.GetTickSpacing(), poolAfterCalc.GetTickSpacing())
			}

			// perform swap
			// TODO: Add sqrtPrice check
			tokenIn, tokenOut, updatedTick, updatedLiquidity, _, err = s.App.ConcentratedLiquidityKeeper.SwapInAmtGivenOut(
				s.Ctx,
				test.tokenOut, test.tokenInDenom,
				DefaultZeroSwapFee, test.priceLimit, pool.GetId())
			if test.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				pool, err = s.App.ConcentratedLiquidityKeeper.GetPoolById(s.Ctx, pool.GetId())
				s.Require().NoError(err)

				// check that tokenIn, tokenOut, tick, and sqrtPrice from SwapOut are all what we expected
				s.Require().Equal(test.expectedTick.String(), updatedTick.String())
				s.Require().Equal(test.expectedTokenIn.String(), tokenIn.String())
				s.Require().Equal(test.expectedTokenOut.String(), tokenOut.String())
				// also ensure the pool's currentTick and currentSqrtPrice was updated due to calling a mutative method
				s.Require().Equal(test.expectedTick.String(), pool.GetCurrentTick().String())

				if test.newLowerPrice.IsNil() && test.newUpperPrice.IsNil() {
					test.newLowerPrice = DefaultLowerPrice
					test.newUpperPrice = DefaultUpperPrice
				}

				newLowerTick, err := math.PriceToTick(test.newLowerPrice, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)
				newUpperTick, err := math.PriceToTick(test.newUpperPrice, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)

				lowerSqrtPrice, err := math.TickToSqrtPrice(newLowerTick, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)
				upperSqrtPrice, err := math.TickToSqrtPrice(newUpperTick, pool.GetPrecisionFactorAtPriceOne())
				s.Require().NoError(err)

				if test.poolLiqAmount0.IsNil() && test.poolLiqAmount1.IsNil() {
					test.poolLiqAmount0 = DefaultAmt0
					test.poolLiqAmount1 = DefaultAmt1
				}

				expectedLiquidity := math.GetLiquidityFromAmounts(DefaultCurrSqrtPrice, lowerSqrtPrice, upperSqrtPrice, test.poolLiqAmount0, test.poolLiqAmount1)
				// check that liquidity is what we expected
				s.Require().Equal(expectedLiquidity.String(), pool.GetLiquidity().String())
				// also ensure the pool's currentLiquidity was updated due to calling a mutative method
				s.Require().Equal(expectedLiquidity.String(), updatedLiquidity.String())
			}
		})

	}
}

func (s *KeeperTestSuite) TestSwapExactAmountIn() {
	type param struct {
		tokenIn           sdk.Coin
		tokenOutDenom     string
		tokenOutMinAmount sdk.Int
		expectedTokenOut  sdk.Int
	}

	tests := []struct {
		name        string
		param       param
		expectedErr error
	}{
		{
			name: "Proper swap usdc > eth",
			// params
			// liquidity: 		 1517882343.751510418088349649
			// sqrtPriceNext:    70.738348247484497717 which is 5003.91391278239310954 https://www.wolframalpha.com/input?i=70.710678118654752440+%2B+42000000+%2F+1517882343.751510418088349649
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  41999999.999 rounded up https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2870.738348247484497717+-+70.710678118654752440%29
			// expectedTokenOut: 8396.7142421 rounded down https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2870.738348247484497717+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+70.738348247484497717%29
			param: param{
				tokenIn:           sdk.NewCoin(USDC, sdk.NewInt(42000000)),
				tokenOutDenom:     ETH,
				tokenOutMinAmount: types.LowerPriceLimit.RoundInt(),
				expectedTokenOut:  sdk.NewInt(8396),
			},
		},
		{
			name: "Proper swap eth > usdc",
			// params
			// liquidity: 		 1517882343.751510418088349649
			// sqrtPriceNext:    70.66666391085714433 which is 4993.77738829003954884402 https://www.wolframalpha.com/input?i=%28%281517882343.751510418088349649%29%29+%2F+%28%28%281517882343.751510418088349649%29+%2F+%2870.710678118654752440%29%29+%2B+%2813370%29%29
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  13370.0000 rounded up https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2870.710678118654752440+-+70.66666391085714433+%29%29+%2F+%2870.66666391085714433+*+70.710678118654752440%29
			// expectedTokenOut: 66808388.890 rounded down https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2870.710678118654752440+-+70.66666391085714433%29
			param: param{
				tokenIn:           sdk.NewCoin(ETH, sdk.NewInt(13370)),
				tokenOutDenom:     USDC,
				tokenOutMinAmount: types.LowerPriceLimit.RoundInt(),
				expectedTokenOut:  sdk.NewInt(66808388),
			},
		},
		{
			name: "out is lesser than min amount",
			param: param{
				tokenIn:           sdk.NewCoin(USDC, sdk.NewInt(42000000)),
				tokenOutDenom:     ETH,
				tokenOutMinAmount: sdk.NewInt(8397),
			},
			expectedErr: types.AmountLessThanMinError{TokenAmount: sdk.NewInt(8396), TokenMin: sdk.NewInt(8397)},
		},
		{
			name: "in and out denom are same",
			param: param{
				tokenIn:           sdk.NewCoin(ETH, sdk.NewInt(13370)),
				tokenOutDenom:     ETH,
				tokenOutMinAmount: types.LowerPriceLimit.RoundInt(),
			},
			expectedErr: types.DenomDuplicatedError{TokenInDenom: ETH, TokenOutDenom: ETH},
		},
		{
			name: "unknown in denom",
			param: param{
				tokenIn:           sdk.NewCoin("etha", sdk.NewInt(13370)),
				tokenOutDenom:     ETH,
				tokenOutMinAmount: types.LowerPriceLimit.RoundInt(),
			},
			expectedErr: types.TokenInDenomNotInPoolError{TokenInDenom: "etha"},
		},
		{
			name: "unknown out denom",
			param: param{
				tokenIn:           sdk.NewCoin(ETH, sdk.NewInt(13370)),
				tokenOutDenom:     "etha",
				tokenOutMinAmount: types.LowerPriceLimit.RoundInt(),
			},
			expectedErr: types.TokenOutDenomNotInPoolError{TokenOutDenom: "etha"},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// Init suite for each test.
			s.Setup()

			// Create a default CL pool
			pool := s.PrepareConcentratedPool()

			// Check the test case to see if we are swapping asset0 for asset1 or vice versa
			asset0 := pool.GetToken0()
			zeroForOne := test.param.tokenIn.Denom == asset0

			// Create a default position to the pool created earlier
			s.SetupDefaultPosition(1)

			// Fund the account with token in.
			s.FundAcc(s.TestAccs[0], sdk.NewCoins(test.param.tokenIn))

			// Retrieve pool post position set up
			pool, err := s.App.ConcentratedLiquidityKeeper.GetPoolById(s.Ctx, pool.GetId())
			s.Require().NoError(err)

			// Note spot price and gas used prior to swap
			spotPriceBefore := pool.GetCurrentSqrtPrice().Power(2)
			prevGasConsumed := s.Ctx.GasMeter().GasConsumed()

			// Execute the swap directed in the test case
			tokenOutAmount, err := s.App.ConcentratedLiquidityKeeper.SwapExactAmountIn(s.Ctx, s.TestAccs[0], pool.(poolmanagertypes.PoolI), test.param.tokenIn, test.param.tokenOutDenom, test.param.tokenOutMinAmount, DefaultZeroSwapFee)
			if test.expectedErr != nil {
				s.Require().Error(err)
				s.Require().ErrorContains(err, test.expectedErr.Error())
			} else {
				s.Require().NoError(err)
				s.Require().Equal(test.param.expectedTokenOut.String(), tokenOutAmount.String())

				gasConsumedForSwap := s.Ctx.GasMeter().GasConsumed() - prevGasConsumed

				// Check that we consume enough gas that a CL pool swap warrants
				// We consume `types.GasFeeForSwap` directly, so the extra I/O operation mean we end up consuming more.
				s.Require().Greater(gasConsumedForSwap, uint64(cltypes.ConcentratedGasFeeForSwap))

				// Assert events
				s.AssertEventEmitted(s.Ctx, cltypes.TypeEvtTokenSwapped, 1)

				// Retrieve pool again post swap
				pool, err = s.App.ConcentratedLiquidityKeeper.GetPoolById(s.Ctx, pool.GetId())
				s.Require().NoError(err)

				spotPriceAfter := pool.GetCurrentSqrtPrice().Power(2)

				// Ratio of the token out should be between the before spot price and after spot price.
				tradeAvgPrice := tokenOutAmount.ToDec().Quo(test.param.tokenIn.Amount.ToDec())

				if zeroForOne {
					s.Require().True(tradeAvgPrice.LT(spotPriceBefore))
					s.Require().True(tradeAvgPrice.GT(spotPriceAfter))
				} else {
					tradeAvgPrice = sdk.OneDec().Quo(tradeAvgPrice)
					s.Require().True(tradeAvgPrice.GT(spotPriceBefore))
					s.Require().True(tradeAvgPrice.LT(spotPriceAfter))
				}

			}
		})
	}
}

func (s *KeeperTestSuite) TestSwapExactAmountOut() {
	type param struct {
		tokenOut         sdk.Coin
		tokenInDenom     string
		tokenInMaxAmount sdk.Int
		expectedTokenIn  sdk.Int
	}

	tests := []struct {
		name        string
		param       param
		expectedErr error
	}{
		{
			name: "Proper swap eth > usdc",
			// params
			// liquidity: 		 1517882343.751510418088349649
			// sqrtPriceNext:    70.738349405152439867 which is 5003.914076565430543175 https://www.wolframalpha.com/input?i=70.710678118654752440+%2B+42000000+%2F+1517882343.751510418088349649
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  42000000.0000 rounded up https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2870.738349405152439867+-+70.710678118654752440%29
			// expectedTokenOut: 8396.714105 rounded down https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2870.738349405152439867+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+70.738349405152439867%29
			param: param{
				tokenOut:         sdk.NewCoin(USDC, sdk.NewInt(42000000)),
				tokenInDenom:     ETH,
				tokenInMaxAmount: types.UpperPriceLimit.RoundInt(),
				expectedTokenIn:  sdk.NewInt(8396),
			},
		},
		{
			name: "Proper swap usdc > eth",
			// params
			// liquidity: 		 1517882343.751510418088349649
			// sqrtPriceNext:    70.666662070529219856 which is 4993.777128190373086350 https://www.wolframalpha.com/input?i=%28%281517882343.751510418088349649%29%29+%2F+%28%28%281517882343.751510418088349649%29+%2F+%2870.710678118654752440%29%29+%2B+%2813370%29%29
			// expectedTokenIn:  13369.9999 rounded up https://www.wolframalpha.com/input?i=%281517882343.751510418088349649+*+%2870.710678118654752440+-+70.666662070529219856+%29%29+%2F+%2870.666662070529219856+*+70.710678118654752440%29
			// expectedTokenOut: 66808387.149 rounded down https://www.wolframalpha.com/input?i=1517882343.751510418088349649+*+%2870.710678118654752440+-+70.666662070529219856%29
			// expectedTick: 	 85163.7 rounded down https://www.wolframalpha.com/input?i2d=true&i=Log%5B1.0001%2C4993.777128190373086350%5D
			param: param{
				tokenOut:         sdk.NewCoin(ETH, sdk.NewInt(13370)),
				tokenInDenom:     USDC,
				tokenInMaxAmount: types.UpperPriceLimit.RoundInt(),
				expectedTokenIn:  sdk.NewInt(66808388),
			},
		},
		{
			name: "out is more than max amount",
			param: param{
				tokenOut:         sdk.NewCoin(USDC, sdk.NewInt(42000000)),
				tokenInDenom:     ETH,
				tokenInMaxAmount: types.LowerPriceLimit.RoundInt(),
			},
			expectedErr: types.AmountGreaterThanMaxError{TokenAmount: sdk.NewInt(8396), TokenMax: types.LowerPriceLimit.RoundInt()},
		},
		{
			name: "in and out denom are same",
			param: param{
				tokenOut:         sdk.NewCoin(ETH, sdk.NewInt(13370)),
				tokenInDenom:     ETH,
				tokenInMaxAmount: types.UpperPriceLimit.RoundInt(),
			},
			expectedErr: types.DenomDuplicatedError{TokenInDenom: ETH, TokenOutDenom: ETH},
		},
		{
			name: "unknown out denom",
			param: param{
				tokenOut:         sdk.NewCoin("etha", sdk.NewInt(13370)),
				tokenInDenom:     ETH,
				tokenInMaxAmount: types.UpperPriceLimit.RoundInt(),
			},
			expectedErr: types.TokenOutDenomNotInPoolError{TokenOutDenom: "etha"},
		},
		{
			name: "unknown in denom",
			param: param{
				tokenOut:         sdk.NewCoin(ETH, sdk.NewInt(13370)),
				tokenInDenom:     "etha",
				tokenInMaxAmount: types.UpperPriceLimit.RoundInt(),
			},
			expectedErr: types.TokenInDenomNotInPoolError{TokenInDenom: "etha"},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// Init suite for each test.
			s.Setup()

			// Create a default CL pool
			pool := s.PrepareConcentratedPool()

			// Check the test case to see if we are swapping asset0 for asset1 or vice versa
			asset0 := pool.GetToken0()
			zeroForOne := test.param.tokenOut.Denom == asset0

			// Chen create a default position to the pool created earlier
			s.SetupDefaultPosition(1)

			// Fund the account with token in.
			s.FundAcc(s.TestAccs[0], sdk.NewCoins(sdk.NewCoin(test.param.tokenInDenom, test.param.tokenInMaxAmount)))

			// Retrieve pool post position set up
			pool, err := s.App.ConcentratedLiquidityKeeper.GetPoolById(s.Ctx, pool.GetId())
			s.Require().NoError(err)

			// Note spot price and gas used prior to swap
			spotPriceBefore := pool.GetCurrentSqrtPrice().Power(2)
			prevGasConsumed := s.Ctx.GasMeter().GasConsumed()

			// Execute the swap directed in the test case
			tokenIn, err := s.App.ConcentratedLiquidityKeeper.SwapExactAmountOut(s.Ctx, s.TestAccs[0], pool.(poolmanagertypes.PoolI), test.param.tokenInDenom, test.param.tokenInMaxAmount, test.param.tokenOut, DefaultZeroSwapFee)

			if test.expectedErr != nil {
				s.Require().Error(err)
				s.Require().ErrorContains(err, test.expectedErr.Error())
			} else {
				s.Require().NoError(err)
				s.Require().Equal(test.param.expectedTokenIn.String(), tokenIn.String())

				gasConsumedForSwap := s.Ctx.GasMeter().GasConsumed() - prevGasConsumed
				// Check that we consume enough gas that a CL pool swap warrants
				// We consume `types.GasFeeForSwap` directly, so the extra I/O operation mean we end up consuming more.
				s.Require().Greater(gasConsumedForSwap, uint64(cltypes.ConcentratedGasFeeForSwap))

				// Assert events
				s.AssertEventEmitted(s.Ctx, cltypes.TypeEvtTokenSwapped, 1)

				// Retrieve pool again post swap
				pool, err = s.App.ConcentratedLiquidityKeeper.GetPoolById(s.Ctx, pool.GetId())
				s.Require().NoError(err)

				spotPriceAfter := pool.GetCurrentSqrtPrice().Power(2)

				// Ratio of the token out should be between the before spot price and after spot price.
				tradeAvgPrice := tokenIn.ToDec().Quo(test.param.tokenOut.Amount.ToDec())

				if zeroForOne {
					s.Require().True(tradeAvgPrice.LT(spotPriceBefore))
					s.Require().True(tradeAvgPrice.GT(spotPriceAfter))
				} else {
					tradeAvgPrice = sdk.OneDec().Quo(tradeAvgPrice)
					s.Require().True(tradeAvgPrice.GT(spotPriceBefore))
					s.Require().True(tradeAvgPrice.LT(spotPriceAfter))
				}

			}
		})
	}
}
