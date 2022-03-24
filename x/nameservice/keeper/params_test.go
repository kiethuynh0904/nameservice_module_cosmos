package keeper_test

import (
	"testing"

	testkeeper "github.com/kiethuynh0904/nameservice/testutil/keeper"
	"github.com/kiethuynh0904/nameservice/x/nameservice/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NameserviceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
