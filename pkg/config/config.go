package config

import (
	"os"
	"planegado/pkg/logger"
)

type Config struct {
	Port string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config := &Config{
		Port: port,
	}

	logger.Info("Loaded config: %+v", config)
	return config
}
