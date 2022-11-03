package types

import (
	"crypto/sha256"

	"github.com/fyxgaming/vm/lib"
)

// type StatusEffect int32

// var (
//
//	Cursed                        StatusEffect = 0
//	Ethereal                      StatusEffect = 1
//	Focused                       StatusEffect = 2
//	Hexed                         StatusEffect = 3
//	Hidden                        StatusEffect = 4
//	Inspired                      StatusEffect = 5
//	Marked                        StatusEffect = 6
//	Meditative                    StatusEffect = 7
//	Poisoned                      StatusEffect = 8
//	Stunned                       StatusEffect = 9
//	Taunted                       StatusEffect = 10
//	Poisonous                     StatusEffect = 11
//	Demoralized                   StatusEffect = 12
//	Swifted                       StatusEffect = 13
//	Burning                       StatusEffect = 14
//	ProtectedAgainstMundaneDamage StatusEffect = 15
//	Electrocuted                  StatusEffect = 16
//	Freezing                      StatusEffect = 17
//
// )

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
