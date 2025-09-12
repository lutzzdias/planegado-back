package entity

type Herd struct {
	NumberOfAnimals int
	TotalWeight     float64
	PriceKg         float64
}

func (h *Herd) CalcMediumWeight() float64 {
	if h.NumberOfAnimals == 0 {
		return 0
	}
	return h.TotalWeight / float64(h.NumberOfAnimals)
}

func (h *Herd) CalcInvestment() *InitialInvestment {
	return &InitialInvestment{
		TotalValue: float64(h.NumberOfAnimals) * h.CalcMediumWeight() * h.PriceKg,
		PerHead:    h.CalcMediumWeight() * h.PriceKg,
		AvgWeight:  h.CalcMediumWeight(),
	}
}
