package service

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestNewCurrencyService(t *testing.T) {
	currencyService := NewCurrencyService()
	assert.NotNil(t, currencyService, "CurrencyService instance should not be nil")
}

func TestConvert(t *testing.T) {
	currencyService := NewCurrencyService()

	testCases := []struct {
		desc          string
		amount        float64
		sourceCurr    string
		targetCurr    string
		expectedError string
		expectedValue float64
	}{
		{
			desc:          "Valid currencies",
			amount:        100.0,
			sourceCurr:    "USD",
			targetCurr:    "EUR",
			expectedError: "",
			expectedValue: 100.0,
		},
		{
			desc:          "Invalid source currency",
			amount:        100.0,
			sourceCurr:    "INVALID",
			targetCurr:    "EUR",
			expectedError: "failed to find Currency by code: INVALID",
			expectedValue: 0.0,
		},
		{
			desc:          "Invalid target currency",
			amount:        100.0,
			sourceCurr:    "USD",
			targetCurr:    "INVALID",
			expectedError: "failed to find Currency by code: INVALID",
			expectedValue: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			convertedAmount, err := currencyService.Convert(tc.amount, tc.sourceCurr, tc.targetCurr)

			if tc.expectedError == "" {
				assert.Nil(t, err, "Error should be nil for valid test case")
			} else {
				assert.NotNil(t, err, "Error should not be nil for invalid test case")
				assert.EqualError(t, err, tc.expectedError, "Error message should match expected value")
			}
			assert.Equal(t, tc.expectedValue, convertedAmount, "Converted amount should match expected value")
		})
	}
}

func TestGlobalCurrencyConverter(t *testing.T) {
	currencyService := CurrencyServiceImpl{}

	amount := 100.0
	source := money.GetCurrency("USD")
	target := money.GetCurrency("EUR")

	assert.NotNil(t, source, "Source currency should not be nil")
	assert.NotNil(t, target, "Target currency should not be nil")

	convertedAmount := currencyService.globalCurrencyConverter(amount, source, target)

	assert.Equal(t, amount, convertedAmount, "Converted amount should equal input amount for 1:1 conversion ratio")
}
