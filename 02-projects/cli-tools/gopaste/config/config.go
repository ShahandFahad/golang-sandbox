// .env setup
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ApiBaseURL string `mapstructure:"API_BASE_URL"`
	ApiKey     string `mapstructure:"API_KEY"`
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
}
