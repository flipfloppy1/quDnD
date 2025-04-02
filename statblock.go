package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Stat string

const (
	AC      Stat = "ac"
	Speed   Stat = "speed"
	Level   Stat = "level"
	PB      Stat = "proficiency"
	HP      Stat = "hp"
	STR     Stat = "str"
	DEX     Stat = "dex"
	CON     Stat = "con"
	INT     Stat = "int"
	WIS     Stat = "wis"
	CHA     Stat = "cha"
	IN      Stat = "initiative"
	STRSave Stat = "strsave"
	DEXSave Stat = "dexsave"
	CONSave Stat = "consave"
	INTSave Stat = "intsave"
	WISSave Stat = "wissave"
	CHASave Stat = "chasave"
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
	{STRSave, "STRSAVE"},
	{DEXSave, "DEXSAVE"},
	{CONSave, "CONSAVE"},
	{INTSave, "INTSAVE"},
	{WISSave, "WISSAVE"},
	{CHASave, "CHASAVE"},
}

type DamageType string

const (
	Acid      DamageType = "dmgacid"
	Blg       DamageType = "dmgbludgeoning"
	Cold      DamageType = "dmgcold"
	Fire      DamageType = "dmgfire"
	Force     DamageType = "dmgforce"
	Lightning DamageType = "dmglightning"
	Necrotic  DamageType = "dmgnecrotic"
	Piercing  DamageType = "dmgpiercing"
	Poison    DamageType = "dmgpoison"
	Psychic   DamageType = "dmgpsychic"
	Radiant   DamageType = "dmgradiant"
	Slashing  DamageType = "dmgslashing"
	Thunder   DamageType = "dmgthunder"
)

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
	Dice   []string `json:"dice"`
	Offset int      `json:"offset"`
}

type StatOffset struct {
	Stat      Stat    `json:"stat"`
	Value     string  `json:"value"`
	Reason    *string `json:"reason"`
	Condition *string `json:"condition"`
}

type Statblock struct {
	Stats            map[Stat]string `json:"stats"`
	StatOffsets      StatOffset      `json:"statOffsets"`
	DamageAffinities DmgAffinity     `json:"dmgAffinities"`
	Weapons          []Weapon        `json:"weapons"`
}

type WpnProperty string

const (
	WpnFinesse    WpnProperty = "Finesse;This weapon may use either DEX or STR for both attack and damage rolls, mutually exclusive."
	WpnAmmunition WpnProperty = "Ammunition;This weapon requires ammunition to fire."
	WpnHeavy      WpnProperty = "Heavy;Small creatures have disadvantage on attack rolls using this weapon."
	WpnLight      WpnProperty = "Light;This weapon is two-handed capable."
	WpnLoading    WpnProperty = "Loading;This weapon can only be fired once per attack action."
	WpnRange      WpnProperty = "Range;This is a ranged weapon."
	WpnSpecial    WpnProperty = "Special;This weapon has additional conditions for use."
	WpnThrown     WpnProperty = "Thrown;This weapon can be thrown."
	WpnTwoHanded  WpnProperty = "TwoHanded;This weapon requires two hands to attack with."
	WpnVersatile  WpnProperty = "Versatile;This weapon can be used with one or two hands."
	WpnImprovised WpnProperty = "Improvised;This makeshift weapon deals 1d4 damage."
	WpnSilvered   WpnProperty = "Silvered;This weapon is effective against certain creatures with nonmagical resistances."
	WpnNatural    WpnProperty = "Natural;This weapon is a part of this creature's body and can only be disarmed through dismemberment."
)

type WeaponRange struct {
	Normal int `json:"normal"`
	Long   int `json:"long"`
}

type WpnEffect struct {
	Effect    string  `json:"effect"`
	Condition *string `json:"condition"`
	Reason    *string `json:"reason"`
}

type WpnUsageCondition struct {
	Condition string  `json:"condition"`
	Reason    *string `json:"reason"`
}

type Weapon struct {
	Name         string              `json:"name"`
	DmgType      DamageType          `json:"dmgType"`
	Dmg          DiceRoll            `json:"dmg"`
	DmgVersatile *DiceRoll           `json:"dmgVersatile"`
	Reach        int                 `json:"reach"`
	WpnRange     WeaponRange         `json:"wpnRange"`
	Conditions   []WpnUsageCondition `json:"conditions"`
	Effects      []WpnEffect         `json:"effects"`
	StatOffsets  []StatOffset        `json:"statOffsets"`
}

func ComposeStatblock(doc *goquery.Document) *Statblock {
	fmt.Println("Before error?")
	statblock := Statblock{}
	statblock.Stats = make(map[Stat]string)
	avSelect := doc.Find(".qud-stats-av")
	dvSelect := doc.Find(".qud-stats-dv")
	speedSelect := doc.Find(".qud-attribute-ms")
	healthSelect := doc.Find(".qud-stats-health")
	var av *string
	var dv *string
	var health *string
	var speed *string

	fmt.Println("Getting ac..")

	if avSelect == nil {
		return nil
	}

	if len(avSelect.Nodes) < 1 {
		return nil
	}

	if avSelect.Nodes[0] == nil {
		return nil
	} else {
		fmt.Println("ac")
		av = &avSelect.Find(".qud-stat-value").Nodes[0].FirstChild.Data
		dv = &dvSelect.Find(".qud-stat-value").Nodes[0].FirstChild.Data
		dvNum, _ := strconv.Atoi(*dv)
		dvNum = max(dvNum, int(math.Abs(float64(dvNum))))
		ac, _ := strconv.Atoi(*av)
		statblock.Stats[AC] = strconv.Itoa(int(math.Ceil(((float64(ac) + float64(dvNum)) / 3) + 10)))
		fmt.Println("ac end")
	}

	fmt.Println("Getting health..")

	if len(healthSelect.Nodes) < 1 {
		return nil
	}

	if healthSelect.Nodes[0] == nil {
		return nil
	} else {
		fmt.Println("health")
		health = &healthSelect.Find(".qud-stat-value").Nodes[0].FirstChild.Data
		statblock.Stats[HP] = *health
		fmt.Println("health end")
	}

	fmt.Println("Getting speed..")

	if len(speedSelect.Nodes) > 0 {
		if speedSelect.Nodes[0] != nil {
			fmt.Println("speed")
			speed = &speedSelect.Nodes[0].NextSibling.FirstChild.Data
			speedFloat, _ := strconv.Atoi(*speed)
			statblock.Stats[Speed] = strconv.Itoa(int(math.Round(((float64(speedFloat)/10)*3)/5) * 5)) // Convert to DnD scales and round to nearest 5ft
			fmt.Println("speed end")
		}
	}

	fmt.Println("Before or after error?")

	return &statblock
}
