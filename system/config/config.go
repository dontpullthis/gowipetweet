package config

type Auth struct {
	ConsumerKey       string `yaml:"consumer_key"`
	ConsumerKeySecret string `yaml:"consumer_key_secret"`
	Token             string `yaml:"token"`
	TokenSecret       string `yaml:"token_secret"`
}

type Config struct {
	Auth Auth
}
