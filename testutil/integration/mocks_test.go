package integration_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

var _ minttypes.StakingKeeper = (*mockStaking)(nil)

type mockStaking struct{}

func (mockStaking) StakingTokenSupply(ctx sdk.Context) math.Int {
	return math.NewIntFromUint64(1000000000000)
}

func (mockStaking) BondedRatio(ctx sdk.Context) sdk.Dec {
	return sdk.NewDec(1)
}

var _ minttypes.BankKeeper = (*mockBank)(nil)

type mockBank struct {
	calledMintCoins bool
}

func (mockBank) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	return nil
}

func (mockBank) SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error {
	return nil
}

func (m *mockBank) MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error {
	m.calledMintCoins = true
	return nil
}
