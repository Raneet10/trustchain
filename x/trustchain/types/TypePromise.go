package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Promise struct {
	Creator            sdk.AccAddress `json:"creator" yaml:"creator"`
	ID                 string         `json:"id" yaml:"id"`
	PromiseDescription string         `json:"promiseDescription" yaml:"promiseDescription"`
	PromiseKeeper      sdk.AccAddress `json:"promiseKeeper" yaml:"promiseKeeper"`
	Reward             sdk.Coins      `json:"reward" yaml:"reward"`
	Deadline           time.Time      `json:"deadline" yaml:"deadline"`
	Confirmed          bool           `json:"confirmed" yaml:"confirmed"`
	Kept               bool           `json:"kept" yaml:"kept"`
}
