package types

import "github.com/fyxgaming/vm/lib"

type BattleReceipt struct {
	// Battle  lib.Outpoint `json:"battle"`
	Fighter lib.Outpoint `json:"fighter"`
	Xp      int32        `json:"xp"`
	Lock    []byte       `json:"lock"`
	Win     bool         `json:"win"`
	PvP     bool         `json:"pvp"`
}
