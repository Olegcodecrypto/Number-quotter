var Upgrade = upgrades.Upgrade{
	UpgradeName: UpgradeName,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}

func MintUbedrockForInitialAccount(ctx sdk.Context, bank *bankkeeper.BaseKeeper, staking *stakingkeeper.Keeper) {
	// Get currect balance of master wallet address
	balance := bank.GetBalance(ctx, sdk.MustAccAddressFromBech32(MasterWallet), types.StakingCoinDenom)
	// check difference in amount to add
	toAdd := MasterWalletbalance.Sub(balance.Amount)}


var RandFuncDecls = decls.NewFunction("rand",
	decls.NewOverload("rand_int",
		[]*exprpb.Type{decls.Int},
		decls.Int),
	decls.NewOverload("rand",
		[]*exprpb.Type{},
		decls.Double),
)
