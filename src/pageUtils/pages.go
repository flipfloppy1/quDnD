package pageUtils

import (
	"github.com/PuerkitoBio/goquery"
)

type Screen string

const (
	Search    Screen = "search"
	Character Screen = "character"
	Concepts  Screen = "concepts"
	Creatures Screen = "creatures"
	Items     Screen = "items"
	Liquids   Screen = "liquids"
	Lore      Screen = "lore"
	Mechanics Screen = "mechanics"
	Mutations Screen = "mutations"
	Other     Screen = "other"
	Custom    Screen = "custom"
)

var AllScreens = []struct {
	Value  Screen
	TSName string
}{
	{Search, "SEARCH"},
	{Character, "CHARACTER"},
	{Concepts, "CONCEPTS"},
	{Creatures, "CREATURES"},
	{Items, "ITEMS"},
	{Liquids, "LIQUIDS"},
	{Lore, "LORE"},
	{Mechanics, "MECHANICS"},
	{Mutations, "MUTATIONS"},
	{Other, "OTHER"},
	{Custom, "CUSTOM"},
}

func GetDescription(doc *goquery.Document) *string {
	selection := doc.Find(".qud-look-modern-text").Find(".poem").Find("p").Find("span")
	if len(selection.Nodes) > 0 {
		if selection.Nodes[0] != nil {
			if selection.Nodes[0].FirstChild != nil {
				return &selection.Nodes[0].FirstChild.Data
			}
		}
	}
	desc := doc.Find(".mw-parser-output > p").First().Text()
	return &desc

}

func GetPageImg(doc *goquery.Document) *string {
	imgSelect := doc.Find(".infobox-imagearea")
	if imgSelect != nil {
		imgTagSelect := imgSelect.Find(".mw-file-element")
		if imgTagSelect == nil {
			return nil
		}
		val, exists := imgTagSelect.Attr("src")
		if exists {
			return &val
		} else {
			return nil
		}
	}

	imgSelect = doc.Find(".qud-infobox-carousel-static-image")
	if imgSelect != nil {
		if len(imgSelect.Nodes) > 0 {
			if imgSelect.Nodes[0] != nil {
				if imgSelect.Nodes[0].FirstChild != nil {
					if imgSelect.Nodes[0].FirstChild.FirstChild != nil {
						if len(imgSelect.Nodes[0].FirstChild.FirstChild.Attr) > 0 {
							for _, attr := range imgSelect.Nodes[0].FirstChild.FirstChild.Attr {
								if attr.Key == "src" {
									return &attr.Val
								}
							}
						}
					}
				}
			}
		}
	}

	return nil
}
