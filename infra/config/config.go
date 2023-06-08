package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Brokers string `json:"BROKERS"`
	Topic   string `json:"TOPIC"`
}

var Env Config

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return envconfig.Process("", &Env)
}
