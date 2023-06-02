package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver               string `mapstructure:"DB_DRIVER"`
	DBSource               string `mapstructure:"DB_SOURCE"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	TokenExpirationMinutes int    `mapstructure:"TOKEN_EXPIRATION_MINUTES"`
	JWT_SECRET             string `mapstructure:"JWT_SECRET"`
	TimeZone               string `mapstructure:"TIME_ZONE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func LoadTestConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config_test")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
