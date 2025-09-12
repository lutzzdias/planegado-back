package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func calculateInitialInvestment(animalNumber int, weightPrice float64, averageWeight float64) float64 {
	return (float64(animalNumber) * averageWeight) * weightPrice
}

func main() {
	log.Info("App start")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": calculateInitialInvestment(10, 12.00, 220)})
	})

	router.Run(":8080")
}
