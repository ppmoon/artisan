package config

import (
	"github.com/spf13/viper"
	"io/ioutil"
)

const TOML = "toml"

func InitConfig(configPath string) {
	viper.AddConfigPath(configPath)
	viper.SetConfigType(TOML)
	files, err := ioutil.ReadDir(configPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		viper.SetConfigName(file.Name())
		err = viper.MergeInConfig()
		if err != nil {
			panic(file.Name() + " " + err.Error())
		}
	}
}
