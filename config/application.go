package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

var ConfigName = getConfigName()

func getConfigName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("ENV")
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/environments")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
