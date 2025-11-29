package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Storage struct {
		Type string `mapstructure:"type"`
		Path string `mapstructure:"path"`
	} `mapstructure:"storage"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}
