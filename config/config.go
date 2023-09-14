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

	// LOCAL
	config.AddConfigPath("config/")


	// func getConfigFile() string {
	// 	// RUNNING ON THE SERVER
	// 	path, _ := os.Executable()
		
	// 	filePath := filepath.Dir(path)
		
	// 	configFolder := fmt.Sprintf("%v/config/production.yaml", filePath)
	
	// 	return configFolder
	// }
	// var configPath = getConfigFile()

	// err = config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("error: file not found %w", err))
	}

}

func GetConfig(env string) *viper.Viper {
	Init(env)

	return config
}