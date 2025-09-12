package main

import (
	"fmt"
	"log"
	"planegado/internal/investment/controller"
	"planegado/pkg/config"
	"planegado/pkg/logger"
)

func main() {
	cfg := config.Load()
	logger.Info("Starting server on port %s", cfg.Port)
	router := controller.NewRouter()

	addr := ":8080"
	if cfg.Port != "" {
		addr = fmt.Sprintf(":%s", cfg.Port)
	}

	if err := router.Run(addr); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
