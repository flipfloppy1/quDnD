package statblock

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"quDnD/src/pageUtils"

	"github.com/PuerkitoBio/goquery"
)

type DbPage struct {
	Page   PageInfo `json:"pageInfo"`
	Exists bool     `json:"exists"`
}

type PageInfo struct {
	PageType    pageUtils.Screen `json:"pageType"`
	PageTitle   string           `json:"pageTitle"`
	ImgLink     *string          `json:"imgSrc"`
	Description *string          `json:"description"`
	Statblock   *Statblock       `json:"statblock"`
	PageId      int              `json:"pageid"`
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
		pageidResp := pageUtils.QudAction("action=query&titles=" + friendlyId)
		var pageResp pageUtils.RestPagesResultJson
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
				switch index {
				case 0:
					penetration, err := strconv.Atoi(node.FirstChild.Data)
					if err != nil {
						return fmt.Errorf("Unable to parse penetration: %s", err.Error())
					}
					weapon.Penetration = penetration
				case 1:
					weapon.Dmg.Dice = append(weapon.Dmg.Dice, node.FirstChild.Data)
				}
			}
		} else {
			return fmt.Errorf("Unable to read item stats for %s", itemUrl)
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

		statblock.Stats[HP] = strconv.Itoa(int(math.Ceil(float64(qCon)/2.0+float64(qLvl)*float64(qLvl))) / 2)
	}

	if len(invSelect.Nodes) > 0 {
		if invLinkSelect := invSelect.Find(".qud-image-link-image-container"); invLinkSelect != nil {
			if invImgSelect := invLinkSelect.Find(".mw-file-element"); invImgSelect != nil {
				currItems := make([]string, 0)

				for _, sibling := range invImgSelect.EachIter() {
					equipmentItem := Weapon{}
					equipmentItem.ImageUrl, _ = sibling.Attr("src")
					equipmentItem.ImageUrl = "https://wiki.cavesofqud.com" + equipmentItem.ImageUrl
					if slices.Contains(currItems, equipmentItem.ImageUrl) {
						continue
					} else {
						currItems = append(currItems, equipmentItem.ImageUrl)
					}
					itemLink, exists := sibling.Parent().Attr("href")
					if !exists {
						fmt.Printf("Link for item %s did not exist\n", equipmentItem.ImageUrl)
						continue
					}
					err := setItemStats(&equipmentItem, itemLink)
					if err != nil {
						fmt.Println(err.Error())
						continue
					} else {
						statblock.Items = append(statblock.Items, equipmentItem)
					}
					if invNameSelect := sibling.ParentsUntil("qud-inv-favilink-wrapper").Find(".qud-image-link"); invNameSelect != nil {
						invNameSelect = invNameSelect.First().Children().First().Children().First()

						statblock.Items[len(statblock.Items)-1].Name = invNameSelect.AttrOr("title", "item")
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
					num += min(int(lvl/(1.0/3.0)), 2)
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
		if quickness >= 150 {
			statblock.Feats = append(statblock.Feats, FeatExtremeSpeed)
		}

		quicknessMul := math.Pow((float64(quickness) / 100.0), 1.5)
		speed, _ := strconv.Atoi(statblock.Stats[Speed])
		speed = int(float64(speed) * quicknessMul)

		// Round down to the nearest 5ft
		speed = speed - speed%5

		statblock.Stats[Speed] = strconv.FormatInt(int64(speed), 10)
	}

	return &statblock
}
