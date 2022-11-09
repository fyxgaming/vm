package main

import (
	"encoding/hex"
	"fmt"

	"github.com/fyxgaming/vm/battle/types"
	fighter "github.com/fyxgaming/vm/fighter/types"
	item "github.com/fyxgaming/vm/item/types"
	"github.com/fyxgaming/vm/lib"
	receipt "github.com/fyxgaming/vm/receipt/types"
	"github.com/fyxgaming/vm/refs"
	"github.com/libsv/go-bk/bec"
	"github.com/mailru/easyjson"
)

// Populate with Fighter/Item Contracts
const VALIDATOR_LOCK = "" //FYX:PUBKEY:

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
		if e.Instance.Kind.String() == refs.Refs["cryptofights/fighter"] {
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
				IsBot:   start.Players[len(battle.Players)].IsBot,
				Fighter: e.Instance,
				Xp:      fighter.Xp,
				Level:   fighter.Level,
			}
			battle.Players = append(battle.Players, player)
		} else if e.Instance.Kind.String() == refs.Refs["cryptofights/item"] {
			player.Items = append(player.Items, e.Instance)
		}
	}

	topics := [][]byte{}
	for _, player := range battle.Players {
		topics = append(topics, player.Pubkey)
	}

	this.Emit("BattleCreated", topics)
	battleData, err := battle.MarshalJSON()
	if err != nil {
		return this.Return(err)
	}
	this.Instance.Satoshis = 1
	this.Instance.Storage = battleData
	lock, err := hex.DecodeString(VALIDATOR_LOCK)
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
	var victor *types.BattlePlayer
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
					victor = battle.Players[1]
				} else {
					victor = battle.Players[0]
				}
				contract, err := lib.NewOutpointFromString(refs.Refs["cryptofights/receipt"])
				if err != nil {
					return this.Return(err)
				}

				var xp int32
				if victor.Level < types.LevelCap && loser.IsBot {
					baseXp := float32(types.XPGains[loser.Level])
					diff := loser.Level - victor.Level

					if victor.Level <= types.XPGains[victor.Level] {
						xp = int32(baseXp + types.XPModifiers[0])
					} else if victor.Level > loser.Level {
						if diff <= -2 {
							xp = int32(baseXp + types.XPModifiers[-2])
						} else {
							xp = int32(baseXp + types.XPModifiers[-1])
						}
					} else {
						if diff > 2 {
							xp = int32(baseXp + types.XPModifiers[2])
						} else {
							xp = int32(baseXp + types.XPModifiers[1])
						}
					}
				}

				if !victor.IsBot {
					// Do item reward stuff
					itemContract, err := lib.NewOutpointFromString(refs.Refs["cryptofights/item"])
					if err != nil {
						return this.Return(err)
					}
					items := []*item.ItemReq{}
					for _, item := range items {
						itemData, err := item.MarshalJSON()
						if err != nil {
							return this.Return(err)
						}
						this.Spawn(itemContract, "New", itemData)
					}
				}

				rec := &receipt.BattleReceipt{
					Fighter: *victor.Fighter.Origin,
					Xp:      xp,
					Win:     true,
					PvP:     len(battle.Players) > 1,
				}
				data, err := rec.MarshalJSON()
				if err != nil {
					return this.Return(err)
				}
				this.Spawn(contract, "Grant", data)

				rec = &receipt.BattleReceipt{
					Fighter: *loser.Fighter.Origin,
					PvP:     len(battle.Players) > 1,
				}
				data, err = rec.MarshalJSON()
				if err != nil {
					return this.Return(err)
				}
				this.Spawn(contract, "Grant", data)

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
