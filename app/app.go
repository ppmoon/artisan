package app

import (
	"artisan/app/api"
	"artisan/app/config"
	"artisan/app/dao"
	"artisan/app/log"
	"fmt"
	"github.com/spf13/viper"
)

type App struct {
}

var a *App

const ConfigPath = "./config"
const Logo = "    _         _   _                 \n" +
	"   / \\   _ __| |_(_)___  __ _ _ __  \n" +
	"  / _ \\ | '__| __| / __|/ _` | '_ \\ \n" +
	" / ___ \\| |  | |_| \\__ \\ (_| | | | |\n" +
	"/_/   \\_\\_|   \\__|_|___/\\__,_|_| |_|"
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
	fmt.Println(Logo)
	// init config
	config.InitConfig(ConfigPath)
	// init log
	log.InitLog()
	log.Info("log init success")
	// init redis
	if viper.GetBool("app.redis_status") {
		dao.InitRedis()
		log.Info("redis init success")
	}
	if viper.GetBool("app.mysql_status") {
		// init mysql
		dao.InitMysql()
		log.Info("mysql init success")
	}
	if viper.GetBool("app.http_status") {
		// gin http must last start
		api.InitHttp()
	}
}
