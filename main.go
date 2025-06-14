package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

type pageSearchResult struct {
	id      int
	key     string
	title   string
	excerpt string
}

type pageSearchResults struct {
	pages []pageSearchResult
}

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	categories := &Categories{}
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
			categories,
		},
		EnumBind: []any{
			AllScreens,
			AllStats,
			AllActions,
			AllDamageTypes,
			AllDmgAffinityLevels,
		},
		Mac: &mac.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
