package controller

import (
	"planegado/internal/profit/domain/usecase"
	"planegado/pkg/logger"
	"planegado/pkg/response"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct {
	sampleUsecase usecase.SampleUsecase
}

func NewSampleHandler(sampleUsecase usecase.SampleUsecase) *SampleHandler {
	return &SampleHandler{sampleUsecase}
}

func (handler *SampleHandler) GetSample(context *gin.Context) {
	logger.Info("GET /sample called")
	value := "test"
	result := handler.sampleUsecase.GetSample(value)
	logger.Info("sampleUsecase executed with result = %s", result.Value)
	response.Success(context, result)
}
