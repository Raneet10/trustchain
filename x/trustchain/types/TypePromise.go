package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var TrustCoin = sdk.Coins{sdk.NewInt64Coin("trustcoin", 1)}

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

func (p Promise) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
ID: %s
PromiseDescription: %s
PromiseKeeper: %s
Reward: %s
Deadline: %s
Confirmed: %t
Kept: %t`, p.Creator, p.ID, p.PromiseDescription, p.PromiseKeeper, p.Reward, p.Deadline, p.Confirmed, p.Kept))
}
