package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("development")
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
