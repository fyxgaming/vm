package types

import (
	"crypto/sha256"

	"github.com/fyxgaming/vm/lib"
)

var LevelCap int32 = 10

var XPGains = []int32{
	0,
	255,
	340,
	380,
	420,
	460,
	505,
	545,
	585,
	625,
	670,
	710,
	750,
	790,
	835,
	875,
	915,
	955,
	1000,
	1040,
	1080,
}

var XPTable = []int32{
	0,
	250,
	330,
	435,
	575,
	760,
	1005,
	1330,
	1760,
	2330,
	3085,
	4085,
	5405,
	7155,
	9470,
	12535,
	16590,
	21960,
	29065,
	29065,
}

var XPModifiers = map[int32]float32{
	-2: 0,
	-1: 0.5,
	0:  1,
	1:  1.5,
	2:  2,
}

type BattleStart struct {
	Engine  []byte
	Players []*Player
}

type Player struct {
	Pubkey  []byte
	Payload []byte
	IsBot   bool
}

type Battle struct {
	Engine  []byte
	Players []*BattlePlayer
	Turns   []*BattleTurn
	Status  BattleStatus
}

type BattlePlayer struct {
	Pubkey  []byte
	Payload []byte
	Fighter *lib.Instance
	Items   []*lib.Instance
	Xp      int32
	Level   int32
	IsBot   bool
}

type BattleStatus int8

const (
	Open       BattleStatus = 0
	Conceded   BattleStatus = 1
	Challenged BattleStatus = -1
)

type BattleTurn struct {
	Seq     int32
	Payload []byte
	Status  BattleStatus
	Sig     []byte
}

func (b *BattleTurn) Hash() []byte {
	msg := []byte{byte(b.Status)}
	msg = append(msg, b.Payload...)

	hash := sha256.Sum256(msg)
	return hash[:]
}

type BattleEnd struct {
	Turns []*BattleTurn
}
