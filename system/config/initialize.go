package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func MustInitialize(cfgPath string) Config {
	cfg := Config{}
	yamlFile, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Fatal("Failed to read config file "+cfgPath+". ", err)
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatal("Failed to parse config file "+cfgPath+". ", err)
	}

	return cfg
}

func MustSave(path string, cfg Config) {
	content, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatal("Failed to convert the config to YAML file "+path+". ", err)
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal("Failed to open the config file for update "+path+". ", err)
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		log.Fatal("Failed to write the config file "+path+". ", err)
	}
}
