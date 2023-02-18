package config

import (
	"github.com/spf13/viper"
)

type Config struct {
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		// logger.Panic(err.Error())
		return
	}
	err = viper.Unmarshal(&config)
	return
}
