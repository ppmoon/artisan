package dao

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
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
	dsn := dbUser + `:` + dbPassWord + `@tcp(` + dbIp + `:` + dbPort + `)/` + dbName + `?charset=utf8mb4&parseTime=True&loc=Local`
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	dbMysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	//检查数据库连接
	if err != nil {
		panic("failed to connect database " + err.Error())
	}
	sqlDb, err := dbMysql.DB()
	if err != nil {
		panic(err)
	}
	//设置数据库连接池
	dbMaxIdle := viper.GetInt("mysql.max_idle")
	dbMaxOpen := viper.GetInt("mysql.max_open")
	dbMaxLifeTime := viper.GetDuration("mysql.max_life_time")
	//最大空闲数量
	sqlDb.SetMaxIdleConns(dbMaxIdle)
	//最大开启数量
	sqlDb.SetMaxOpenConns(dbMaxOpen)
	//超时设置
	sqlDb.SetConnMaxLifetime(time.Second * dbMaxLifeTime)
}
