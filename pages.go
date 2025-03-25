package main

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
	Description *string    `json:"description"`
	Statblock   *Statblock `json:"statblock"`
}
