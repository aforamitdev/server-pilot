package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	fmt.Println("stratup ")
	for {
		time.Sleep(time.Second * 1)
		runtime.EventsEmit(a.ctx, "LOG:RECEIVED", fmt.Sprintf("send event on %d", time.Now().String()))
	}
}

// // Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {

// }
