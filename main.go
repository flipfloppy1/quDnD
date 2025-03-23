package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type Screen string

type pageSearchResult struct {
	id      int
	key     string
	title   string
	excerpt string
}

type pageSearchResults struct {
	pages []pageSearchResult
}

const (
	Search    Screen = "Search"
	Creatures Screen = "Creatures"
	Objects   Screen = "Objects"
	Liquids   Screen = "Liquids"
	Lore      Screen = "Lore"
	Mechanics Screen = "Mechanics"
)

var AllScreens = []struct {
	Value  Screen
	TSName string
}{
	{Search, "SEARCH"},
	{Creatures, "CREATURES"},
	{Objects, "OBJECTS"},
	{Liquids, "LIQUIDS"},
	{Lore, "LORE"},
	{Mechanics, "MECHANICS"},
}

type PageInfo struct {
	pageType Screen
}

type CategoryMembers struct {
	categoryMap map[string][]int
}

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "quDnD",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 15, G: 59, B: 58, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
		},
		EnumBind: []any{
			AllScreens,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
