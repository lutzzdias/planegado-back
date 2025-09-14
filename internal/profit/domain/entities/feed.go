package entity

import (
	"planegado/internal/profit/domain/enums"
	log "planegado/pkg/logger"
)

type Feed struct {
	PercentageAliveWeight float64
	ExpectedGMD           float64
	PriceKg               float64
	Regime                enums.FeedRegime
}

func (f *Feed) CalcDailyConsumption(avgWeight float64) float64 {
	return avgWeight * (f.PercentageAliveWeight / 100)
}

func (f *Feed) CalcTotalCost(herd *Herd, days int) float64 {
	avgWeight := herd.CalcMediumWeight()
	daily := f.CalcDailyConsumption(avgWeight)
	baseCost := float64(herd.NumberOfAnimals) * daily * f.PriceKg * float64(days)

	var costMultiplier float64
	var regimeDescription string

	switch f.Regime {
	case enums.FeedRegimeOpen:
		costMultiplier = 0.3
		regimeDescription = "Regime Aberto"
	case enums.FeedRegimeConfinedSemi:
		costMultiplier = 1.0
		regimeDescription = "Regime Confinado/Semi-confinado"
	default:
		costMultiplier = 1.0
		regimeDescription = "Regime Não Definido"
	}

	finalCost := baseCost * costMultiplier

	log.Info("Feed Cost Calculation - %s: %.1f%% da ração base aplicado. Custo base: R$ %.2f, Custo final: R$ %.2f",
		regimeDescription,
		costMultiplier*100,
		baseCost,
		finalCost)

	return finalCost
}
