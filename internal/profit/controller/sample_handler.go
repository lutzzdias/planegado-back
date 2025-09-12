package controller

import (
	log "planegado/pkg/logger"
	"planegado/pkg/response"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct {
}

func NewSampleHandler() *SampleHandler {
	return &SampleHandler{}
}

func (handler *SampleHandler) GetSample(context *gin.Context) {
	log.Info("GET /sample called")

	log.Info("executed with result = empty", "success")
	response.Success(context, gin.H{})
}
