package domain_test

import (
	"planegado/domain"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateInvestment(t *testing.T) {
	inv := domain.NewInvestment()
	err := inv.Validade()

	require.Error(t, err)
}

func TestInvestmentIdIsNotAUuid(t *testing.T) {
	inv := domain.NewInvestment()

	inv.ID = "123"
	inv.Amount = 1000.00
	err := inv.Validade()
	require.Error(t, err)
}

func TestInvestmentValidation(t *testing.T) {
	inv := domain.NewInvestment()

	inv.ID = uuid.NewV4().String()
	inv.Amount = 1000.00
	err := inv.Validade()
	require.Nil(t, err)
}
