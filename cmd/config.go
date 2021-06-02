package main

import (
	"github.com/spf13/viper"

	"github.com/rs/zerolog/log"
)

type Config struct {
	Port int `mapstructure:"port"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	err = viper.Unmarshal(&config)
	return
}
