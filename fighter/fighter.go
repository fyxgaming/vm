package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/fyxgaming/vm/fighter/types"
	"github.com/fyxgaming/vm/fighter/types/skills"
	"github.com/fyxgaming/vm/lib"
	"github.com/mailru/easyjson"
	"golang.org/x/exp/slices"
)

// Populate with Xp Contract
const Xp string = ""

var Validator []byte = []byte{}

func main() {}

func New() (retCode int) {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	fighter := types.Fighter{}
	err = easyjson.Unmarshal([]byte(this.CallData), &fighter)
	if err != nil {
		return this.Return(err)
	}

	if ok := fighter.ValidateStartingScores(); !ok {
		return this.Return(fmt.Errorf("invalid-scores"))
	}

	fighter.Level = 1

	fighter.Skills = []skills.SkillType{
		skills.Attack,
		skills.Focus,
		skills.Hide,
		types.InitialSkills[fighter.Race],
	}
	fighter.LevelUpHpBonusDice = make([]int32, 24)
	fighter.LevelUpHpBonusMods = make([]int32, 24)

	return this.Return(nil)
}

func RecordBattle() (retCode int) {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	fighter := types.Fighter{}
	err = easyjson.Unmarshal([]byte(this.Instance.Storage), &fighter)
	if err != nil {
		return this.Return(err)
	}

	var xp int64
	for _, exec := range this.Stack {
		for _, e := range exec.Events {
			if e.Id != "claim-xp" ||
				exec.Contract.String() != Xp ||
				e.Topics[0] != this.Instance.Outpoint.String() ||
				!bytes.Equal(Validator, exec.Parent.Lock) {

				continue
			}

			xp, err = strconv.ParseInt(e.Topics[1], 10, 32)
			if err != nil {
				return this.Return(err)
			}
			if fighter.Level < 10 {
				fighter.Xp += int32(xp)
			}

			win, err := strconv.ParseBool(e.Topics[2])
			if err != nil {
				return this.Return(err)
			}

			pvp, err := strconv.ParseBool(e.Topics[3])
			if err != nil {
				return this.Return(err)
			}

			if pvp {
				fighter.PvpBattlesFought++
				if win {
					fighter.PvpBattlesWon++
				}
			} else {
				fighter.PveBattlesFought++
				if win {
					fighter.PveBattlesWon++
				}
			}
		}
	}

	return this.Return(nil)
}

//export LevelUp
func LevelUp() (retCode int) {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	req := types.LevelUpRequest{}
	err = easyjson.Unmarshal([]byte(this.CallData), &req)
	if err != nil {
		return this.Return(err)
	}

	fighter := types.Fighter{}
	err = easyjson.Unmarshal([]byte(this.Instance.Storage), &fighter)
	if err != nil {
		return this.Return(err)
	}

	reqXp := types.XPTable[fighter.Level]
	if fighter.Xp < reqXp {
		return this.Return(fmt.Errorf("inadequate-xp"))
	}

	fighter.Level++
	if fighter.Level == 7 && fighter.Race == types.Human {
		fighter.Skills = append(fighter.Skills, skills.ReTrain)
	}
	fighter.Xp -= reqXp

	if _, ok := types.SkillLevels[fighter.Level]; ok {
		if req.SkillType == 0 {
			return this.Return(fmt.Errorf("missing-skill"))
		}
		skillInfo, ok := skills.SkillData[req.SkillType]
		if !ok {
			return this.Return(fmt.Errorf("invalid-skill"))
		}
		if fighter.Level-1 < skillInfo.RequiredLevel {
			return this.Return(fmt.Errorf("insufficient-level"))
		}
		for _, reqSkill := range skillInfo.RequiredSkills {
			if slices.Contains(fighter.Skills, reqSkill)
			found := false
			for _, skill := range fighter.Skills {
				if reqSkill == skill {
					found = true
					break
				}
			}
			if !found {
				return this.Return(fmt.Errorf("missing-required-skill"))
			}
		}

		fighter.Skills = append(fighter.Skills, req.SkillType)
		for _, skill := range skillInfo.UnlockableSkills {
			fighter.Skills = append(fighter.Skills, skill)
		}
	}

	if _, ok := types.AbilityLevels[fighter.Level]; ok {
		fighter.AbilityScores[req.Ability]++
	}

	strMod := (fighter.AbilityScores[0] - 10) / 2

	//TODO: implement Dice rolls
	fighter.LevelUpHpBonusMods[fighter.Level] = strMod

	return this.Return(nil)
}
