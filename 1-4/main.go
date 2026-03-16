package main

import (
	"1-4/backend/services/app"
	"1-4/backend/services/cyphers"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	var a cyphers.Atbash
	var s cyphers.Scytale
	var p cyphers.Polybius
	var c cyphers.Caesar

	// Create an instance of the window structure
	window := app.NewApp(&a, &s, &p, &c)
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "1-4",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        window.Startup,
		Bind: []interface{}{
			window,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
