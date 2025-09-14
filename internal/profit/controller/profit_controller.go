package controller

import (
	"net/http"
	dto "planegado/internal/profit/domain/entities"
	"planegado/internal/profit/domain/usecase"

	"github.com/gin-gonic/gin"
)

type ProfitController struct{}

// CalculateProfit godoc
// @Summary Calcula o lucro esperado
// @Description Recebe dados de rebanho, ração e venda e retorna o lucro esperado
// @Tags Profit
// @Accept json
// @Produce json
// @Param input body entity.ProfitRequest true "Dados de entrada"
// @Success 200 {object} entity.ExpectedProfit
// @Router /profit/calculate [post]
func (pc *ProfitController) CalculateProfit(c *gin.Context) {
	var req dto.ProfitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecaseCalc := usecase.CalculateExpectedProfit{}
	result := usecaseCalc.Execute(&req.Herd, &req.Feed, &req.Sale, req.Days)
	c.JSON(http.StatusOK, result)
}
