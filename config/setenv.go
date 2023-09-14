package config

// set file name
// var EnVar = GetConfig("env")					// if runing locally

// func getConfigFile() string {
// 	// RUNNING ON THE SERVER
// 	path, _ := os.Executable()

// 	filePath := filepath.Dir(path)

// 	configFolder := fmt.Sprintf("%v/config/production.yaml", filePath)

// 	return configFolder
// }
// var configPath = getConfigFile()

var EnVar = GetConfig("production")