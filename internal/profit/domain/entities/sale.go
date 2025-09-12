package entity

import "time"

type Sale struct {
	FinalAvgWeight float64
	PriceKg        float64
	ExpectedDate   time.Time
}

func (s *Sale) CalcTotalValue(herd *Herd) float64 {
	return float64(herd.NumberOfAnimals) * s.FinalAvgWeight * s.PriceKg
}
