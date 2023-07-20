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

// env: SERVER_XXX
type ServerConfig struct {
	Timeout time.Duration `default:"30s"`
	Port    int           `default:"8080"`
}

// env: DB_XXX
type DBConfig struct {
	User                   string        `default:"root" split_words:"true"`
	Pass                   string        `split_words:"true"`
	Name                   string        `default:"go-server-poc" split_words:"true"`
	InstanceConnectionName string        `split_words:"true"`
	PrivateIP              string        `split_words:"true"`
	MaxOpenConns           int           `default:"25" split_words:"true"`
	MaxIdleConns           int           `default:"50" split_words:"true"`
	ConnMaxLifetime        time.Duration `default:"5m" split_words:"true"`
	ConnMaxIdletime        time.Duration `default:"5m" split_words:"true"`
}
