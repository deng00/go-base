package config

import (
	"bytes"
	"github.com/spf13/viper"
)

type Config = viper.Viper

func GetConfigFromLocal() (*Config, error) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./conf")
	config.AddConfigPath("../conf")
	config.AddConfigPath("../../conf")
	config.AddConfigPath("../../../conf")
	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func GetConfigFromContent(confContent []byte) *Config {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	var err error
	err = config.ReadConfig(bytes.NewBuffer(confContent))
	if err != nil {
		println("go-base-config: failed to parse config, %s", err)
		return nil
	}
	return config
}
