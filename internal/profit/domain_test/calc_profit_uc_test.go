package entities_test

import (
	"math"
	"testing"
	"time"

	entity "planegado/internal/profit/domain/entities"
	"planegado/internal/profit/domain/enums"
	"planegado/internal/profit/domain/usecase"
)

func TestExpectedProfitCalculation(t *testing.T) {
	t.Run("calculates expected profit with known inputs", func(t *testing.T) {
		herd := &entity.Herd{
			NumberOfAnimals:       10,
			TotalWeight:           2500,
			PriceKg:               20,
			WeightCalculationType: enums.WeightCalculationAverage,
		}
		feed := &entity.Feed{
			PercentageAliveWeight: 3.0,
			ExpectedGMD:           1.5,
			PriceKg:               2,
			Regime:                enums.FeedRegimeConfinedSemi,
		}
		sale := &entity.Sale{
			FinalAvgWeight:  350,
			PriceKg:         25,
			ExpectedDate:    time.Now(),
			CalculationType: enums.WeightCalculationAverage,
		}
		days := 100

		uc := usecase.CalculateExpectedProfit{}
		res := uc.Execute(herd, feed, sale, days)

		t.Logf("avgWeight=%.2f", herd.CalcMediumWeight())
		t.Logf("investment=%.2f", res.Investment.TotalValue)
		t.Logf("feedCost=%.2f", feed.CalcTotalCost(herd, days))
		t.Logf("saleValue=%.2f", sale.CalcTotalValue(herd))
		t.Logf("profit=%.2f", res.Profit)

		// Assert - Cálculo para regime confinado:
		// Investment: 10 * 250 * 20 = 50000
		// Feed Cost: 10 * (250 * 0.03) * 2 * 100 * 1.0 = 15000 (regime confinado = 100%)
		// Sale Value: 10 * 350 * 25 = 87500
		// Profit: 87500 - 50000 - 15000 = 22500
		want := 22500.0

		if math.Abs(res.Profit-want) > 1e-6 { // verifica imprecisão de ponto flutuante
			t.Fatalf("unexpected profit: got=%.2f want=%.2f", res.Profit, want)
		}
	})
}
