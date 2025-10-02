package main

import (
	"context"
	"fmt"

	"github.com/BrandenWilliams/VHS/ffconvert"
	"github.com/BrandenWilliams/VHS/ffconvert/linuxcliargs"
)

// App struct
type App struct {
	ctx context.Context

	ffc     ffconvert.FFConvert
	preSets []linuxcliargs.PreSet
}

func NewApp() *App {
	return &App{}
}

func (a *App) getPreSets() {
	a.preSets = a.ffc.LCliA.GetPreSets()
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	a.getPreSets()
}

// DELETE LONG TERM
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
