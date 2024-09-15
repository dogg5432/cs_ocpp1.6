package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct{
	Server ServerConfig `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct{
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct{
	Ip string `mapstructture:"ip"`
	Port int `mapstructure:"port"`
}

func Load() (*Config, error){
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}
	return &config, nil
}