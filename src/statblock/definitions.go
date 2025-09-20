package statblock

type Attack struct {
	DamageType DamageType `json:"dmgType"`
	Damage     DiceRoll   `json:"damage"`
	Conditions []string   `json:"conditions"`
}

type Ability struct {
	Id          string    `json:"id"`
	UseTime     Duration  `json:"useTime"`
	Duration    *DiceRoll `json:"duration"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Conditions  []string  `json:"conditions"`
	Attacks     []Attack  `json:"attacks"`
	Effects     []Effect  `json:"effects"`
	Indefinite  bool      `json:"indefinite"`
}

type Feat struct {
	Id            string     `json:"id"`
	Name          string     `json:"name"`
	Buffs         []FeatBuff `json:"buffs"`
	Abilities     []Ability  `json:"abilities"`
	Description   string     `json:"description"`
	Prerequisites []string   `json:"prereqs"`
	ImageUrl      string     `json:"imgUrl"`
}

type FeatBuff struct {
	Stat       Stat     `json:"stat"`
	Value      string   `json:"value"`
	Conditions []string `json:"conditions"`
}

var (
	AbilityJuke         Ability            = Ability{"juke", Action, nil, "Whirl past an opponent, swapping places with it", "You use an action to swap places with a creature within 5ft of you that is your size or smaller. You and your allies will not provoke opportunity attacks from the target until your next turn.", []string{"1 action", "target is within 5 feet", "target is creature's size or smaller"}, []Attack{}, []Effect{}, false}
	AbilityFlurry       Ability            = Ability{"flurry", Action, nil, "Make an attack action with every hand at once, including hands granted by mutation or technology", "Once per encounter, you may expend an action to make an attack using every hand you have. For the purposes of other abilities, these attacks count as discrete attack actions.", []string{"1 action", "target is in melee range", "once per encounter"}, []Attack{}, []Effect{}, false}
	AbilityCharge       Ability            = Ability{"charge", Action, nil, "Perform a melee attack after charging between 10-20ft forward", "Once per encounter, you can charge between 10-20ft towards an enemy of your choosing, making an attack with your primary weapon with +1 to-hit.", []string{"1 action", "target is between 10 and 20 feet", "once per encounter"}, []Attack{}, []Effect{}, false}
	AbilityExtremeSpeed Ability            = Ability{"extreme speed", "", nil, "Take two turns each round of combat", "When you enter combat, roll initiative twice. Use the highest roll as your first turn and the lowest as your second. Abilities that may be used every turn can be used in both turns.", []string{"must be an enemy", "150 quickness or greater"}, []Attack{}, []Effect{}, true}
	AbilityBludgeon     Ability            = Ability{"bludgeon", Action, nil, "Make an attack with a cudgel, dazing an opponent", "When you attack with a cudgel, roll a d4. On a 4, your attack inflicts Dazed on your opponent. If your opponent is already Dazed you instead Stun them for 1 round.", []string{"1 action", "target is in melee range"}, []Attack{}, []Effect{{DAZED, []string{"opponent is not already dazed", "4 on 1d4 to Daze", "attacking with a cudgel"}, []string{}, DiceRoll{[]string{"1d4"}, 0, StatNone}}, {STUNNED, []string{"opponent is already dazed", "4 on 1d4 to Daze", "attacking with a cudgel"}, []string{}, DiceRoll{[]string{}, 1, StatNone}}}, false}
	Abilities           map[string]Ability = map[string]Ability{
		AbilityJuke.Id:         AbilityJuke,
		AbilityFlurry.Id:       AbilityFlurry,
		AbilityCharge.Id:       AbilityCharge,
		AbilityExtremeSpeed.Id: AbilityExtremeSpeed,
		AbilityBludgeon.Id:     AbilityBludgeon,
	}
)

var (
	FeatSprint                 Feat            = Feat{"sprint", "Sprint", []FeatBuff{{Speed, "10", []string{}}}, []Ability{}, "This creature is capable of moving quickly", []string{}, "https://wiki.cavesofqud.com/images/4/41/Sprint_skill.png"}
	FeatSwiftReflexes          Feat            = Feat{"swift reflexes", "Swift Reflexes", []FeatBuff{{AC, "2", []string{"As a reaction to a projectile attack that targets the creature"}}}, []Ability{}, "This creature has a +2 AC bonus when flinching away from projectile attacks", []string{}, "https://wiki.cavesofqud.com/images/6/6d/Swift_reflexes_skill.png"}
	FeatSpry                   Feat            = Feat{"spry", "Spry", []FeatBuff{{DEX, "3", []string{}}}, []Ability{}, "This creature is nimble and gains a +3 bonus to its DEX", []string{}, "https://wiki.cavesofqud.com/images/b/b8/Spry_skill.png"}
	FeatTumble                 Feat            = Feat{"tumble", "Tumble", []FeatBuff{{DEX, "1", []string{}}}, []Ability{}, "This creature gains a +1 bonus to its DEX and can Juke as a bonus action if it has the feat", []string{}, "https://wiki.cavesofqud.com/images/6/62/Tumble_skill.png"}
	FeatJuke                   Feat            = Feat{"juke", "Juke", []FeatBuff{}, []Ability{AbilityJuke}, "This creature can whirl past an opponent, swapping places with it", []string{}, "https://wiki.cavesofqud.com/images/c/ce/Juke_skill.png"}
	FeatAxeProficiency         Feat            = Feat{"axe proficiency", "Axe Proficiency", []FeatBuff{{TOHIT, "proficiency", []string{"using an axe"}}}, []Ability{}, "This creature is proficient with axes", []string{}, "https://wiki.cavesofqud.com/images/6/6b/Axe_proficiency_skill.png"}
	FeatSteadyHands            Feat            = Feat{"steady hands", "Bow and Rifle Proficiency", []FeatBuff{{TOHIT, "proficiency", []string{"using a bow or rifle"}}}, []Ability{}, "This creature is skilled with bows and rifles", []string{}, "https://wiki.cavesofqud.com/images/3/37/Steady_hands_skill.png"}
	FeatCudgelProficiency      Feat            = Feat{"cudgel proficiency", "Cudgel Proficiency", []FeatBuff{{TOHIT, "proficiency", []string{"using a cudgel"}}}, []Ability{}, "This creature is skilled with crushing and bludgeoning weapons", []string{}, "https://wiki.cavesofqud.com/images/7/7d/Cudgel_proficiency_skill.png"}
	FeatLongBladeProficiency   Feat            = Feat{"long blade proficiency", "Long Blade Proficiency", []FeatBuff{{TOHIT, "proficiency", []string{"using a long blade"}}}, []Ability{}, "This creature is skilled with long thrusting and slashing blades", []string{}, "https://wiki.cavesofqud.com/images/e/eb/Long_blade_proficiency_skill.png"}
	FeatSteadyHand             Feat            = Feat{"steady hand", "Pistol Proficiency", []FeatBuff{{TOHIT, "proficiency", []string{"using a pistol"}}}, []Ability{}, "This creature is skilled with pistols of various kinds", []string{}, "https://wiki.cavesofqud.com/images/b/b6/Steady_hand_skill.png"}
	FeatShortBladeExpertise    Feat            = Feat{"short blade expertise", "Short Blade Expertise", []FeatBuff{{TOHIT, "proficiency", []string{"using a short blade"}}}, []Ability{}, "This creature is skilled with small one-handed knives and blades", []string{}, "https://wiki.cavesofqud.com/images/6/67/Short_blade_expertise_skill.png"}
	FeatShortBlade             Feat            = Feat{"short blade", "Short Blade Proficiency", []FeatBuff{{TOHIT, "2", []string{"using a short blade"}}}, []Ability{}, "This creature is skilled with small one-handed knives and blades", []string{}, "https://wiki.cavesofqud.com/images/4/49/Short_blade_skill.png"}
	FeatFlurry                 Feat            = Feat{"flurry", "Flurry", []FeatBuff{}, []Ability{AbilityFlurry}, "This creature can make an attack with every hand at once", []string{}, "https://wiki.cavesofqud.com/images/2/20/Flurry_skill.png"}
	FeatMultiweaponFighting    Feat            = Feat{"multiweapon fighting", "Multiweapon Fighting", []FeatBuff{}, []Ability{}, "This creature can fight with multiple weapons at a time", []string{}, "https://wiki.cavesofqud.com/images/7/70/Multiweapon_fighting_skill.png"}
	FeatMultiweaponProficiency Feat            = Feat{"multiweapon proficiency", "Multiweapon Proficiency I", []FeatBuff{}, []Ability{}, "This creature has a chance to strike with its offhand weapons", []string{"multiweapon fighting"}, "https://wiki.cavesofqud.com/images/f/ff/Multiweapon_proficiency_skill.png"}
	FeatMultiweaponExpertise   Feat            = Feat{"multiweapon expertise", "Multiweapon Proficiency II", []FeatBuff{}, []Ability{}, "This creature has an improved chance to strike with its offhand weapons", []string{"multiweapon proficiency", "multiweapon fighting"}, "https://wiki.cavesofqud.com/images/d/d6/Multiweapon_expertise_skill.png"}
	FeatMultiweaponMastery     Feat            = Feat{"multiweapon mastery", "Multiweapon Proficiency III", []FeatBuff{}, []Ability{}, "This creature has a high chance to strike with its offhand weapons", []string{"multiweapon proficiency", "multiweapon expertise", "multiweapon fighting"}, "https://wiki.cavesofqud.com/images/2/25/Multiweapon_mastery_skill.png"}
	FeatTactics                Feat            = Feat{"tactics", "Tactics", []FeatBuff{}, []Ability{}, "This creature understands basic battle tactics", []string{}, "https://wiki.cavesofqud.com/images/d/de/Tactics_skill.png"}
	FeatCudgel                 Feat            = Feat{"cudgel", "Cudgel Use", []FeatBuff{}, []Ability{}, "This creature can use crushing and bludgeoning weapons", []string{}, "https://wiki.cavesofqud.com/images/f/f9/Cudgel_skill.png"}
	FeatBludgeon               Feat            = Feat{"bludgeon", "Bludgeon", []FeatBuff{}, []Ability{AbilityBludgeon}, "This creature can daze opponents while using a cudgel weapon", []string{"cudgel"}, "https://wiki.cavesofqud.com/images/d/d8/Bludgeon_skill.png"}
	FeatCharge                 Feat            = Feat{"charge", "Charge", []FeatBuff{{TOHIT, "1", []string{"during a charge"}}}, []Ability{AbilityCharge}, "This creature can charge forward to attack an enemy", []string{"tactics"}, "https://wiki.cavesofqud.com/images/3/34/Charge_skill.png"}
	FeatExtremeSpeed           Feat            = Feat{"extreme speed", "Extreme Speed", []FeatBuff{{TOHIT, "1", []string{"during a charge"}}}, []Ability{AbilityExtremeSpeed}, "This creature acts much quicker than others", []string{}, "https://wiki.cavesofqud.com/images/7/7e/Heightened_quickness_mutation.png"}
	Feats                      map[string]Feat = map[string]Feat{FeatSprint.Id: FeatSprint, FeatSwiftReflexes.Id: FeatSwiftReflexes, FeatSpry.Id: FeatSpry, FeatTumble.Id: FeatTumble, FeatJuke.Id: FeatJuke, FeatAxeProficiency.Id: FeatAxeProficiency, FeatSteadyHands.Id: FeatSteadyHands, FeatCudgelProficiency.Id: FeatCudgelProficiency, FeatLongBladeProficiency.Id: FeatLongBladeProficiency, FeatSteadyHand.Id: FeatSteadyHand, FeatShortBladeExpertise.Id: FeatShortBladeExpertise, FeatFlurry.Id: FeatFlurry, FeatShortBlade.Id: FeatShortBlade, FeatMultiweaponProficiency.Id: FeatMultiweaponProficiency, FeatMultiweaponExpertise.Id: FeatMultiweaponExpertise, FeatMultiweaponMastery.Id: FeatMultiweaponMastery, FeatMultiweaponFighting.Id: FeatMultiweaponFighting, FeatTactics.Id: FeatTactics, FeatCudgel.Id: FeatCudgel, FeatBludgeon.Id: FeatBludgeon, FeatCharge.Id: FeatCharge, FeatExtremeSpeed.Id: FeatExtremeSpeed}
)

type Duration string

var AllActions = []struct {
	Value  Duration
	TSName string
}{
	{Action, "ACTION"},
	{Reaction, "REACTION"},
	{ItemInteraction, "ITEM_INTERACTION"},
	{BonusAction, "BONUS_ACTION"},
	{FreeAction, "FREE_ACTION"},
	{Indefinite, "INDEFINITE"},
}

const (
	Action          Duration = "action"
	Reaction        Duration = "reaction"
	ItemInteraction Duration = "item_interaction"
	BonusAction     Duration = "bonus_action"
	FreeAction      Duration = "free_action"
	Indefinite      Duration = "indefinite"
)

type Stat string

const (
	StatNone  Stat = "none"
	AC        Stat = "AC"
	Speed     Stat = "Movement Speed"
	Level     Stat = "Level"
	PB        Stat = "Proficiency"
	HP        Stat = "Max HP"
	STR       Stat = "STR stat"
	DEX       Stat = "DEX stat"
	CON       Stat = "CON stat"
	INT       Stat = "INT stat"
	WIS       Stat = "WIS stat"
	CHA       Stat = "CHA stat"
	STR_BONUS Stat = "STR bonus"
	DEX_BONUS Stat = "DEX bonus"
	CON_BONUS Stat = "CON bonus"
	INT_BONUS Stat = "INT bonus"
	WIS_BONUS Stat = "WIS bonus"
	CHA_BONUS Stat = "CHA bonus"
	TOHIT     Stat = "to-hit"
	IN        Stat = "Initiative"
	MUT       Stat = "Mutation rank"
	STRSave   Stat = "STR save"
	DEXSave   Stat = "DEX save"
	CONSave   Stat = "CON save"
	INTSave   Stat = "INT save"
	WISSave   Stat = "WIS save"
	CHASave   Stat = "CHA save"
)

var AllStats = []struct {
	Value  Stat
	TSName string
}{
	{AC, "AC"},
	{Speed, "SPEED"},
	{Level, "LEVEL"},
	{PB, "PROFICIENCY"},
	{HP, "HP"},
	{STR, "STR"},
	{DEX, "DEX"},
	{CON, "CON"},
	{INT, "INT"},
	{WIS, "WIS"},
	{CHA, "CHA"},
	{IN, "INITIATIVE"},
	{MUT, "MUTATIONRANK"},
	{STRSave, "STRSAVE"},
	{DEXSave, "DEXSAVE"},
	{CONSave, "CONSAVE"},
	{INTSave, "INTSAVE"},
	{WISSave, "WISSAVE"},
	{CHASave, "CHASAVE"},
}

var AllDamageTypes = []struct {
	Value  DamageType
	TSName string
}{
	{Acid, "DMGACID"},
	{Blg, "DMGBLUDGEONING"},
	{Cold, "DMGCOLD"},
	{Fire, "DMGFIRE"},
	{Force, "DMGFORCE"},
	{Lightning, "DMGLIGHTNING"},
	{Necrotic, "DMGNECROTIC"},
	{Piercing, "DMGPIERCING"},
	{Poison, "DMGPOISON"},
	{Psychic, "DMGPSYCHIC"},
	{Radiant, "DMGRADIANT"},
	{Slashing, "DMGSLASHING"},
	{Thunder, "DMGTHUNDER"},
	{Umbral, "DMGUMBRAL"},
	{Cosmic, "DMGCOSMIC"},
}

type DamageType string

const (
	Acid      DamageType = "Acid Damage"
	Blg       DamageType = "Bludgeoning Damage"
	Cold      DamageType = "Cold Damage"
	Fire      DamageType = "Fire Damage"
	Force     DamageType = "Force Damage"
	Lightning DamageType = "Lightning Damage"
	Necrotic  DamageType = "Necrotic Damage"
	Piercing  DamageType = "Piercing Damage"
	Poison    DamageType = "Poison Damage"
	Psychic   DamageType = "Psychic Damage"
	Radiant   DamageType = "Radiant Damage"
	Slashing  DamageType = "Slashing Damage"
	Thunder   DamageType = "Thunder Damage"
	Umbral    DamageType = "Umbral Damage"
	Cosmic    DamageType = "Cosmic Damage"
)

var AllDmgAffinityLevels = []struct {
	Value  DmgAffinityLevel
	TSName string
}{
	{AffinityResistant, "RESISTANT"},
	{AffinityWeak, "WEAK"},
	{AffinityImmune, "IMMUNE"},
}

type DmgAffinityLevel string

const (
	AffinityResistant DmgAffinityLevel = "resistant"
	AffinityWeak      DmgAffinityLevel = "weak"
	AffinityImmune    DmgAffinityLevel = "immune"
)

type DmgAffinity struct {
	Level     DmgAffinityLevel `json:"level"`
	DmgType   DamageType       `json:"dmgType"`
	Reason    *string          `json:"reason"`
	Condition *string          `json:"condition"`
}

type DiceRoll struct {
	Dice      []string `json:"dice"`
	Offset    int      `json:"offset"`
	StatBonus Stat     `json:"statBonus"`
}

type StatOffset struct {
	Stat      Stat    `json:"stat"`
	Value     string  `json:"value"`
	Reason    *string `json:"reason"`
	Condition *string `json:"condition"`
}

type Statblock struct {
	Stats            map[Stat]string `json:"stats"`
	StatOffsets      []StatOffset    `json:"statOffsets"`
	DamageAffinities []DmgAffinity   `json:"dmgAffinities"`
	Items            []Weapon        `json:"items"`
	Feats            []Feat          `json:"feats"`
}

type WpnProperty string

const (
	WpnFinesse    WpnProperty = "Finesse;This weapon may use either DEX or STR for both attack and damage rolls."
	WpnAmmunition WpnProperty = "Ammunition;This weapon requires ammunition to fire."
	WpnHeavy      WpnProperty = "Heavy;Small creatures have disadvantage on attack rolls using this weapon."
	WpnLight      WpnProperty = "Light;This weapon is two-handed capable."
	WpnLoading    WpnProperty = "Loading;This weapon can only be fired once per attack action."
	WpnRange      WpnProperty = "Range;This is a ranged weapon."
	WpnSpecial    WpnProperty = "Special;This weapon has additional conditions for use."
	WpnThrown     WpnProperty = "Thrown;This weapon can be thrown."
	WpnTwoHanded  WpnProperty = "TwoHanded;This weapon requires two hands to attack with."
	WpnVersatile  WpnProperty = "Versatile;This weapon can be used with one or two hands."
	WpnImprovised WpnProperty = "Improvised;This makeshift weapon deals 1d4 damage, or the damage of a weapon it resembles."
	WpnSilvered   WpnProperty = "Silvered;This weapon is effective against certain creatures with nonmagical resistances."
	WpnNatural    WpnProperty = "Natural;This weapon is a part of this creature's body and can only be disarmed through dismemberment."
)

type WeaponRange struct {
	Normal int `json:"normal"`
	Long   int `json:"long"`
}

type EffectType string

const (
	DAZED   EffectType = "Dazed"
	STUNNED EffectType = "Stunned"
)

type Effect struct {
	Effect     EffectType `json:"effect"`
	Conditions []string   `json:"conditions"`
	Reasons    []string   `json:"reasons"`
	Rounds     DiceRoll   `json:"rounds"`
}

type WpnUsageCondition struct {
	Condition string   `json:"condition"`
	Reasons   []string `json:"reason"`
}

type Weapon struct {
	Name         string              `json:"name"`
	ImageUrl     string              `json:"imageUrl"`
	DmgType      DamageType          `json:"dmgType"`
	Dmg          DiceRoll            `json:"dmg"`
	DmgVersatile *DiceRoll           `json:"dmgVersatile"`
	Penetration  int                 `json:"penetration"`
	WpnRange     WeaponRange         `json:"wpnRange"`
	Conditions   []WpnUsageCondition `json:"conditions"`
	Effects      []Effect            `json:"effects"`
	StatOffsets  []StatOffset        `json:"statOffsets"`
	PageId       int                 `json:"pageid"`
}
