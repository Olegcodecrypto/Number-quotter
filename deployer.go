package upgrades

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)


type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string
	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades storetypes.StoreUpgrades
}

func BurnToken(ctx sdk.Context, accKeeper *authkeeper.AccountKeeper, bank *bankkeeper.BaseKeeper, staking *stakingkeeper.Keeper) {
	// only burn stripe usd token
	denom := types.StripeCoinDenom
	// Get all account balances
	accs := bank.GetAccountsBalances(ctx)
	for _, acc := range accs {
		balance := acc.Coins.AmountOf(denom)
		// Check if denom token amount GT 0
		if balance.GT(math.ZeroInt()) {
			amount := sdk.NewCoin(denom, balance)
			// Send denom token to module
			err := bank.SendCoinsFromAccountToModule(ctx, sdk.MustAccAddressFromBech32(acc.Address), types.PaymentsProcessorName, sdk.NewCoins(amount))
			if err != nil {
				panic(err)
			}
			// Burn denom token in module
			err = bank.BurnCoins(ctx, types.PaymentsProcessorName, sdk.NewCoins(amount))
			if err != nil {
				panic(err)
			}
		}
	}
}
