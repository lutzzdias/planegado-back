package entity

import (
	"planegado/internal/profit/domain/enums"
	log "planegado/pkg/logger"
)

// swagger:model Herd
type Herd struct {
	NumberOfAnimals       int                         `json:"numberOfAnimals"`
	TotalWeight           float64                     `json:"totalWeight"`
	PriceKg               float64                     `json:"priceKg"`
	WeightCalculationType enums.WeightCalculationType `json:"weightCalculationType"`
}

func (h *Herd) CalcMediumWeight() float64 {
	if h.NumberOfAnimals == 0 {
		return 0
	}
	return h.TotalWeight / float64(h.NumberOfAnimals)
}

func (h *Herd) CalcInvestment() *InitialInvestment {

	if h.WeightCalculationType == enums.WeightCalculationTotal {
		log.Info("Calculating investment using total weight")
		return &InitialInvestment{
			TotalValue: h.TotalWeight * h.PriceKg,
			PerHead:    h.TotalWeight * h.PriceKg / float64(h.NumberOfAnimals),
			AvgWeight:  h.CalcMediumWeight(),
		}
	}

	avgWeight := h.CalcMediumWeight()
	return &InitialInvestment{
		TotalValue: float64(h.NumberOfAnimals) * avgWeight * h.PriceKg,
		PerHead:    avgWeight * h.PriceKg,
		AvgWeight:  avgWeight,
	}
}
