package mocks

import "fmt"

type MockCurrencyService struct {
	ShouldError bool
}

func (c *MockCurrencyService) Convert(amount float64, sourceCurr, targetCurr string) (float64, error) {
	if c.ShouldError {
		return amount, fmt.Errorf("Conversion error")
	}

	return amount, nil
}
