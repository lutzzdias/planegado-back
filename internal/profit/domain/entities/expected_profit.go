package entity

type ExpectedProfit struct {
	Investment *InitialInvestment
	Feed       *Feed
	Sale       *Sale
	Profit     float64
}

func (ep *ExpectedProfit) CalcProfit(herd *Herd, days int) float64 {
	investment := herd.CalcInvestment()
	feedCost := ep.Feed.CalcTotalCost(herd, days)
	saleValue := ep.Sale.CalcTotalValue(herd)

	ep.Investment = investment
	ep.Profit = saleValue - investment.TotalValue - feedCost
	return ep.Profit
}
