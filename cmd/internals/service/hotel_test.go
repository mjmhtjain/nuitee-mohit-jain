package service

import (
	"fmt"
	"testing"

	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/client"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"
	"github.com/stretchr/testify/assert"
)

// Mock HotelBeds client
type mockHotelBedsClient struct {
	shouldError bool
	invalidRate bool
}

func (m *mockHotelBedsClient) SearchHotels(request *dto.HotelBedsSearchRequest) (*dto.HotelbedsResponse, error) {
	if m.shouldError {
		return nil, fmt.Errorf("client error")
	}

	minRate := "199.99"
	if m.invalidRate {
		minRate = "invalid"
	}

	return &dto.HotelbedsResponse{
		Hotels: dto.Hotels{
			Hotels: []dto.Hotel{
				{
					Code:     1234,
					MinRate:  minRate,
					Currency: "EUR",
				},
				{
					Code:     5678,
					MinRate:  "299.99",
					Currency: "EUR",
				},
			},
		},
	}, nil
}

type mockCurrencyService struct {
	shouldError bool
}

func (c *mockCurrencyService) Convert(amount float64, sourceCurr, targetCurr string) (float64, error) {
	if c.shouldError {
		return amount, fmt.Errorf("Conversion error")
	}

	return amount, nil
}

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
			client:      &mockHotelBedsClient{},
			currService: &mockCurrencyService{},
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
			currService: &mockCurrencyService{},
			client:      &mockHotelBedsClient{shouldError: true},
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
			client:      &mockHotelBedsClient{invalidRate: true},
			currService: &mockCurrencyService{},
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
			client:      &mockHotelBedsClient{},
			currService: &mockCurrencyService{},
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
			client:      &mockHotelBedsClient{},
			currService: &mockCurrencyService{shouldError: true},
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
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedLen, len(result))
				if len(result) > 0 {
					assert.Equal(t, "1234", result[0].HotelID)
					assert.Equal(t, tt.expectedCurr, result[0].Currency)
					assert.Equal(t, 199.99, result[0].Price)
				}
			}
		})
	}
}
