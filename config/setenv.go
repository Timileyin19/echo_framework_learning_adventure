package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// set file name
// var EnVar = GetConfig("env")					// if runing locally


func getConfigFile() string {
	// RUNNING ON THE SERVER
	path, _ := os.Executable()
	
	filePath := filepath.Dir(path)
	
	configFile := fmt.Sprintf("%v/config/production.yaml", filePath)

	return configFile
}
var configPath = getConfigFile()

var EnVar = GetConfig(configPath)