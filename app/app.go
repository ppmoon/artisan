package app

import (
	"github.com/ppmoon/artisan/app/config"
)

type App struct {
}

var a *App

func init() {
	a = NewApp()
}

// new app
func NewApp() *App {
	return &App{}
}

// run app
func Run() {
	a.Run()
}
func (a *App) Run() {
	// init config
	config.InitConfig("./config")
}
