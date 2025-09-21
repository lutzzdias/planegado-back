package handler

import (
	profController "planegado/internal/profit/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	pc := profController.ProfitController{}
	routerGroup.POST("/profit/calculate", pc.CalculateProfit)
}
