package utils

import (
	"log"

	"github.com/spf13/viper"
)

var Config *AppConfig

type AppConfig struct {
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	Config = &AppConfig{
		JWTSecret: viper.GetString("JWT_SECRET"),
	}
}
