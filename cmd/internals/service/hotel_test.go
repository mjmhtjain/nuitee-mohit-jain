package service

import (
	"testing"

	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/client"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSearchHotels(t *testing.T) {
	tests := []struct {
		name          string
		client        client.HotelBedsClient
		currService   CurrencyService
		params        dto.HotelSearchServiceParams
		expectedError string
		expectedLen   int
		expectedCurr  string
	}{
		{
			name:        "Success case",
			client:      &mocks.MockHotelBedsClient{},
			currService: &mocks.MockCurrencyService{},
			params: dto.HotelSearchServiceParams{
				CheckIn:  "2024-12-25",
				CheckOut: "2024-12-26",
				HotelIDs: []int{1234, 5678},
				Currency: "EUR",
				Occupancies: []dto.Occupancy{
					{
						Rooms:    1,
						Adults:   2,
						Children: 1,
					},
				},
			},
			expectedLen:  2,
			expectedCurr: "EUR",
		},
		{
			name:        "Client error",
			currService: &mocks.MockCurrencyService{},
			client:      &mocks.MockHotelBedsClient{ShouldError: true},
			params: dto.HotelSearchServiceParams{
				CheckIn:  "2024-12-25",
				CheckOut: "2024-12-26",
				HotelIDs: []int{1234},
				Currency: "EUR",
				Occupancies: []dto.Occupancy{
					{
						Rooms:    1,
						Adults:   2,
						Children: 1,
					},
				},
			},
			expectedError: "client error",
			expectedCurr:  "EUR",
		},
		{
			name:        "Invalid rate parsing",
			client:      &mocks.MockHotelBedsClient{InvalidRate: true},
			currService: &mocks.MockCurrencyService{},
			params: dto.HotelSearchServiceParams{
				CheckIn:  "2024-12-25",
				CheckOut: "2024-12-26",
				HotelIDs: []int{1234},
				Currency: "EUR",
				Occupancies: []dto.Occupancy{
					{
						Rooms:    1,
						Adults:   2,
						Children: 1,
					},
				},
			},
			expectedError: "failed to get Price for Hotel: 1234",
			expectedCurr:  "EUR",
		},
		{
			name:        "Different Currency from ServiceParams",
			client:      &mocks.MockHotelBedsClient{},
			currService: &mocks.MockCurrencyService{},
			params: dto.HotelSearchServiceParams{
				CheckIn:  "2024-12-25",
				CheckOut: "2024-12-26",
				HotelIDs: []int{1234, 5678},
				Currency: "USD",
				Occupancies: []dto.Occupancy{
					{
						Rooms:    1,
						Adults:   2,
						Children: 1,
					},
				},
			},
			expectedLen:  2,
			expectedCurr: "USD",
		},
		{
			name:        "Bad Currency from ServiceParams",
			client:      &mocks.MockHotelBedsClient{},
			currService: &mocks.MockCurrencyService{ShouldError: true},
			params: dto.HotelSearchServiceParams{
				CheckIn:  "2024-12-25",
				CheckOut: "2024-12-26",
				HotelIDs: []int{1234, 5678},
				Currency: "asd",
				Occupancies: []dto.Occupancy{
					{
						Rooms:    1,
						Adults:   2,
						Children: 1,
					},
				},
			},
			expectedError: "failed to convert Currency: Conversion error",
		},
		{
			name:        "Unmarshal Error of Response",
			client:      &mocks.MockHotelBedsClient{InvalidResponse: true},
			currService: &mocks.MockCurrencyService{},
			params: dto.HotelSearchServiceParams{
				CheckIn:  "2024-12-25",
				CheckOut: "2024-12-26",
				HotelIDs: []int{1234, 5678},
				Currency: "EUR",
				Occupancies: []dto.Occupancy{
					{
						Rooms:    1,
						Adults:   2,
						Children: 1,
					},
				},
			},
			expectedError: "failed to unmarshal response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hotelService := &HotelServiceImpl{
				client:      tt.client,
				currService: tt.currService,
			}

			result, err := hotelService.SearchHotels(tt.params)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedLen, len(result.HotelPrices))
				if len(result.HotelPrices) > 0 {
					assert.Equal(t, "1234", result.HotelPrices[0].HotelID)
					assert.Equal(t, tt.expectedCurr, result.HotelPrices[0].Currency)
					assert.Equal(t, 199.99, result.HotelPrices[0].Price)
					assert.NotEmpty(t, result.SupplierRequest)
					assert.NotEmpty(t, result.SupplierResponse)
				}
			}
		})
	}
}
