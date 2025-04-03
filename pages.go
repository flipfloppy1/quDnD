package main

import "github.com/PuerkitoBio/goquery"

type Screen string

const (
	Character Screen = "character"
	Concepts  Screen = "concepts"
	Creatures Screen = "creatures"
	Items     Screen = "items"
	Liquids   Screen = "liquids"
	Lore      Screen = "lore"
	Mechanics Screen = "mechanics"
	Custom    Screen = "custom"
)

var AllScreens = []struct {
	Value  Screen
	TSName string
}{
	{Character, "CHARACTER"},
	{Concepts, "CONCEPTS"},
	{Creatures, "CREATURES"},
	{Items, "ITEMS"},
	{Liquids, "LIQUIDS"},
	{Lore, "LORE"},
	{Mechanics, "MECHANICS"},
	{Custom, "CUSTOM"},
}

type PageInfo struct {
	PageType    Screen     `json:"pageType"`
	PageTitle   string     `json:"pageTitle"`
	ImgLink     *string    `json:"imgSrc"`
	Description *string    `json:"description"`
	Statblock   *Statblock `json:"statblock"`
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
		if len(imgSelect.Nodes) > 0 {
			if imgSelect.Nodes[0] != nil {
				if imgSelect.Nodes[0].FirstChild != nil {
					if imgSelect.Nodes[0].FirstChild.FirstChild != nil {
						if imgSelect.Nodes[0].FirstChild.FirstChild.FirstChild != nil {
							if len(imgSelect.Nodes[0].FirstChild.FirstChild.FirstChild.Attr) > 0 {
								for _, attr := range imgSelect.Nodes[0].FirstChild.FirstChild.FirstChild.Attr {
									if attr.Key == "src" {
										return &attr.Val
									}
								}
							}
						} else if len(imgSelect.Nodes[0].FirstChild.FirstChild.Attr) > 0 {
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
