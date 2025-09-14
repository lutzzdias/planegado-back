package usecase

import entity "planegado/internal/profit/domain/entities"

type CalculateExpectedProfit struct{}

func (c *CalculateExpectedProfit) Execute(herd *entity.Herd, feed *entity.Feed, sale *entity.Sale, days int) *entity.ExpectedProfit {
	expectedProfit := &entity.ExpectedProfit{
		Feed: feed,
		Sale: sale,
	}
	expectedProfit.CalcProfit(herd, days)
	return expectedProfit
}
