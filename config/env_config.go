package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{}

type ConfigInf interface {
	Get(key string) string
}

func New(filenames ...string) ConfigInf {

	if err := godotenv.Load(filenames...); err != nil {
		panic(err.Error())
	}
	return &Config{}
}
func (c *Config) Get(key string) string {

	return os.Getenv(key)
}
