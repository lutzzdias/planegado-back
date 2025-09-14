package entity

import (
	"planegado/internal/profit/domain/enums"
	log "planegado/pkg/logger"
	"time"
)

// swagger:model Sale
type Sale struct {
	FinalAvgWeight  float64                     `json:"finalAvgWeight" example:"350.0"`
	PriceKg         float64                     `json:"priceKg" example:"20.0"`
	ExpectedDate    time.Time                   `json:"expectedDate" example:"2025-09-20T00:00:00Z"`
	CalculationType enums.WeightCalculationType `json:"calculationType" enums:"0,1" example:"0"` // 0=Média, 1=Total
}

func (s *Sale) CalcTotalValue(herd *Herd) float64 {
	var totalValue float64
	var calculationDescription string

	switch s.CalculationType {
	case enums.WeightCalculationAverage:
		totalValue = float64(herd.NumberOfAnimals) * s.FinalAvgWeight * s.PriceKg
		calculationDescription = "Peso Médio por Animal"

	case enums.WeightCalculationTotal:
		totalValue = herd.TotalWeight * s.PriceKg
		calculationDescription = "Peso Total do Rebanho"

	default:
		totalValue = float64(herd.NumberOfAnimals) * s.FinalAvgWeight * s.PriceKg
		calculationDescription = "Peso Médio por Animal (Default)"
	}

	log.Info("Sale Calculation - %s: Valor total calculado: R$ %.2f",
		calculationDescription,
		totalValue)

	return totalValue
}
