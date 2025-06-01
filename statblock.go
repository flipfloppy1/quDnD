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
	StatOffsets      []StatOffset    `json:"statOffsets"`
	DamageAffinities []DmgAffinity   `json:"dmgAffinities"`
	Items            []Weapon        `json:"items"`
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
	ImageUrl     string              `json:"imageUrl"`
	DmgType      DamageType          `json:"dmgType"`
	Dmg          DiceRoll            `json:"dmg"`
	DmgVersatile *DiceRoll           `json:"dmgVersatile"`
	Penetration  int                 `json:"penetration"`
	WpnRange     WeaponRange         `json:"wpnRange"`
	Conditions   []WpnUsageCondition `json:"conditions"`
	Effects      []WpnEffect         `json:"effects"`
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
		fmt.Println(pageidResp)
		var pageResp RestPagesResultJson
		err = json.Unmarshal([]byte(pageidResp), &pageResp)
		if err != nil {
			return fmt.Errorf("Unable to get item pageid: %s", err.Error())
		}
		fmt.Println(pageResp.Query.PageMap)
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
	fmt.Println("Before error?")
	statblock := Statblock{}
	statblock.Stats = make(map[Stat]string)
	avSelect := doc.Find(".qud-stats-av")
	dvSelect := doc.Find(".qud-stats-dv")
	speedSelect := doc.Find(".qud-attribute-ms")
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

	return &statblock
}
