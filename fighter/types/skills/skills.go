package skills

type SkillType int32

const (
	Attack              SkillType = 0
	Focus               SkillType = 1
	Hide                SkillType = 2
	Burn                SkillType = 3
	Chill               SkillType = 4
	CoatWeapon          SkillType = 5
	DirtyFighting       SkillType = 6
	Heal                SkillType = 7
	Hex                 SkillType = 8
	Mark                SkillType = 9
	Meditate            SkillType = 10
	PhaseShift          SkillType = 11
	PinningStrike       SkillType = 12
	PowerAttack         SkillType = 13
	Shock               SkillType = 14
	Smite               SkillType = 15
	SneakAttack         SkillType = 16
	Stun                SkillType = 17
	Taunt               SkillType = 18
	BattleCry           SkillType = 19
	Cleanse             SkillType = 20
	Curse               SkillType = 21
	Fireball            SkillType = 22
	Freeze              SkillType = 23
	Gore                SkillType = 24
	LightningBolt       SkillType = 25
	LockAndLoad         SkillType = 26
	SecondWind          SkillType = 27
	Elementalist        SkillType = 28
	Runesmith           SkillType = 29
	Sharpshooter        SkillType = 30
	Dwarven             SkillType = 31
	Elven               SkillType = 32
	Human               SkillType = 33
	Ambidextrous        SkillType = 34
	Armored             SkillType = 35
	Cirurgical          SkillType = 36
	Durable             SkillType = 37
	Strong              SkillType = 38
	Unarmed             SkillType = 39
	ImprovedArmored     SkillType = 40
	ImprovedPowerAttack SkillType = 41
	ImprovedSneakAttack SkillType = 42
	Opportunistic       SkillType = 43
	Precise             SkillType = 44
	Vengeful            SkillType = 45
	Warmongering        SkillType = 46
	Slippery            SkillType = 47
	Scoundrel           SkillType = 58
	Shadowy             SkillType = 59
	Stab                SkillType = 60
	Strike              SkillType = 61
	Charge              SkillType = 62
	MundaneProtection   SkillType = 63
	ReTrain             SkillType = 64
)

type Rank int32

const (
	Action Rank = 0
	Talent Rank = 1
	Trait  Rank = 2
)

type Skill struct {
	Id               string
	RequiredLevel    int32
	RequiredSkills   []SkillType
	UnlockableSkills []SkillType
	Rank             Rank
	AffectsOpponent  bool
	IsAttackAction   bool
	IsSpellAction    bool
	Special          bool
}

var SkillData map[SkillType]*Skill = map[SkillType]*Skill{
	Dwarven: {Id: "dwarven", Rank: Trait},
	Elven:   {Id: "elven", Rank: Trait},
	Human:   {Id: "human", Rank: Trait},
	Attack: {
		Id:              "attack",
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	Focus: {Id: "focus"},
	Hide:  {Id: "hide"},
	Burn: {
		Id:              "burn",
		RequiredSkills:  []SkillType{Elementalist},
		AffectsOpponent: true,
		IsAttackAction:  true,
		IsSpellAction:   true,
	},
	Chill: {
		Id:              "chill",
		RequiredSkills:  []SkillType{Elementalist},
		AffectsOpponent: true,
		IsAttackAction:  true,
		IsSpellAction:   true,
	},
	CoatWeapon: {Id: "coat-weapon"},
	DirtyFighting: {
		Id:              "dirty-fighting",
		RequiredSkills:  []SkillType{Scoundrel},
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	Heal: {
		Id:             "heal",
		RequiredSkills: []SkillType{Cleanse},
		IsSpellAction:  true,
	},
	Hex: {
		Id:              "hex",
		AffectsOpponent: true,
		IsSpellAction:   true,
	},
	Mark: {
		Id:              "mark",
		AffectsOpponent: true,
	},
	Meditate: {Id: "meditate"},
	PhaseShift: {
		Id:            "phase-shift",
		IsSpellAction: true,
	},
	PinningStrike: {
		Id:              "pinning-strike",
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	PowerAttack: {
		Id:              "power-attack",
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	Shock: {
		Id:              "shock",
		RequiredSkills:  []SkillType{Elementalist},
		AffectsOpponent: true,
		IsAttackAction:  true,
		IsSpellAction:   true,
	},
	Smite: {
		Id:              "smite",
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	SneakAttack: {
		Id:              "sneak-attack",
		RequiredSkills:  []SkillType{Scoundrel},
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	Stun: {
		Id:              "stun",
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	Taunt: {
		Id:              "taunt",
		AffectsOpponent: true,
	},
	BattleCry: {
		Id:              "battle-cry",
		RequiredLevel:   3,
		AffectsOpponent: true,
	},
	Cleanse: {Id: "clense"},
	Curse: {
		Id:              "curse",
		RequiredSkills:  []SkillType{Hex},
		AffectsOpponent: true,
		IsSpellAction:   true,
	},
	Fireball: {
		Id:              "fireball",
		RequiredSkills:  []SkillType{Elementalist},
		AffectsOpponent: true,
		IsAttackAction:  true,
		IsSpellAction:   true,
	},
	Freeze: {
		Id:              "freeze",
		RequiredSkills:  []SkillType{Elementalist},
		AffectsOpponent: true,
		IsAttackAction:  true,
		IsSpellAction:   true,
	},
	Gore: {
		Id:              "gore",
		RequiredSkills:  []SkillType{PinningStrike},
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	LightningBolt: {
		Id:              "lightning-bolt",
		RequiredSkills:  []SkillType{Elementalist},
		AffectsOpponent: true,
		IsAttackAction:  true,
		IsSpellAction:   true,
	},
	LockAndLoad: {
		Id:              "lock-and-load",
		RequiredSkills:  []SkillType{Sharpshooter},
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	SecondWind: {
		Id:             "second-wind",
		RequiredSkills: []SkillType{Durable},
	},
	//Talents
	Elementalist: {
		Id:               "elementalist",
		UnlockableSkills: []SkillType{Burn, Chill, Shock},
		Rank:             Talent,
	},
	Runesmith: {
		Id:               "runesmith",
		RequiredSkills:   []SkillType{Armored},
		UnlockableSkills: []SkillType{ImprovedArmored, Warmongering},
		Rank:             Talent,
	},
	Sharpshooter: {
		Id:               "sharpshooter",
		RequiredSkills:   []SkillType{Cirurgical},
		UnlockableSkills: []SkillType{LockAndLoad, Precise},
		Rank:             Talent,
	},
	//Traits
	Ambidextrous: {
		Id:   "ambidextrous",
		Rank: Trait,
	},
	Armored: {
		Id:   "armored",
		Rank: Trait,
	},
	Cirurgical: {
		Id:   "cirurgical",
		Rank: Trait,
	},
	Durable: {
		Id:   "durable",
		Rank: Trait,
	},
	Strong: {
		Id:   "strong",
		Rank: Trait,
	},
	Unarmed: {
		Id:   "unarmed",
		Rank: Trait,
	},
	ImprovedArmored: {
		Id:             "improved-armored",
		RequiredSkills: []SkillType{Runesmith},
		Rank:           Trait,
	},
	ImprovedPowerAttack: {
		Id:             "improved-power-attack",
		RequiredSkills: []SkillType{PowerAttack},
		Rank:           Trait,
	},
	ImprovedSneakAttack: {
		Id:             "improved-sneak-attack",
		RequiredSkills: []SkillType{SneakAttack},
		Rank:           Trait,
	},
	Opportunistic: {
		Id:            "opportunistic",
		RequiredLevel: 3,
		Rank:          Trait,
	},
	Precise: {
		Id:             "precise",
		RequiredSkills: []SkillType{Sharpshooter},
		Rank:           Trait,
	},
	Vengeful: {
		Id:            "vengeful",
		RequiredLevel: 3,
		Rank:          Trait,
	},
	Warmongering: {
		Id:             "warmongering",
		RequiredSkills: []SkillType{Runesmith},
		Rank:           Trait,
	},
	Slippery: {
		Id:             "slippery",
		RequiredSkills: []SkillType{Runesmith},
		Rank:           Trait,
	},
	Scoundrel: {
		Id:               "scoundrel",
		UnlockableSkills: []SkillType{Shadowy, DirtyFighting, SneakAttack},
		Rank:             Talent,
	},
	Shadowy: {
		Id:             "shadowy",
		RequiredSkills: []SkillType{Scoundrel},
		Rank:           Trait,
	},
	Strike: {
		Id:              "strike",
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	Stab: {
		Id:              "stab",
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	Charge: {
		Id:              "charge",
		AffectsOpponent: true,
		IsAttackAction:  true,
	},
	MundaneProtection: {
		Id:            "mundane-protection",
		RequiredLevel: 3,
	},
	ReTrain: {
		Id:             "re-train",
		RequiredLevel:  7,
		RequiredSkills: []SkillType{Human},
		Special:        true,
	},
}
