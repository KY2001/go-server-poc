package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var conf Config

func NewConfig() Config {
	err := envconfig.Process("app", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conf
}

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Timeout    int `default:"30"`
	PortNumber int `default:"8000"`
}

type DBConfig struct {
}
