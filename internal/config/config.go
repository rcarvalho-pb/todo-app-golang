package config

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	Port       int
	SecretKey  string
	HashKey    []byte
	BlockKey   []byte
	CookieName string
}

var EnvConfigs *envConfigs

func InitEvenConfigs() {
	EnvConfigs = LoadEnvVariables()
}

func LoadEnvVariables() (config *envConfigs) {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file: ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
