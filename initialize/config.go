package initialize

import (
	"blog/server/global"
	"os"

	"gopkg.in/yaml.v3"
)

// Parse the config file
func Config(test bool) {
	var yamlFile []byte
	var err error
	if test {
		yamlFile, err = os.ReadFile("../config.yaml")
		if err != nil { 
			panic(err)
		}
	} else {
		yamlFile, err = os.ReadFile("./config.yaml")
		if err != nil { 
			panic(err)
		}
	}
	err = yaml.Unmarshal(yamlFile, &global.GLOBAL_CONFIG)
	if err != nil {
		panic(err)
	}
}