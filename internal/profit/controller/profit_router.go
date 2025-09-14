package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	pc := ProfitController{}
	routerGroup.POST("/profit/calculate", pc.CalculateProfit)
}
