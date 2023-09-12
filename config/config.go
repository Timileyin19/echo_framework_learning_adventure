package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	var err error

	config = viper.New()

	config.SetConfigType("yaml")

	config.SetConfigName(env)

	config.AddConfigPath("config/")

	err = config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("error: file not found %w", err))
	}

}

func GetConfig(env string) *viper.Viper {
	Init(env)

	return config
}