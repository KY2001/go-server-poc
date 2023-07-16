package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var conf Config

func NewConfig() Config {
	err := envconfig.Process("", &conf)
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
	Timeout time.Duration `default:"30s"`
	Port    int           `default:"8080"`
}

type DBConfig struct {
	DBUser                string `split_words:"true"`
	DBPass                string `split_words:"true"`
	DBName                string `split_words:"true"`
	InstanceConnectonName string `split_words:"true"`
	PrivateIP             string `split_words:"true"`
}
