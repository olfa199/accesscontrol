package keeper_test

import (
	"testing"

	keepertest "crossdomain/testutil/keeper"
	"crossdomain/testutil/nullify"
	"crossdomain/x/cdaccesscontrol/keeper"
	"crossdomain/x/cdaccesscontrol/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNCertificate(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Certificate {
	items := make([]types.Certificate, n)
	for i := range items {
		items[i].Id = keeper.AppendCertificate(ctx, items[i])
	}
	return items
}

func TestCertificateGet(t *testing.T) {
	keeper, ctx := keepertest.CdaccesscontrolKeeper(t)
	items := createNCertificate(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCertificate(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCertificateRemove(t *testing.T) {
	keeper, ctx := keepertest.CdaccesscontrolKeeper(t)
	items := createNCertificate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCertificate(ctx, item.Id)
		_, found := keeper.GetCertificate(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCertificateGetAll(t *testing.T) {
	keeper, ctx := keepertest.CdaccesscontrolKeeper(t)
	items := createNCertificate(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCertificate(ctx)),
	)
}

func TestCertificateCount(t *testing.T) {
	keeper, ctx := keepertest.CdaccesscontrolKeeper(t)
	items := createNCertificate(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCertificateCount(ctx))
}
