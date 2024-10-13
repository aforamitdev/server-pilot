package main

import (
	"context"
	"embed"

	driver "github.com/aforamitdev/server-pilot/api/spilot/driver/grpc"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	grpc_driver := driver.NewGrpcDriver()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Server Pilot",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			grpc_driver.Startup(ctx)

		},
		Bind: []interface{}{
			app,
			grpc_driver,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
