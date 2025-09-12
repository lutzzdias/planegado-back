package controller

import (
	log "planegado/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	log.Info("Initializing router...")
	return gin.Default()
}
