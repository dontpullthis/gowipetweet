package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func MustInitialize() Config {
	config = Config{}
	yamlFile, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		log.Fatal("Failed to read config file "+ConfigPath+". ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal("Failed to parse config file "+ConfigPath+". ", err)
	}

	return config
}
