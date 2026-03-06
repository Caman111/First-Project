package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	return Config{
		Key: os.Getenv("X_MASTER_KEY"),
	}, nil
}
