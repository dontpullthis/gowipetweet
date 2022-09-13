package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Auth struct {
	ConsumerKey       string `yaml:"consumer_key"`
	ConsumerKeySecret string `yaml:"consumer_key_secret"`
	Token             string `yaml:"token"`
	TokenSecret       string `yaml:"token_secret"`
}

type Config struct {
	Auth Auth
}

var config Config
var ConfigPath = "config.yaml"

func GetConfig() *Config {
	return &config
}

func MustSave() {
	content, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal("Failed to convert the config to YAML file "+ConfigPath+". ", err)
	}

	f, err := os.Create(ConfigPath)
	if err != nil {
		log.Fatal("Failed to open the config file for update "+ConfigPath+". ", err)
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		log.Fatal("Failed to write the config file "+ConfigPath+". ", err)
	}
}
