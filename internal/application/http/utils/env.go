package utils

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppPort string `mapstructure:"APP_PORT"`
	DBName  string `mapstructure:"DB_NAME"`
	DBUrl   string `mapstructure:"DB_URL"`
}

func NewEnv() *Env {
	var env Env
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("can't find the env file: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("environment can't be loaded: ", err)
	}

	return &env
}
