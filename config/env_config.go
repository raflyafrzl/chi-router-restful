package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{}

func New(filenames ...string) *Config {

	if err := godotenv.Load(filenames...); err != nil {
		panic(err.Error())
	}
	return &Config{}
}
func (c *Config) Get(key string) string {

	return os.Getenv(key)
}
