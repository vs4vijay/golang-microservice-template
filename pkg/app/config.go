package app

import (
	"os"
)

type Config struct {
	Port string
}

func NewConfig() Config {
	config := Config{
		Port: os.Getenv("PORT"),
	}

	return config
}
