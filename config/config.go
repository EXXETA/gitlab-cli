package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	// default path of the config file
	defaultConfigFilePath = "$HOME"

	// default config file name
	defaultConfigFileName = ".gitlab-cli-config"
)

// InitConfig initializes the config file
func InitConfig(configFileFromFlag string) {

	// if the config file was set in a flag
	if configFileFromFlag != "" {
		viper.SetConfigFile(configFileFromFlag)
	}

	// set the default config file info
	viper.SetConfigName(defaultConfigFileName)
	viper.AddConfigPath(defaultConfigFilePath)

	// read the config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error wile reading the config file.", err)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// GetConfigValue returns the value of a config parameter
func GetConfigValue(parameterName string) string {
	return viper.GetString(parameterName)
}
