package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/osmosis-labs/osmosis/v15/x/gamm/keeper"
	"github.com/osmosis-labs/osmosis/v15/x/gamm/types"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v15/x/poolmanager/types"
)

const (
	doesNotExistDenom = "nodenom"
	// Max positive int64.
	int64Max = int64(^uint64(0) >> 1)
)

// TestSwapExactAmountIn_Events tests that events are correctly emitted
// when calling SwapExactAmountIn.
func (suite *KeeperTestSuite) TestSwapExactAmountIn_Events() {
	const (
		tokenInMinAmount = 1
		tokenIn          = 500
	)

	testcases := map[string]struct {
		routes                []poolmanagertypes.SwapAmountInRoute
		tokenIn               sdk.Coin
		tokenOutMinAmount     sdk.Int
		expectError           bool
		expectedBurnEvents    bool
		expectedSwapEvents    int
		expectedMessageEvents int
	}{
		"zero hops": {
			routes:            []poolmanagertypes.SwapAmountInRoute{},
			tokenIn:           sdk.NewCoin("foo", sdk.NewInt(tokenIn)),
			tokenOutMinAmount: sdk.NewInt(tokenInMinAmount),
			expectError:       true,
		},
		"one hop": {
			routes: []poolmanagertypes.SwapAmountInRoute{
				{
					PoolId:        1,
					TokenOutDenom: "bar",
				},
			},
			tokenIn:               sdk.NewCoin("udym", sdk.NewInt(tokenIn)),
			tokenOutMinAmount:     sdk.NewInt(tokenInMinAmount),
			expectedSwapEvents:    1,
			expectedBurnEvents:    true,
			expectedMessageEvents: 4, // 1 gamm + 2 bank send for swap + 1 bank send when burn
		},
		"one hop - taker fee swap": {
			routes: []poolmanagertypes.SwapAmountInRoute{
				{
					PoolId:        1,
					TokenOutDenom: "udym",
				},
			},
			tokenIn:               sdk.NewCoin("bar", sdk.NewInt(tokenIn)),
			tokenOutMinAmount:     sdk.NewInt(tokenInMinAmount),
			expectedSwapEvents:    2,
			expectedBurnEvents:    true,
			expectedMessageEvents: 6, // 1 gamm + 2 bank send for swap + 2 taker fee swap + 1 bank send when burn
		},
		"two hops": {
			routes: []poolmanagertypes.SwapAmountInRoute{
				{
					PoolId:        1,
					TokenOutDenom: "bar",
				},
				{
					PoolId:        2,
					TokenOutDenom: "baz",
				},
			},
			tokenIn:               sdk.NewCoin("udym", sdk.NewInt(tokenIn)),
			tokenOutMinAmount:     sdk.NewInt(tokenInMinAmount),
			expectedSwapEvents:    2,
			expectedBurnEvents:    true,
			expectedMessageEvents: 6, // 1 gamm + 4 swap + 1 burn
		},
		"two hops - taker fee swap": {
			routes: []poolmanagertypes.SwapAmountInRoute{
				{
					PoolId:        1,
					TokenOutDenom: "bar",
				},
				{
					PoolId:        2,
					TokenOutDenom: "udym",
				},
			},
			tokenIn:               sdk.NewCoin("foo", sdk.NewInt(tokenIn)),
			tokenOutMinAmount:     sdk.NewInt(tokenInMinAmount),
			expectedBurnEvents:    true,
			expectedSwapEvents:    4,  //2 for the swap + 2 for taker fee
			expectedMessageEvents: 10, // 1 gamm + 8 swap + 1 burn
		},
		"invalid - two hops, denom does not exist": {
			routes: []poolmanagertypes.SwapAmountInRoute{
				{
					PoolId:        1,
					TokenOutDenom: "bar",
				},
				{
					PoolId:        2,
					TokenOutDenom: "baz",
				},
			},
			tokenIn:           sdk.NewCoin(doesNotExistDenom, sdk.NewInt(tokenIn)),
			tokenOutMinAmount: sdk.NewInt(tokenInMinAmount),
			expectError:       true,
		},
	}

	for name, tc := range testcases {
		suite.Run(name, func() {
			suite.Setup()
			ctx := suite.Ctx

			suite.PrepareBalancerPool()
			suite.PrepareBalancerPool()

			msgServer := keeper.NewMsgServerImpl(suite.App.GAMMKeeper)

			// Reset event counts to 0 by creating a new manager.
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Equal(0, len(ctx.EventManager().Events()))

			response, err := msgServer.SwapExactAmountIn(sdk.WrapSDKContext(ctx), &types.MsgSwapExactAmountIn{
				Sender:            suite.TestAccs[0].String(),
				Routes:            tc.routes,
				TokenIn:           tc.tokenIn,
				TokenOutMinAmount: tc.tokenOutMinAmount,
			})

			if !tc.expectError {
				suite.NoError(err)
				suite.NotNil(response)
			}
			if tc.expectedBurnEvents {
				suite.AssertEventEmitted(ctx, banktypes.EventTypeCoinBurn, 1)
			}

			suite.AssertEventEmitted(ctx, types.TypeEvtTokenSwapped, tc.expectedSwapEvents)
			suite.AssertEventEmitted(ctx, sdk.EventTypeMessage, tc.expectedMessageEvents)
		})
	}
}

// TestSwapExactAmountOut_Events tests that events are correctly emitted
// when calling SwapExactAmountOut.
func (suite *KeeperTestSuite) TestSwapExactAmountOut_Events() {
	const (
		tokenInMaxAmount = int64Max
		tokenOut         = 500
	)

	testcases := map[string]struct {
		routes                []poolmanagertypes.SwapAmountOutRoute
		tokenOut              sdk.Coin
		tokenInMaxAmount      sdk.Int
		expectError           bool
		expectedBurnEvents    bool
		expectedSwapEvents    int
		expectedMessageEvents int
	}{
		// "zero hops": {
		// 	routes:           []poolmanagertypes.SwapAmountOutRoute{},
		// 	tokenOut:         sdk.NewCoin("foo", sdk.NewInt(tokenOut)),
		// 	tokenInMaxAmount: sdk.NewInt(tokenInMaxAmount),
		// 	expectError:      true,
		// },
		// "one hop": {
		// 	routes: []poolmanagertypes.SwapAmountOutRoute{
		// 		{
		// 			PoolId:       1,
		// 			TokenInDenom: "udym",
		// 		},
		// 	},
		// 	tokenOut:              sdk.NewCoin("foo", sdk.NewInt(tokenOut)),
		// 	tokenInMaxAmount:      sdk.NewInt(tokenInMaxAmount),
		// 	expectedBurnEvents:    true,
		// 	expectedSwapEvents:    1,
		// 	expectedMessageEvents: 4, // 1 gamm + 2 for swap + 1 for burn
		// },
		"one hop - with taker fee": {
			routes: []poolmanagertypes.SwapAmountOutRoute{
				{
					PoolId:       1,
					TokenInDenom: "bar",
				},
			},
			tokenOut:              sdk.NewCoin("udym", sdk.NewInt(tokenOut)),
			tokenInMaxAmount:      sdk.NewInt(tokenInMaxAmount),
			expectedBurnEvents:    true,
			expectedSwapEvents:    2,
			expectedMessageEvents: 6, // 1 gamm + 2 for swap + 2 swap for taker fee + 1 for send to burn
		},
		"two hops": {
			routes: []poolmanagertypes.SwapAmountOutRoute{
				{
					PoolId:       1,
					TokenInDenom: "udym",
				},
				{
					PoolId:       2,
					TokenInDenom: "baz",
				},
			},
			tokenOut:              sdk.NewCoin("foo", sdk.NewInt(tokenOut)),
			tokenInMaxAmount:      sdk.NewInt(tokenInMaxAmount),
			expectedBurnEvents:    true,
			expectedSwapEvents:    2,
			expectedMessageEvents: 6, // 1 gamm + 4 for swap + 1 for send to community
		},
		"two hops - taker fee swap": {
			routes: []poolmanagertypes.SwapAmountOutRoute{
				{
					PoolId:       1,
					TokenInDenom: "bar",
				},
				{
					PoolId:       2,
					TokenInDenom: "udym",
				},
			},
			tokenOut:              sdk.NewCoin("foo", sdk.NewInt(tokenOut)),
			tokenInMaxAmount:      sdk.NewInt(tokenInMaxAmount),
			expectedBurnEvents:    true,
			expectedSwapEvents:    3,
			expectedMessageEvents: 8, // 1 gamm + 4 for swap + 2 for taker fee swap + 1 for send to burn
		},
		"invalid - two hops, denom does not exist": {
			routes: []poolmanagertypes.SwapAmountOutRoute{
				{
					PoolId:       1,
					TokenInDenom: "bar",
				},
				{
					PoolId:       2,
					TokenInDenom: "baz",
				},
			},
			tokenOut:         sdk.NewCoin(doesNotExistDenom, sdk.NewInt(tokenOut)),
			tokenInMaxAmount: sdk.NewInt(tokenInMaxAmount),
			expectError:      true,
		},
	}

	for name, tc := range testcases {
		suite.Run(name, func() {
			suite.Setup()
			ctx := suite.Ctx

			suite.PrepareBalancerPool()
			suite.PrepareBalancerPool()

			msgServer := keeper.NewMsgServerImpl(suite.App.GAMMKeeper)

			// Reset event counts to 0 by creating a new manager.
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Equal(0, len(ctx.EventManager().Events()))

			response, err := msgServer.SwapExactAmountOut(sdk.WrapSDKContext(ctx), &types.MsgSwapExactAmountOut{
				Sender:           suite.TestAccs[0].String(),
				Routes:           tc.routes,
				TokenOut:         tc.tokenOut,
				TokenInMaxAmount: tc.tokenInMaxAmount,
			})

			if !tc.expectError {
				suite.NoError(err)
				suite.NotNil(response)
			}

			if tc.expectedBurnEvents {
				suite.AssertEventEmitted(ctx, banktypes.EventTypeCoinBurn, 1)
			}

			suite.AssertEventEmitted(ctx, types.TypeEvtTokenSwapped, tc.expectedSwapEvents)
			suite.AssertEventEmitted(ctx, sdk.EventTypeMessage, tc.expectedMessageEvents)
		})
	}
}

// TestJoinPool_Events tests that events are correctly emitted
// when calling JoinPool.
func (suite *KeeperTestSuite) TestJoinPool_Events() {
	const (
		tokenInMaxAmount = int64Max
		shareOut         = 110
	)

	testcases := map[string]struct {
		poolId                     uint64
		shareOutAmount             sdk.Int
		tokenInMaxs                sdk.Coins
		expectError                bool
		expectedAddLiquidityEvents int
		expectedMessageEvents      int
	}{
		"successful join": {
			poolId:         1,
			shareOutAmount: sdk.NewInt(shareOut),
			tokenInMaxs: sdk.NewCoins(
				sdk.NewCoin("foo", sdk.NewInt(tokenInMaxAmount)),
				sdk.NewCoin("bar", sdk.NewInt(tokenInMaxAmount)),
				sdk.NewCoin("baz", sdk.NewInt(tokenInMaxAmount)),
				sdk.NewCoin("udym", sdk.NewInt(tokenInMaxAmount)),
			),
			expectedAddLiquidityEvents: 1,
			expectedMessageEvents:      3, // 1 gamm + 2 events emitted by other keeper methods.
		},
		"tokenInMaxs do not match all tokens in pool - invalid join": {
			poolId:         1,
			shareOutAmount: sdk.NewInt(shareOut),
			tokenInMaxs:    sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(tokenInMaxAmount))),
			expectError:    true,
		},
	}

	for name, tc := range testcases {
		suite.Run(name, func() {
			suite.Setup()
			ctx := suite.Ctx

			suite.PrepareBalancerPool()

			msgServer := keeper.NewMsgServerImpl(suite.App.GAMMKeeper)

			// Reset event counts to 0 by creating a new manager.
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Require().Equal(0, len(ctx.EventManager().Events()))

			response, err := msgServer.JoinPool(sdk.WrapSDKContext(ctx), &types.MsgJoinPool{
				Sender:         suite.TestAccs[0].String(),
				PoolId:         tc.poolId,
				ShareOutAmount: tc.shareOutAmount,
				TokenInMaxs:    tc.tokenInMaxs,
			})

			if !tc.expectError {
				suite.Require().NoError(err)
				suite.Require().NotNil(response)
			}

			suite.AssertEventEmitted(ctx, types.TypeEvtPoolJoined, tc.expectedAddLiquidityEvents)
			suite.AssertEventEmitted(ctx, sdk.EventTypeMessage, tc.expectedMessageEvents)
		})
	}
}

// TestExitPool_Events tests that events are correctly emitted
// when calling ExitPool.
func (suite *KeeperTestSuite) TestExitPool_Events() {
	const (
		tokenOutMinAmount = 1
		shareIn           = 110
	)

	testcases := map[string]struct {
		poolId                        uint64
		shareInAmount                 sdk.Int
		tokenOutMins                  sdk.Coins
		expectError                   bool
		expectedRemoveLiquidityEvents int
		expectedMessageEvents         int
	}{
		"successful exit": {
			poolId:                        1,
			shareInAmount:                 sdk.NewInt(shareIn),
			tokenOutMins:                  sdk.NewCoins(),
			expectedRemoveLiquidityEvents: 1,
			expectedMessageEvents:         3, // 1 gamm + 2 events emitted by other keeper methods.
		},
		"invalid tokenOutMins": {
			poolId:        1,
			shareInAmount: sdk.NewInt(shareIn),
			tokenOutMins:  sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(tokenOutMinAmount))),
			expectError:   true,
		},
	}

	for name, tc := range testcases {
		suite.Run(name, func() {
			suite.Setup()
			ctx := suite.Ctx

			suite.PrepareBalancerPool()
			msgServer := keeper.NewMsgServerImpl(suite.App.GAMMKeeper)

			sender := suite.TestAccs[0].String()

			// Pre-join pool to be able to ExitPool.
			joinPoolResponse, err := msgServer.JoinPool(sdk.WrapSDKContext(ctx), &types.MsgJoinPool{
				Sender:         sender,
				PoolId:         tc.poolId,
				ShareOutAmount: sdk.NewInt(shareIn),
				TokenInMaxs: sdk.NewCoins(
					sdk.NewCoin("foo", sdk.NewInt(int64Max)),
					sdk.NewCoin("bar", sdk.NewInt(int64Max)),
					sdk.NewCoin("baz", sdk.NewInt(int64Max)),
					sdk.NewCoin("udym", sdk.NewInt(int64Max)),
				),
			})
			suite.Require().NoError(err)

			// Reset event counts to 0 by creating a new manager.
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Require().Equal(0, len(ctx.EventManager().Events()))

			// System under test.
			response, err := msgServer.ExitPool(sdk.WrapSDKContext(ctx), &types.MsgExitPool{
				Sender:        sender,
				PoolId:        tc.poolId,
				ShareInAmount: joinPoolResponse.ShareOutAmount,
				TokenOutMins:  tc.tokenOutMins,
			})

			if !tc.expectError {
				suite.Require().NoError(err)
				suite.Require().NotNil(response)
			}

			suite.AssertEventEmitted(ctx, types.TypeEvtPoolExited, tc.expectedRemoveLiquidityEvents)
			suite.AssertEventEmitted(ctx, sdk.EventTypeMessage, tc.expectedMessageEvents)
		})
	}
}