package controller

import (
	"planegado/pkg/logger"

	"planegado/internal/profit/domain/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	logger.Info("Initializing sampleUsecase")
	sampleUsecase := usecase.NewSampleUsecase()

	logger.Info("Initializing sampleHandler")
	sampleHandler := NewSampleHandler(sampleUsecase)

	router.GET("/sample", sampleHandler.GetSample)

	return router
}
