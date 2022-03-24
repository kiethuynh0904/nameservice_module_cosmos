package keeper

import (
	"github.com/kiethuynh0904/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
