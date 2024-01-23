package upgrades

import (
	"fmt"
	"time"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authvestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	"github.com/CosmosContracts/juno/v19/app/keepers"
)

// Stops a vesting account and returns all tokens back to the Core-1 SubDAO.
func MoveVestingCoinFromVestingAccount(ctx sdk.Context, keepers *keepers.AppKeepers, bondDenom string, name string, accAddr sdk.AccAddress, councilAccAddr sdk.AccAddress) error {
	now := ctx.BlockHeader().Time

	stdAcc := keepers.AccountKeeper.GetAccount(ctx, accAddr)
	vacc, ok := stdAcc.(*authvestingtypes.PeriodicVestingAccount)
	if !ok {
		// For e2e testing
		fmt.Printf("account " + accAddr.String() + " is not a vesting account.\n")
		return nil
	}

	fmt.Printf("\n\n== Vesting Account Address: %s (%s) ==\n", vacc.GetAddress().String(), name)

	// Gets vesting coins (These get returned back to Core-1 SubDAO / a new vesting contract made from)
	vestingCoins := vacc.GetVestingCoins(now)
	fmt.Printf("Vesting Coins: %v\n", vestingCoins)

	// Instantly complete any re-deleations.
	amt, err := completeAllRedelegations(ctx, now, keepers, accAddr)
	if err != nil {
		return err
	}
	fmt.Println("Redelegated Amount: ", amt)

	// Instantly unbond all delegations.
	amt, err = unbondAllAndFinish(ctx, now, keepers, accAddr)
	if err != nil {
		return err
	}
	fmt.Println("Unbonded Amount: ", amt)

	// Pre transfer balance
	councilBeforeBal := keepers.BankKeeper.GetBalance(ctx, councilAccAddr, bondDenom)

	// Moves unvested tokens to the Core-1 SubDAO.
	if err := transferUnvestedTokensToCouncil(ctx, keepers, bondDenom, accAddr, councilAccAddr, vestingCoins); err != nil {
		return err
	}

	// Set the vesting account to a base account
	keepers.AccountKeeper.SetAccount(ctx, vacc.BaseAccount)

	// Ensure the post validation checks are met.
	err = postValidation(ctx, keepers, bondDenom, accAddr, councilAccAddr, vestingCoins, councilBeforeBal)
	return err
}

func postValidation(ctx sdk.Context, keepers *keepers.AppKeepers, bondDenom string, accAddr sdk.AccAddress, councilAccAddr sdk.AccAddress, vestingCoins sdk.Coins, councilBeforeBal sdk.Coin) error {
	// Council balance should only increase by exactly the council + vestedCoins
	councilAfterBal := keepers.BankKeeper.GetBalance(ctx, councilAccAddr, bondDenom)
	if !councilBeforeBal.Add(vestingCoins[0]).IsEqual(councilAfterBal) {
		return fmt.Errorf("ERROR: core1BeforeBal (%v) + unvestedCoins (%v) != core1BalAfter (%v)", councilBeforeBal, vestingCoins, councilAfterBal)
	}

	// vesting account should have no future vesting periods
	newVacc := keepers.AccountKeeper.GetAccount(ctx, accAddr)
	if _, ok := newVacc.(*authvestingtypes.PeriodicVestingAccount); ok {
		return fmt.Errorf("ERROR: account %s still is a vesting account", accAddr.String())
	}

	// ensure the account has 0 delegations, redelegations, or unbonding delegations
	delegations := keepers.StakingKeeper.GetAllDelegatorDelegations(ctx, accAddr)
	if len(delegations) != 0 {
		return fmt.Errorf("ERROR: account %s still has delegations", accAddr.String())
	}

	redelegations := keepers.StakingKeeper.GetRedelegations(ctx, accAddr, 65535)
	if len(redelegations) != 0 {
		return fmt.Errorf("ERROR: account %s still has redelegations", accAddr.String())
	}

	unbondingDelegations := keepers.StakingKeeper.GetAllUnbondingDelegations(ctx, accAddr)
	if len(unbondingDelegations) != 0 {
		return fmt.Errorf("ERROR: account %s still has unbonding delegations", accAddr.String())
	}

	return nil
}

// Transfer funds from the vesting account to the Council SubDAO.
func transferUnvestedTokensToCouncil(ctx sdk.Context, keepers *keepers.AppKeepers, bondDenom string, accAddr, councilAccAddr sdk.AccAddress, unvestedCoins sdk.Coins) error {
	fmt.Printf("Sending Vesting Coins to Council: %v\n", unvestedCoins)
	if err := keepers.BankKeeper.SendCoins(ctx, accAddr, councilAccAddr, unvestedCoins); err != nil {
		return err
	}

	councilBal := keepers.BankKeeper.GetBalance(ctx, councilAccAddr, bondDenom)
	fmt.Printf("Updated Council SubDAO Balance: %v\n", councilBal)

	return nil
}

// Completes all re-delegations and returns the amount of tokens which were re-delegated.
func completeAllRedelegations(ctx sdk.Context, now time.Time, keepers *keepers.AppKeepers, accAddr sdk.AccAddress) (math.Int, error) {
	redelegatedAmt := math.ZeroInt()

	for _, activeRedelegation := range keepers.StakingKeeper.GetRedelegations(ctx, accAddr, 65535) {
		redelegationSrc, _ := sdk.ValAddressFromBech32(activeRedelegation.ValidatorSrcAddress)
		redelegationDst, _ := sdk.ValAddressFromBech32(activeRedelegation.ValidatorDstAddress)

		// set all entry completionTime to now so we can complete re-delegation
		for i := range activeRedelegation.Entries {
			activeRedelegation.Entries[i].CompletionTime = now
			redelegatedAmt = redelegatedAmt.Add(math.Int(activeRedelegation.Entries[i].SharesDst))
		}

		keepers.StakingKeeper.SetRedelegation(ctx, activeRedelegation)
		_, err := keepers.StakingKeeper.CompleteRedelegation(ctx, accAddr, redelegationSrc, redelegationDst)
		if err != nil {
			return redelegatedAmt, err
		}
	}

	return redelegatedAmt, nil
}

// Returns the amount of tokens which were unbonded (not rewards)
func unbondAllAndFinish(ctx sdk.Context, now time.Time, keepers *keepers.AppKeepers, accAddr sdk.AccAddress) (math.Int, error) {
	unbondedAmt := math.ZeroInt()

	// Unbond all delegations from the account
	for _, delegation := range keepers.StakingKeeper.GetAllDelegatorDelegations(ctx, accAddr) {
		validatorValAddr := delegation.GetValidatorAddr()
		_, found := keepers.StakingKeeper.GetValidator(ctx, validatorValAddr)
		if !found {
			continue
		}

		_, err := keepers.StakingKeeper.Undelegate(ctx, accAddr, validatorValAddr, delegation.GetShares())
		if err != nil {
			return math.ZeroInt(), err
		}
	}

	// Take all unbonding and complete them.
	for _, unbondingDelegation := range keepers.StakingKeeper.GetAllUnbondingDelegations(ctx, accAddr) {
		validatorStringAddr := unbondingDelegation.ValidatorAddress
		validatorValAddr, _ := sdk.ValAddressFromBech32(validatorStringAddr)

		// Complete unbonding delegation
		for i := range unbondingDelegation.Entries {
			unbondingDelegation.Entries[i].CompletionTime = now
			unbondedAmt = unbondedAmt.Add(unbondingDelegation.Entries[i].Balance)
		}

		keepers.StakingKeeper.SetUnbondingDelegation(ctx, unbondingDelegation)
		_, err := keepers.StakingKeeper.CompleteUnbonding(ctx, accAddr, validatorValAddr)
		if err != nil {
			return math.ZeroInt(), err
		}
	}

	return unbondedAmt, nil
}