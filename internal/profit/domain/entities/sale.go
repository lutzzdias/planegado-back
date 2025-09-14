package entity

import (
	"planegado/internal/profit/domain/enums"
	log "planegado/pkg/logger"
	"time"
)

type Sale struct {
	FinalAvgWeight  float64
	PriceKg         float64
	ExpectedDate    time.Time
	CalculationType enums.WeightCalculationType
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
