package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type Screen string

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

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	currScreen := struct {
		Value  Screen
		TSName string
	}{"Search", "SEARCH"}
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "quDnD",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			&currScreen,
		},
		EnumBind: []interface{}{
			AllScreens,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
