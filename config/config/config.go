package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config = viper.New()

// 初始化配置
func init() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")

	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal("Config not found: ", err)
	}
}
