package app

import (
	"artisan/app/api"
	"artisan/app/config"
	"artisan/app/dao"
	"artisan/app/log"
)

type App struct {
}

var a *App

const ConfigPath = "./config"

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
	config.InitConfig(ConfigPath)
	// init log
	log.InitLog()
	log.Info("log init success")
	// init redis
	dao.InitRedis()
	log.Info("redis init success")
	// init mysql
	dao.InitMysql()
	log.Info("mysql init success")
	// gin http must last start
	api.InitHttp()
}
