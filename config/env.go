package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

// struct to map env values
type envConfigs struct {
	LocalServerPort string `mapstructure:"LOCAL_SERVER_PORT"`
	SecretKey       string `mapstructure:"SECRET_KEY"`
}

func loadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath(".")

	viper.SetConfigName("app")

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
