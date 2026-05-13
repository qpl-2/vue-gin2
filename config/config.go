package config

import (
	"github.com/spf13/viper"
	"log"
)

// 接收yml配置结构体
type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host         string
		Port         string
		User         string
		Password     string
		Name         string
		MaxTdleConns int
		MaxOpenConns int
	}
}

var Appconfig *Config

// viper读取文件
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}
	Appconfig = &Config{}

	err = viper.Unmarshal(Appconfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	initDB()
}
