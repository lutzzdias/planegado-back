package entity

// ExpectedProfit resultado do cálculo
// swagger:model ExpectedProfit
type ExpectedProfit struct {
	Investment *InitialInvestment `json:"investment"`
	Feed       *Feed              `json:"feed"`
	Sale       *Sale              `json:"sale"`
	Profit     float64            `json:"profit"`
}

func (ep *ExpectedProfit) CalcProfit(herd *Herd, days int) float64 {
	investment := herd.CalcInvestment()
	feedCost := ep.Feed.CalcTotalCost(herd, days)
	saleValue := ep.Sale.CalcTotalValue(herd)

	ep.Investment = investment
	ep.Profit = saleValue - investment.TotalValue - feedCost
	return ep.Profit
}
