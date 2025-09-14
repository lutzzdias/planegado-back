// @title Profit API
// @version 1.0
// @description API para cálculo de lucro esperado
// @BasePath /api/v1
package main

import (
	"fmt"
	"log"
	_ "planegado/docs"
	profitController "planegado/internal/profit/controller"
	"planegado/pkg/config"
	"planegado/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.Load()
	logger.Info("Starting server on port %s", cfg.Port)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	profitController.RegisterRoutes(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := ":8080"
	if cfg.Port != "" {
		addr = fmt.Sprintf(":%s", cfg.Port)
	}

	if err := router.Run(addr); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
