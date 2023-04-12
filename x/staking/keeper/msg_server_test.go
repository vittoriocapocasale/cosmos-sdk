package keeper_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"

	"github.com/golang/mock/gomock"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var PKS = simtestutil.CreateTestPubKeys(3)

func (s *KeeperTestSuite) TestMsgCreateValidator() {
	ctx, keeper, msgServer := s.ctx, s.stakingKeeper, s.msgServer
	require := s.Require()

	delAddr := sdk.AccAddress(PKS[0].Address())
	pk1 := ed25519.GenPrivKey().PubKey()
	require.NotNil(pk1)

	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	valAddr := sdk.ValAddress(authority)

	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), stakingtypes.ModuleName, gomock.Any()).AnyTimes()
	amt := keeper.TokensFromConsensusPower(s.ctx, int64(100))
	s.bankKeeper.MintCoins(s.ctx, stakingtypes.ModuleName, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, amt)))

	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), stakingtypes.ModuleName, delAddr, gomock.Any()).AnyTimes()
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, stakingtypes.ModuleName, delAddr, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, amt)))
	s.bankKeeper.EXPECT().DelegateCoinsFromAccountToModule(gomock.Any(), authority, stakingtypes.NotBondedPoolName, gomock.Any()).AnyTimes()

	pubkey, err := codectypes.NewAnyWithValue(pk1)
	require.NoError(err)

	testCases := []struct {
		name      string
		input     *stakingtypes.MsgCreateValidator
		expErr    bool
		expErrMsg string
	}{
		{
			name: "empty description",
			input: &stakingtypes.MsgCreateValidator{
				Description: stakingtypes.Description{},
				Commission: stakingtypes.CommissionRates{
					Rate:          sdk.NewDecWithPrec(5, 1),
					MaxRate:       sdk.NewDecWithPrec(5, 1),
					MaxChangeRate: math.LegacyNewDec(0),
				},
				MinSelfDelegation: math.NewInt(1),
				DelegatorAddress:  delAddr.String(),
				ValidatorAddress:  valAddr.String(),
				Pubkey:            pubkey,
				Value:             sdk.NewInt64Coin("stake", 10000),
			},
			expErr:    true,
			expErrMsg: "empty description",
		},
		{
			name: "invalid validator address",
			input: &stakingtypes.MsgCreateValidator{
				Description: stakingtypes.Description{
					Moniker: "NewValidator",
				},
				Commission: stakingtypes.CommissionRates{
					Rate:          sdk.NewDecWithPrec(5, 1),
					MaxRate:       sdk.NewDecWithPrec(5, 1),
					MaxChangeRate: math.LegacyNewDec(0),
				},
				MinSelfDelegation: math.NewInt(1),
				DelegatorAddress:  delAddr.String(),
				ValidatorAddress:  sdk.AccAddress([]byte("invalid")).String(),
				Pubkey:            pubkey,
				Value:             sdk.NewInt64Coin("stake", 10000),
			},
			expErr:    true,
			expErrMsg: "invalid validator address",
		},
		{
			name: "empty validator pubkey",
			input: &stakingtypes.MsgCreateValidator{
				Description: stakingtypes.Description{
					Moniker: "NewValidator",
				},
				Commission: stakingtypes.CommissionRates{
					Rate:          sdk.NewDecWithPrec(5, 1),
					MaxRate:       sdk.NewDecWithPrec(5, 1),
					MaxChangeRate: math.LegacyNewDec(0),
				},
				MinSelfDelegation: math.NewInt(1),
				DelegatorAddress:  delAddr.String(),
				ValidatorAddress:  valAddr.String(),
				Pubkey:            nil,
				Value:             sdk.NewInt64Coin("stake", 10000),
			},
			expErr:    true,
			expErrMsg: "empty validator public key",
		},
		{
			name: "empty delegation amount",
			input: &stakingtypes.MsgCreateValidator{
				Description: stakingtypes.Description{
					Moniker: "NewValidator",
				},
				Commission: stakingtypes.CommissionRates{
					Rate:          sdk.NewDecWithPrec(5, 1),
					MaxRate:       sdk.NewDecWithPrec(5, 1),
					MaxChangeRate: math.LegacyNewDec(0),
				},
				MinSelfDelegation: math.NewInt(1),
				DelegatorAddress:  delAddr.String(),
				ValidatorAddress:  valAddr.String(),
				Pubkey:            pubkey,
				Value:             sdk.NewInt64Coin("stake", 0),
			},
			expErr:    true,
			expErrMsg: "invalid delegation amount",
		},
		{
			name: "nil delegation amount",
			input: &stakingtypes.MsgCreateValidator{
				Description: stakingtypes.Description{
					Moniker: "NewValidator",
				},
				Commission: stakingtypes.CommissionRates{
					Rate:          sdk.NewDecWithPrec(5, 1),
					MaxRate:       sdk.NewDecWithPrec(5, 1),
					MaxChangeRate: math.LegacyNewDec(0),
				},
				MinSelfDelegation: math.NewInt(1),
				DelegatorAddress:  delAddr.String(),
				ValidatorAddress:  valAddr.String(),
				Pubkey:            pubkey,
				Value:             sdk.Coin{},
			},
			expErr:    true,
			expErrMsg: "invalid delegation amount",
		},
		{
			name: "zero minimum self delegation",
			input: &stakingtypes.MsgCreateValidator{
				Description: stakingtypes.Description{
					Moniker: "NewValidator",
				},
				Commission: stakingtypes.CommissionRates{
					Rate:          sdk.NewDecWithPrec(5, 1),
					MaxRate:       sdk.NewDecWithPrec(5, 1),
					MaxChangeRate: math.LegacyNewDec(0),
				},
				MinSelfDelegation: math.NewInt(0),
				DelegatorAddress:  delAddr.String(),
				ValidatorAddress:  valAddr.String(),
				Pubkey:            pubkey,
				Value:             sdk.NewInt64Coin("stake", 10000),
			},
			expErr:    true,
			expErrMsg: "minimum self delegation must be a positive integer",
		},
		{
			name: "negative minimum self delegation",
			input: &stakingtypes.MsgCreateValidator{
				Description: stakingtypes.Description{
					Moniker: "NewValidator",
				},
				Commission: stakingtypes.CommissionRates{
					Rate:          sdk.NewDecWithPrec(5, 1),
					MaxRate:       sdk.NewDecWithPrec(5, 1),
					MaxChangeRate: math.LegacyNewDec(0),
				},
				MinSelfDelegation: math.NewInt(-1),
				DelegatorAddress:  delAddr.String(),
				ValidatorAddress:  valAddr.String(),
				Pubkey:            pubkey,
				Value:             sdk.NewInt64Coin("stake", 10000),
			},
			expErr:    true,
			expErrMsg: "minimum self delegation must be a positive integer",
		},
		{
			name: "delegation less than minimum self delegation",
			input: &stakingtypes.MsgCreateValidator{
				Description: stakingtypes.Description{
					Moniker: "NewValidator",
				},
				Commission: stakingtypes.CommissionRates{
					Rate:          sdk.NewDecWithPrec(5, 1),
					MaxRate:       sdk.NewDecWithPrec(5, 1),
					MaxChangeRate: math.LegacyNewDec(0),
				},
				MinSelfDelegation: math.NewInt(100),
				DelegatorAddress:  delAddr.String(),
				ValidatorAddress:  valAddr.String(),
				Pubkey:            pubkey,
				Value:             sdk.NewInt64Coin("stake", 10),
			},
			expErr:    true,
			expErrMsg: "validator's self delegation must be greater than their minimum self delegation",
		},
		{
			name: "valid msg",
			input: &stakingtypes.MsgCreateValidator{
				Description: stakingtypes.Description{
					Moniker: "NewValidator",
				},
				Commission: stakingtypes.CommissionRates{
					Rate:          sdk.NewDecWithPrec(5, 1),
					MaxRate:       sdk.NewDecWithPrec(5, 1),
					MaxChangeRate: math.LegacyNewDec(0),
				},
				MinSelfDelegation: math.NewInt(1),
				DelegatorAddress:  delAddr.String(),
				ValidatorAddress:  valAddr.String(),
				Pubkey:            pubkey,
				Value:             sdk.NewInt64Coin("stake", 10000),
			},
			expErr: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.T().Run(tc.name, func(t *testing.T) {
			_, err := msgServer.CreateValidator(ctx, tc.input)
			if tc.expErr {
				require.Error(err)
				require.Contains(err.Error(), tc.expErrMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgUpdateParams() {
	ctx, keeper, msgServer := s.ctx, s.stakingKeeper, s.msgServer
	require := s.Require()

	testCases := []struct {
		name      string
		input     *stakingtypes.MsgUpdateParams
		expErr    bool
		expErrMsg string
	}{
		{
			name: "valid params",
			input: &stakingtypes.MsgUpdateParams{
				Authority: keeper.GetAuthority(),
				Params:    stakingtypes.DefaultParams(),
			},
			expErr: false,
		},
		{
			name: "invalid authority",
			input: &stakingtypes.MsgUpdateParams{
				Authority: "invalid",
				Params:    stakingtypes.DefaultParams(),
			},
			expErr:    true,
			expErrMsg: "invalid authority",
		},
		{
			name: "negative commission rate",
			input: &stakingtypes.MsgUpdateParams{
				Authority: keeper.GetAuthority(),
				Params: stakingtypes.Params{
					MinCommissionRate: math.LegacyNewDec(-10),
					UnbondingTime:     stakingtypes.DefaultUnbondingTime,
					MaxValidators:     stakingtypes.DefaultMaxValidators,
					MaxEntries:        stakingtypes.DefaultMaxEntries,
					HistoricalEntries: stakingtypes.DefaultHistoricalEntries,
					BondDenom:         stakingtypes.BondStatusBonded,
				},
			},
			expErr:    true,
			expErrMsg: "minimum commission rate cannot be negative",
		},
		{
			name: "commission rate cannot be bigger than 100",
			input: &stakingtypes.MsgUpdateParams{
				Authority: keeper.GetAuthority(),
				Params: stakingtypes.Params{
					MinCommissionRate: math.LegacyNewDec(2),
					UnbondingTime:     stakingtypes.DefaultUnbondingTime,
					MaxValidators:     stakingtypes.DefaultMaxValidators,
					MaxEntries:        stakingtypes.DefaultMaxEntries,
					HistoricalEntries: stakingtypes.DefaultHistoricalEntries,
					BondDenom:         stakingtypes.BondStatusBonded,
				},
			},
			expErr:    true,
			expErrMsg: "minimum commission rate cannot be greater than 100%",
		},
		{
			name: "invalid bond denom",
			input: &stakingtypes.MsgUpdateParams{
				Authority: keeper.GetAuthority(),
				Params: stakingtypes.Params{
					MinCommissionRate: stakingtypes.DefaultMinCommissionRate,
					UnbondingTime:     stakingtypes.DefaultUnbondingTime,
					MaxValidators:     stakingtypes.DefaultMaxValidators,
					MaxEntries:        stakingtypes.DefaultMaxEntries,
					HistoricalEntries: stakingtypes.DefaultHistoricalEntries,
					BondDenom:         "",
				},
			},
			expErr:    true,
			expErrMsg: "bond denom cannot be blank",
		},
		{
			name: "max validators must be positive",
			input: &stakingtypes.MsgUpdateParams{
				Authority: keeper.GetAuthority(),
				Params: stakingtypes.Params{
					MinCommissionRate: stakingtypes.DefaultMinCommissionRate,
					UnbondingTime:     stakingtypes.DefaultUnbondingTime,
					MaxValidators:     0,
					MaxEntries:        stakingtypes.DefaultMaxEntries,
					HistoricalEntries: stakingtypes.DefaultHistoricalEntries,
					BondDenom:         stakingtypes.BondStatusBonded,
				},
			},
			expErr:    true,
			expErrMsg: "max validators must be positive",
		},
		{
			name: "max entries most be positive",
			input: &stakingtypes.MsgUpdateParams{
				Authority: keeper.GetAuthority(),
				Params: stakingtypes.Params{
					MinCommissionRate: stakingtypes.DefaultMinCommissionRate,
					UnbondingTime:     stakingtypes.DefaultUnbondingTime,
					MaxValidators:     stakingtypes.DefaultMaxValidators,
					MaxEntries:        0,
					HistoricalEntries: stakingtypes.DefaultHistoricalEntries,
					BondDenom:         stakingtypes.BondStatusBonded,
				},
			},
			expErr:    true,
			expErrMsg: "max entries must be positive",
		},
		{
			name: "negative unbounding time",
			input: &stakingtypes.MsgUpdateParams{
				Authority: keeper.GetAuthority(),
				Params: stakingtypes.Params{
					UnbondingTime:     time.Hour * 24 * 7 * 3 * -1,
					MaxEntries:        stakingtypes.DefaultMaxEntries,
					MaxValidators:     stakingtypes.DefaultMaxValidators,
					HistoricalEntries: stakingtypes.DefaultHistoricalEntries,
					MinCommissionRate: stakingtypes.DefaultMinCommissionRate,
					BondDenom:         "denom",
				},
			},
			expErr:    true,
			expErrMsg: "unbonding time must be positive",
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.T().Run(tc.name, func(t *testing.T) {
			_, err := msgServer.UpdateParams(ctx, tc.input)
			if tc.expErr {
				require.Error(err)
				require.Contains(err.Error(), tc.expErrMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}
