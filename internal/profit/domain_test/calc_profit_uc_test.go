package entities_test

import (
	entity "planegado/internal/profit/domain/entities"
	"planegado/internal/profit/domain/usecase"

	"math"
	"testing"
	"time"
)

func TestExpectedProfitCalculation(t *testing.T) {
	t.Run("calculates expected profit with known inputs", func(t *testing.T) {
		// Arrange
		herd := &entity.Herd{NumberOfAnimals: 10, TotalWeight: 2500, PriceKg: 20}
		feed := &entity.Feed{PercentageAliveWeight: 0.03, ExpectedGMD: 1.5, PriceKg: 2}
		sale := &entity.Sale{FinalAvgWeight: 350, PriceKg: 25, ExpectedDate: time.Now()}
		days := 100

		// Act
		uc := usecase.CalculateExpectedProfit{}
		res := uc.Execute(herd, feed, sale, days)

		t.Logf("avgWeight=%.2f", herd.CalcMediumWeight())
		t.Logf("investment=%.2f", res.Investment.TotalValue)
		t.Logf("feedCost=%.2f", feed.CalcTotalCost(herd, days))
		t.Logf("saleValue=%.2f", sale.CalcTotalValue(herd))
		t.Logf("profit=%.2f", res.Profit)

		// Assert (comparação numérica com epsilon)
		want := 22500.0
		if math.Abs(res.Profit-want) > 1e-6 {
			t.Fatalf("unexpected profit: got=%.2f want=%.2f", res.Profit, want)
		}
	})
}
