package v2_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CosmosContracts/juno/v16/x/globalfee"
	"github.com/CosmosContracts/juno/v16/x/globalfee/exported"
	v2 "github.com/CosmosContracts/juno/v16/x/globalfee/migrations/v2"
	"github.com/CosmosContracts/juno/v16/x/globalfee/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps exported.ParamSet) {
	*ps.(*types.Params) = ms.ps
}

func TestMigrate(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(globalfee.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := sdk.NewKVStoreKey(v2.ModuleName)
	tKey := sdk.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	store := ctx.KVStore(storeKey)

	legacySubspace := newMockSubspace(types.DefaultParams())
	require.NoError(t, v2.Migrate(ctx, store, legacySubspace, cdc))

	var res types.Params
	bz := store.Get(v2.ParamsKey)
	require.NoError(t, cdc.Unmarshal(bz, &res))
	require.Equal(t, legacySubspace.ps, res)
}