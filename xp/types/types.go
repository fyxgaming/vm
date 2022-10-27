package types

import "github.com/fyxgaming/vm/lib"

type XpGrant struct {
	Fighter lib.Outpoint `json:"fighter"`
	Xp      int32        `json:"xp"`
	Lock    []byte       `json:"lock"`
	Win     bool         `json:"win"`
	PvP     bool         `json:"pvp"`
}
