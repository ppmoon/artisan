package dao

import (
	"artisan/app/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
)

var dbMysql *gorm.DB

func InitMysql() {
	//从config.toml当中读取数据库配置参数
	dbUser := viper.GetString("mysql.user")
	dbPassWord := viper.GetString("mysql.password")
	dbIp := viper.GetString("mysql.ip")
	dbPort := viper.GetString("mysql.port")
	dbName := viper.GetString("mysql.name")
	//设置数据库连接
	var err error
	dbMysql, err = gorm.Open("mysql", dbUser+`:`+dbPassWord+`@tcp(`+dbIp+`:`+dbPort+`)/`+dbName+`?charset=utf8mb4&parseTime=True&loc=Local`)
	//检查数据库连接
	if err != nil {
		log.Error("connect to mysql fails: ", err.Error())
		panic("failed to connect database")
	}
	//设置数据库连接池
	dbMaxIdle := viper.GetInt("mysql.max_idle")
	dbMaxOpen := viper.GetInt("mysql.max_open")
	dbMaxLifeTime := viper.GetDuration("mysql.max_life_time")
	//最大空闲数量
	dbMysql.DB().SetMaxIdleConns(dbMaxIdle)
	//最大开启数量
	dbMysql.DB().SetMaxOpenConns(dbMaxOpen)
	//超时设置
	dbMysql.DB().SetConnMaxLifetime(time.Second * dbMaxLifeTime)
	dbMysql.LogMode(viper.GetBool("mysql.log_mode"))
}
