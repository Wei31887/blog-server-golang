package initialize

import (
	"blog/server/initialize/global"
	"fmt"

	"github.com/spf13/viper"
)

// Parse the config file
func Config(path string) (err error){
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)               // optionally look for config in the working directory
	viper.AutomaticEnv()

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&global.GLOBAL_CONFIG)
	if err != nil { // Handle errors unmarshaling the config
	    panic(fmt.Errorf("fatal error config file: %w", err))
	}
    return
}