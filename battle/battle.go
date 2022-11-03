package main

import (
	"encoding/hex"
	"fmt"

	"github.com/fyxgaming/vm/refs"
	"github.com/fyxgaming/vm/battle/types"
	fighter "github.com/fyxgaming/vm/fighter/types"
	"github.com/fyxgaming/vm/lib"
	receipt "github.com/fyxgaming/vm/receipt/types"
	"github.com/libsv/go-bk/bec"
	"github.com/mailru/easyjson"
)

// Populate with Fighter/Item Contracts
const RECEIPT string = ""   //FYX:CONTRACT:cryptofights/receipt
const FIGHTER string = refs.Refs["cryptofights/fighter"]
const ITEM string = ""      //FYX:CONTRACT:cryptofights/item
// const VALIDATOR string = "" //FYX:PUBKEY:validator/cryptofights

func main() {}

func Init() int {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	var start types.BattleStart
	err = easyjson.Unmarshal([]byte(this.CallData), &start)
	if err != nil {
		return this.Return(err)
	}

	battle := types.Battle{
		Engine: start.Engine,
	}
	var player *types.BattlePlayer
	for _, e := range this.Stack {
		if e.Instance.Kind.String() == FIGHTER {
			if len(battle.Players) > len(start.Players) {
				return this.Return(fmt.Errorf("excess-fighters"))
			}
			var fighter *fighter.Fighter
			err = easyjson.Unmarshal([]byte(e.Instance.Storage), fighter)
			if err != nil {
				return this.Return(err)
			}
			player = &types.BattlePlayer{
				Pubkey:  start.Players[len(battle.Players)].Pubkey,
				Payload: start.Players[len(battle.Players)].Payload,
				Fighter: e.Instance,
				Xp:      fighter.Xp,
			}
			battle.Players = append(battle.Players, player)
		} else if e.Instance.Kind.String() == ITEM {
			player.Items = append(player.Items, e.Instance)
		}
	}

	battleData, err := battle.MarshalJSON()
	if err != nil {
		return this.Return(err)
	}
	this.Instance.Satoshis = 1
	this.Instance.Storage = string(battleData)
	lock, err := hex.DecodeString(VALIDATOR)
	if err != nil {
		return this.Return(err)
	}

	this.Instance.Lock = lock

	return this.Return(nil)
}

//export Resolve
func Resolve() int {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	battleEnd := &types.BattleEnd{}
	err = easyjson.Unmarshal([]byte(this.CallData), battleEnd)
	if err != nil {
		return this.Return(err)
	}
	if len(battleEnd.Turns) == 0 {
		return this.Return(fmt.Errorf("empty-battle"))
	}

	battle := &types.Battle{}
	err = easyjson.Unmarshal([]byte(this.Instance.Storage), battle)
	if err != nil {
		return this.Return(err)
	}

	battle.Turns = battleEnd.Turns
	last := battleEnd.Turns[len(battleEnd.Turns)-1]
	var winner *types.BattlePlayer
	var loser *types.BattlePlayer
	var player *types.BattlePlayer

	for i, p := range battle.Players {
		player = p
		sig, err := bec.ParseSignature(last.Sig, bec.S256())
		if err != nil {
			return this.Return(fmt.Errorf("malformed-sig"))
		}
		pub, err := bec.ParsePubKey(p.Pubkey, bec.S256())
		if err != nil {
			return this.Return(err)
		}

		if sig.Verify(last.Hash(), pub) {
			switch last.Status {
			case types.Conceded:
				loser = p
				if i == 0 {
					winner = battle.Players[1]
				} else {
					winner = battle.Players[0]
				}
				contract, err := lib.NewOutpointFromString(RECEIPT)
				if err != nil {
					return this.Return(err)
				}
				var xp int32
				if winner.
				receipt := &receipt.BattleReceipt{
					Fighter: *winner.Fighter,
					Xp:      0,
					Win:     true,
					PvP:     len(battle.Players) > 1,
				}

				data, err := receipt.MarshalJSON()
				this.Spawn(contract, "Grant", string(data))

			case types.Challenged:
			default:
				return this.Return(fmt.Errorf("battle-not-complete"))
			}
			break
		}
	}
	if player == nil {
		return this.Return(fmt.Errorf("invalid-resolve"))
	}

	return this.Return(nil)
}
