package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Attack struct {
	DamageType DamageType `json:"dmgType"`
	Damage     DiceRoll   `json:"damage"`
	Conditions []string   `json:"conditions"`
}

type Ability struct {
	Duration    Duration `json:"duration"`
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	Conditions  []string `json:"conditions"`
	Attacks     []Attack `json:"attacks"`
	Effects     []Effect `json:"effects"`
}

type Feat struct {
	Id            string     `json:"id"`
	Name          string     `json:"name"`
	Buffs         []FeatBuff `json:"buffs"`
	Abilities     []Ability  `json:"abilities"`
	Description   string     `json:"description"`
	Prerequisites []string   `json:"prereqs"`
}

type FeatBuff struct {
	Stat       Stat     `json:"stat"`
	Value      string   `json:"value"`
	Conditions []string `json:"conditions"`
}

var (
	AbilityJuke     Ability = Ability{Action, "Whirl past an opponent, swapping places with it", "You use an action to swap places with a creature within 5ft of you that is your size or smaller. You and your allies will not provoke opportunity attacks from the target until your next turn.", []string{"1 action", "target is within 5 feet", "target is creature's size or smaller"}, []Attack{}, []Effect{}}
	AbilityFlurry   Ability = Ability{Action, "Make an attack action with every hand at once, including hands granted by mutation or technology", "Once per encounter, you may expend an action to make an attack using every hand you have. For the purposes of other abilities, these attacks count as discrete attack actions.", []string{"1 action", "target is in melee range", "once per encounter"}, []Attack{}, []Effect{}}
	AbilityCharge   Ability = Ability{Action, "Perform a melee attack after charging between 10-20ft forward", "Once per encounter, you can charge between 10-20ft towards an enemy of your choosing, making an attack with your primary weapon with +1 to-hit.", []string{"1 action", "target is between 10 and 20 feet", "once per encounter"}, []Attack{}, []Effect{}}
	AbilityBludgeon Ability = Ability{Action, "Make an attack with a cudgel, dazing an opponent", "When you attack with a cudgel, roll a d4. On a 4, your attack inflicts Dazed on your opponent. If your opponent is already Dazed you instead Stun them for 1 round.", []string{"1 action", "target is in melee range"}, []Attack{}, []Effect{{DAZED, []string{"opponent is not already dazed", "4 on 1d4 to Daze"}, []string{}, DiceRoll{[]string{"1d4"}, 0, StatNone}}, {STUNNED, []string{"opponent is already dazed", "4 on 1d4 to Daze"}, []string{}, DiceRoll{[]string{}, 1, StatNone}}}}
)

var (
	FeatSprint                 Feat            = Feat{"sprint", "Sprint", []FeatBuff{{Speed, "10", []string{}}}, []Ability{}, "This creature is capable of moving quickly", []string{}}
	FeatSwiftReflexes          Feat            = Feat{"swift reflexes", "Swift Reflexes", []FeatBuff{{AC, "2", []string{"As a reaction to a projectile attack that targets the creature"}}}, []Ability{}, "This creature has a +2 AC bonus when flinching away from projectile attacks", []string{}}
	FeatSpry                   Feat            = Feat{"spry", "Spry", []FeatBuff{{DEX, "3", []string{}}}, []Ability{}, "This creature is nimble and gains a +3 bonus to its DEX", []string{}}
	FeatTumble                 Feat            = Feat{"tumble", "Tumble", []FeatBuff{{DEX, "1", []string{}}}, []Ability{}, "This creature gains a +1 bonus to its DEX and can Juke as a bonus action if it has the feat", []string{}}
	FeatJuke                   Feat            = Feat{"juke", "Juke", []FeatBuff{}, []Ability{AbilityJuke}, "This creature can whirl past an opponent, swapping places with it", []string{}}
	FeatAxeProficiency         Feat            = Feat{"axe proficiency", "Axe Proficiency", []FeatBuff{{TOHIT, "proficiency", []string{"using an axe"}}}, []Ability{}, "This creature is proficient with axes", []string{}}
	FeatSteadyHands            Feat            = Feat{"steady hands", "Bow and Rifle Proficiency", []FeatBuff{{TOHIT, "proficiency", []string{"using a bow or rifle"}}}, []Ability{}, "This creature is skilled with bows and rifles", []string{}}
	FeatCudgelProficiency      Feat            = Feat{"cudgel proficiency", "Cudgel Proficiency", []FeatBuff{{TOHIT, "proficiency", []string{"using a cudgel"}}}, []Ability{}, "This creature is skilled with crushing and bludgeoning weapons", []string{}}
	FeatLongBladeProficiency   Feat            = Feat{"long blade proficiency", "Long Blade Proficiency", []FeatBuff{{TOHIT, "proficiency", []string{"using a long blade"}}}, []Ability{}, "This creature is skilled with long thrusting and slashing blades", []string{}}
	FeatSteadyHand             Feat            = Feat{"steady hand", "Pistol Proficiency", []FeatBuff{{TOHIT, "proficiency", []string{"using a pistol"}}}, []Ability{}, "This creature is skilled with pistols of various kinds", []string{}}
	FeatShortBladeExpertise    Feat            = Feat{"short blade expertise", "Short Blade Expertise", []FeatBuff{{TOHIT, "proficiency", []string{"using a short blade"}}}, []Ability{}, "This creature is skilled with small one-handed knives and blades", []string{}}
	FeatShortBlade             Feat            = Feat{"short blade", "Short Blade Proficiency", []FeatBuff{{TOHIT, "2", []string{"using a short blade"}}}, []Ability{}, "This creature is skilled with small one-handed knives and blades", []string{}}
	FeatFlurry                 Feat            = Feat{"flurry", "Flurry", []FeatBuff{}, []Ability{AbilityFlurry}, "This creature can make an attack with every hand at once", []string{}}
	FeatMultiweaponFighting    Feat            = Feat{"multiweapon fighting", "Multiweapon Fighting", []FeatBuff{}, []Ability{}, "This creature can fight with multiple weapons at a time", []string{}}
	FeatMultiweaponProficiency Feat            = Feat{"multiweapon proficiency", "Multiweapon Proficiency I", []FeatBuff{}, []Ability{}, "This creature has a chance to strike with its offhand weapons", []string{"multiweapon fighting"}}
	FeatMultiweaponExpertise   Feat            = Feat{"multiweapon expertise", "Multiweapon Proficiency II", []FeatBuff{}, []Ability{}, "This creature has an improved chance to strike with its offhand weapons", []string{"multiweapon proficiency", "multiweapon fighting"}}
	FeatMultiweaponMastery     Feat            = Feat{"multiweapon mastery", "Multiweapon Proficiency III", []FeatBuff{}, []Ability{}, "This creature has a high chance to strike with its offhand weapons", []string{"multiweapon proficiency", "multiweapon expertise", "multiweapon fighting"}}
	FeatTactics                Feat            = Feat{"tactics", "Tactics", []FeatBuff{}, []Ability{}, "This creature understands basic battle tactics", []string{}}
	FeatCudgel                 Feat            = Feat{"cudgel", "Cudgel Use", []FeatBuff{}, []Ability{}, "This creature can use crushing and bludgeoning weapons", []string{}}
	FeatBludgeon               Feat            = Feat{"bludgeon", "Bludgeon", []FeatBuff{}, []Ability{}, "This creature can daze opponents while using a cudgel weapon", []string{"cudgel"}}
	FeatCharge                 Feat            = Feat{"charge", "Charge", []FeatBuff{{TOHIT, "1", []string{"during a charge"}}}, []Ability{AbilityCharge}, "This creature can charge forward to attack an enemy", []string{"tactics"}}
	Feats                      map[string]Feat = map[string]Feat{FeatSprint.Id: FeatSprint, FeatSwiftReflexes.Id: FeatSwiftReflexes, FeatSpry.Id: FeatSpry, FeatTumble.Id: FeatTumble, FeatJuke.Id: FeatJuke, FeatAxeProficiency.Id: FeatAxeProficiency, FeatSteadyHands.Id: FeatSteadyHands, FeatCudgelProficiency.Id: FeatCudgelProficiency, FeatLongBladeProficiency.Id: FeatLongBladeProficiency, FeatSteadyHand.Id: FeatSteadyHand, FeatShortBladeExpertise.Id: FeatShortBladeExpertise, FeatFlurry.Id: FeatFlurry, FeatShortBlade.Id: FeatShortBlade, FeatMultiweaponProficiency.Id: FeatMultiweaponProficiency, FeatMultiweaponExpertise.Id: FeatMultiweaponExpertise, FeatMultiweaponMastery.Id: FeatMultiweaponMastery, FeatMultiweaponFighting.Id: FeatMultiweaponFighting, FeatTactics.Id: FeatTactics, FeatCudgel.Id: FeatCudgel, FeatBludgeon.Id: FeatBludgeon, FeatCharge.Id: FeatCharge}
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
}

const (
	Action          Duration = "action"
	Reaction        Duration = "reaction"
	ItemInteraction Duration = "item_interaction"
	BonusAction     Duration = "bonus_action"
	FreeAction      Duration = "free_action"
)

type Stat string

const (
	StatNone  Stat = "none"
	AC        Stat = "ac"
	Speed     Stat = "speed"
	Level     Stat = "level"
	PB        Stat = "proficiency"
	HP        Stat = "hp"
	STR       Stat = "str"
	DEX       Stat = "dex"
	CON       Stat = "con"
	INT       Stat = "int"
	WIS       Stat = "wis"
	CHA       Stat = "cha"
	STR_BONUS Stat = "strBonus"
	DEX_BONUS Stat = "dexBonus"
	CON_BONUS Stat = "conBonus"
	INT_BONUS Stat = "intBonus"
	WIS_BONUS Stat = "wisBonus"
	CHA_BONUS Stat = "chaBonus"
	TOHIT     Stat = "tohit"
	IN        Stat = "initiative"
	STRSave   Stat = "strsave"
	DEXSave   Stat = "dexsave"
	CONSave   Stat = "consave"
	INTSave   Stat = "intsave"
	WISSave   Stat = "wissave"
	CHASave   Stat = "chasave"
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
	Rounds     DiceRoll
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

func setItemStats(weapon *Weapon, itemUrl string) error {
	errorFormat := "Item stats error: %s"
	resp, err := http.Get("https://wiki.cavesofqud.com" + itemUrl)

	if err != nil {
		resp.Body.Close()
		return fmt.Errorf(errorFormat, err.Error())
	}

	if resp.StatusCode < 300 && resp.StatusCode >= 200 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)

		friendlyId, _ := strings.CutPrefix(itemUrl, "/wiki/")
		pageidResp := qudAction("action=query&titles=" + friendlyId)
		var pageResp RestPagesResultJson
		err = json.Unmarshal([]byte(pageidResp), &pageResp)
		if err != nil {
			return fmt.Errorf("Unable to get item pageid: %s", err.Error())
		}
		for _, pageRes := range pageResp.Query.PageMap {
			weapon.PageId = pageRes.Pageid
		}

		statSelect := doc.Find(".qud-item-stat-value")

		if len(statSelect.Nodes) > 0 {
			for index, node := range statSelect.Nodes {
				if index == 0 {
					penetration, err := strconv.Atoi(node.FirstChild.Data)
					if err != nil {
						return fmt.Errorf("Unable to parse penetration: %s", err.Error())
					}
					weapon.Penetration = penetration
				} else if index == 1 {
					weapon.Dmg.Dice = append(weapon.Dmg.Dice, node.FirstChild.Data)
				}
			}
		} else {
			return fmt.Errorf("Unable to read item stats for %s", itemUrl)
		}

		titleSelect := doc.Find(".mw-first-heading")

		if len(titleSelect.Nodes) > 0 {
			weapon.Name, _ = strings.CutPrefix(titleSelect.Nodes[0].FirstChild.Data, "Data:")
		} else {
			return fmt.Errorf("Unable to find item name for %s", itemUrl)
		}

	} else {
		return fmt.Errorf(errorFormat, resp.Status)
	}

	return nil
}

func ComposeStatblock(doc *goquery.Document) *Statblock {
	statblock := Statblock{}
	statblock.Stats = make(map[Stat]string)
	avSelect := doc.Find(".qud-stats-av")
	dvSelect := doc.Find(".qud-stats-dv")
	skillSelect := doc.Find("#collapsible-qud-qud-skills")
	speedSelect := doc.Find(".qud-attribute-ms")
	quicknessSelect := doc.Find(".qud-attribute-qn")
	healthSelect := doc.Find(".qud-stats-health")
	attrSelect := doc.Find(".qud-attributes-wrapper")
	invSelect := doc.Find(".qud-inventory-item")
	var av *string
	var dv *string
	var health *string
	var speed *string

	if avSelect == nil {
		return nil
	}

	if len(avSelect.Nodes) < 1 {
		return nil
	}

	if avSelect.Nodes[0] == nil {
		return nil
	} else {
		av = &avSelect.Find(".qud-stat-value").Nodes[0].FirstChild.Data
		dv = &dvSelect.Find(".qud-stat-value").Nodes[0].FirstChild.Data
		dvNum, _ := strconv.Atoi(*dv)
		dvNum = max(dvNum, int(math.Abs(float64(dvNum))))
		ac, _ := strconv.Atoi(*av)
		baseAC := (float64(ac)+float64(dvNum))/3 + 10
		skewMul := 1.15
		distMiddle := 15.0
		scale := 2500.0
		finalAC := (skewMul - (math.Pow((distMiddle-baseAC), 2) / scale) - math.Pow(baseAC, 2)/(2*scale)) * baseAC
		statblock.Stats[AC] = strconv.Itoa(int(math.Ceil(finalAC)))
	}

	if len(healthSelect.Nodes) < 1 {
		return nil
	}

	if healthSelect.Nodes[0] == nil {
		return nil
	} else {
		health = &healthSelect.Find(".qud-stat-value").Nodes[0].FirstChild.Data
		statblock.Stats[HP] = *health
	}

	if len(speedSelect.Nodes) > 0 {
		if speedSelect.Nodes[0] != nil {
			speed = &speedSelect.Nodes[0].NextSibling.FirstChild.Data
			speedFloat, _ := strconv.Atoi(*speed)
			statblock.Stats[Speed] = strconv.Itoa(int(math.Round(((float64(speedFloat)/10)*3)/5) * 5)) // Convert to DnD scales and round to nearest 5ft
		}
	}

	if len(attrSelect.Nodes) > 0 {
		qStr, _ := strconv.Atoi(attrSelect.Find(".qud-attribute-st").Nodes[0].NextSibling.FirstChild.Data)
		qDex, _ := strconv.Atoi(attrSelect.Find(".qud-attribute-ag").Nodes[0].NextSibling.FirstChild.Data)
		qCon, _ := strconv.Atoi(attrSelect.Find(".qud-attribute-to").Nodes[0].NextSibling.FirstChild.Data)
		qInt, _ := strconv.Atoi(attrSelect.Find(".qud-attribute-in").Nodes[0].NextSibling.FirstChild.Data)
		qWis, _ := strconv.Atoi(attrSelect.Find(".qud-attribute-wi").Nodes[0].NextSibling.FirstChild.Data)
		qCha, _ := strconv.Atoi(attrSelect.Find(".qud-attribute-eg").Nodes[0].NextSibling.FirstChild.Data)
		qLvl, _ := strconv.Atoi(strings.TrimSpace(doc.Find(".qud-character-level-value").Nodes[0].FirstChild.Data))

		statblock.Stats[STR] = strconv.Itoa(qStr * 2 / 3)
		statblock.Stats[DEX] = strconv.Itoa(qDex * 2 / 3)
		statblock.Stats[CON] = strconv.Itoa(qCon * 2 / 3)
		statblock.Stats[INT] = strconv.Itoa(qInt * 2 / 3)
		statblock.Stats[WIS] = strconv.Itoa(qWis * 2 / 3)
		statblock.Stats[CHA] = strconv.Itoa(qCha * 2 / 3)

		lvl := float64(qLvl) / 5 * 3
		if lvl == 0 {
			statblock.Stats[Level] = "0"
		} else if lvl <= 0.25 {
			statblock.Stats[Level] = "1/8"
		} else if lvl <= 0.5 {
			statblock.Stats[Level] = "1/4"
		} else if lvl <= 1 {
			statblock.Stats[Level] = "1/2"
		} else {
			statblock.Stats[Level] = strconv.Itoa(int(lvl))
		}

		statblock.Stats[HP] = strconv.Itoa(int(math.Ceil(float64(qCon) / 2.0 * float64(qLvl))))
	}

	if len(invSelect.Nodes) > 0 {
		invLinkSelect := invSelect.Find(".qud-image-link-image-container")
		currItems := make([]string, 0)

		for _, node := range invLinkSelect.Nodes {
			equipmentItem := Weapon{}
			for _, attr := range node.FirstChild.FirstChild.Attr {
				if attr.Key == "src" {
					equipmentItem.ImageUrl = "https://wiki.cavesofqud.com" + attr.Val
				}
			}
			if slices.Contains(currItems, equipmentItem.ImageUrl) {
				break
			} else {
				currItems = append(currItems, equipmentItem.ImageUrl)
			}
			for _, attr := range node.FirstChild.Attr {
				if attr.Key == "href" {
					err := setItemStats(&equipmentItem, attr.Val)
					if err != nil {
						fmt.Println(err.Error())
						continue
					} else {
						statblock.Items = append(statblock.Items, equipmentItem)
					}
				}
			}
		}
	}

	if len(skillSelect.Nodes) > 0 {
		skills := skillSelect.Find(".qud-skill-entry")
		for _, node := range skills.Nodes {
			if node.FirstChild != nil && node.FirstChild.FirstChild != nil {
				statblock.Feats = append(statblock.Feats, Feats[strings.ToLower(node.FirstChild.FirstChild.Data)])
			}
		}
	}

	for _, feat := range statblock.Feats {
		for _, buff := range feat.Buffs {
			if !(len(buff.Conditions) > 0) {
				stat := statblock.Stats[buff.Stat]
				if buff.Value == "proficiency" {
					num, err := strconv.Atoi(stat)
					if err != nil {
						continue
					}
					lvl := 0.0
					if strings.Contains(statblock.Stats[Level], "/") {
						numer, _ := strconv.Atoi(strings.Split(statblock.Stats[Level], "/")[0])
						denom, _ := strconv.Atoi(strings.Split(statblock.Stats[Level], "/")[1])
						lvl = float64(numer) / float64(denom)
					} else {
						lvlInt, _ := strconv.Atoi(statblock.Stats[Level])
						lvl = float64(lvlInt)
					}
					num += min(int(lvl/(1/3)), 2)
					statblock.Stats[buff.Stat] = strconv.FormatInt(int64(num), 10)
				} else {
					stat, err := strconv.Atoi(statblock.Stats[buff.Stat])
					if err != nil {
						continue
					}
					buffVal, err := strconv.Atoi(buff.Value)
					stat += buffVal
					statblock.Stats[buff.Stat] = strconv.FormatInt(int64(stat), 10)
				}
			}
		}
	}

	if len(quicknessSelect.Nodes) > 0 {
		quickStr := quicknessSelect.Nodes[0].NextSibling.FirstChild.Data
		quickness, _ := strconv.Atoi(quickStr)

		quicknessMul := math.Pow((float64(quickness) / 100.0), 1.5)
		speed, _ := strconv.Atoi(statblock.Stats[Speed])
		speed = int(float64(speed) * quicknessMul)

		// Round down to the nearest 5ft
		speed = speed - speed%5

		statblock.Stats[Speed] = strconv.FormatInt(int64(speed), 10)
	}

	return &statblock
}
