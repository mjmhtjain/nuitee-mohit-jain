package service

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

type CurrencyService interface {
	Convert(amount float64, sourceCurr, targetCurr string) (float64, error)
}

type CurrencyServiceImpl struct {
}

func NewCurrencyService() CurrencyService {
	return &CurrencyServiceImpl{}
}

// Convert converts the currency from source to target currency
func (c *CurrencyServiceImpl) Convert(amount float64, sourceCurr, targetCurr string) (float64, error) {
	var val float64 = 0
	source := money.GetCurrency(sourceCurr)
	target := money.GetCurrency(targetCurr)

	if source == nil {
		return val, fmt.Errorf("failed to find Currency by code: %v", sourceCurr)
	}

	if target == nil {
		return val, fmt.Errorf("failed to find Currency by code: %v", targetCurr)
	}

	val = c.globalCurrencyConverter(amount, source, target)

	return val, nil
}

// globalCurrencyConverter converts the currencies: Currently has a 1:1 currency conversion ratio
func (c *CurrencyServiceImpl) globalCurrencyConverter(amount float64, source, target *money.Currency) float64 {
	return amount
}
