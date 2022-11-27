package initialize

import (
	"blog/server/global"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Parse the config file
func InitializeConfig() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &global.GLOBAL_CONFIG)
	if err != nil {
		panic(err)
	}
}