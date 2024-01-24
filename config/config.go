package config

import "github.com/tkanos/gonfig"

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Config {
	result := Config{}
	gonfig.GetConf("config/config.json", &result)
	return result
}
