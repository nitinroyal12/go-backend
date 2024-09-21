package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Keys() Config {
	godotenv.Load()
	return Config{
		Port: getEnv("PORT",":4001"),
	}
}

func getEnv(key, fallback string) string {
	if value,ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}