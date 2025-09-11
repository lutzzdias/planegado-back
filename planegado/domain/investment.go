package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Investment struct {
	ID        string    `valid:"uuid"`
	Amount    float64   `valid:"notnull"`
	CreatedAt time.Time `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewInvestment() *Investment {
	return &Investment{}
}

func (inv Investment) Validade() error {
	_, err := govalidator.ValidateStruct(inv)

	if err != nil {
		return err
	}
	return nil
}
