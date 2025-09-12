package entity

type Feed struct {
	PercentageAliveWeight float64
	ExpectedGMD           float64
	PriceKg               float64
}

func (f *Feed) CalcDailyConsumption(avgWeight float64) float64 {
	return avgWeight * f.PercentageAliveWeight
}

func (f *Feed) CalcTotalCost(herd *Herd, days int) float64 {
	avgWeight := herd.CalcMediumWeight()
	daily := f.CalcDailyConsumption(avgWeight)
	return float64(herd.NumberOfAnimals) * daily * f.PriceKg * float64(days)
}
