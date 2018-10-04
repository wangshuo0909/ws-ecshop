package config

import (
	"log"
	"github.com/spf13/viper"
)

var DefaultConfig *Config

type Config struct {
	JWTSecret string
}

func LoadConfig() *Config {
	config := viper.New()
	config.SetConfigName("Config")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error context file: %s \n", err)
	}
	return &Config{
		JWTSecret: config.Get("JWT.secret").(string),
	}
}

func init() {
	DefaultConfig = LoadConfig()
}
