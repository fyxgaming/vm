package types

import (
	"log"

	"github.com/fyxgaming/vm/fighter/types/skills"
)

type Race int32

const (
	Human Race = 0
	Dwarf Race = 1
	Elf   Race = 2
)

type Gender int32

const (
	Male   Gender = 0
	Female Gender = 1
)

var InitialScores = [][]int32{{8, 8, 8}, {9, 7, 7}, {7, 7, 7}}
var InitialSkills = []skills.SkillType{
	skills.Human,
	skills.Dwarven,
	skills.Elven,
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

var SkillLevels map[int32]bool = map[int32]bool{
	2:  true,
	4:  true,
	6:  true,
	8:  true,
	10: true,
	16: true,
	20: true,
}

type Appearance struct {
	Gender       int32 `json:"gender"`
	FaceId       int32 `json:"faceIdentifier"`
	HairId       int32 `json:"hairIdentifier"`
	BeardId      int32 `json:"beardIdentifier"`
	SkinToneId   int32 `json:"skinToneIdentifier"`
	HairColorId  int32 `json:"hairColorIdentifier"`
	EyeColorId   int32 `json:"eyeColorIdentifier"`
	ArmorColorId int32 `json:"armorColorIdentifier"`
	HatId        int32 `json:"hatIdentifier"`
}

type Fighter struct {
	Name               string             `json:"name"`
	Race               Race               `json:"race"`
	AbilityScores      []int32            `json:"abilityScores"`
	Appearance         Appearance         `json:"appearance"`
	Skills             []skills.SkillType `json:"skills"`
	Level              int32              `json:"level"`
	Xp                 int32              `json:"xp"`
	LevelUpHpBonusDice []int32            `json:"levelUpHpBonusDice"`
	LevelUpHpBonusMods []int32            `json:"levelUpHpBonusModifiers"`
	PvpBattlesFought   int32              `json:"pvpBattlesFought"`
	PvpBattlesWon      int32              `json:"pvpBattlesWon"`
	PveBattlesFought   int32              `json:"pveBattlesFought"`
	PveBattlesWon      int32              `json:"pveBattlesWon"`
}

var Costs = []int32{0, 1, 2, 4, 7, 11, 16}

func (f *Fighter) ValidateStartingScores() bool {
	cost := int32(0)
	for ability, score := range f.AbilityScores {
		start := InitialScores[f.Race][ability]
		upgrade := score - start
		if upgrade < 0 {
			log.Printf("Invalid Ability Scores")
			return false
		}
		if upgrade > 0 {
			points := Costs[upgrade]
			if points <= 0 {
				log.Printf("Invalid Ability Scores")
				return false
			}
			cost += points
		}
	}
	return cost == 15
}

type LevelUpRequest struct {
	Ability   int32            `json:"ability"`
	SkillType skills.SkillType `json:"skillType"`
}

type Ability int32

const (
	Strength     Ability = 0
	Dexterity    Ability = 1
	Intelligence Ability = 2
)

var AbilityLevels map[int32]bool = map[int32]bool{
	3:  true,
	5:  true,
	7:  true,
	9:  true,
	14: true,
	18: true,
}
