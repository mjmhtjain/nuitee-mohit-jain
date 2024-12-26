package mocks

import (
	"encoding/json"
	"fmt"

	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"
)

// Mock HotelBeds client
type MockHotelBedsClient struct {
	ShouldError     bool
	InvalidRate     bool
	InvalidResponse bool
}

func (m *MockHotelBedsClient) SearchHotels(request []byte) ([]byte, error) {
	if m.ShouldError {
		return nil, fmt.Errorf("client error")
	}

	minRate := "199.99"
	if m.InvalidRate {
		minRate = "invalid"
	}

	if m.InvalidResponse {
		return []byte("asdfa"), nil
	}

	result := dto.HotelbedsResponse{
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
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return jsonResult, nil
}
