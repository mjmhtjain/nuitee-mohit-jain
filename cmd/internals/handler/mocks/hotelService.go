package mocks

import (
	"fmt"

	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"
)

// Mock hotel service for testing
type MockHotelService struct{}

func (m *MockHotelService) SearchHotels(params dto.HotelSearchServiceParams) (dto.HotelSearchServiceResponse, error) {

	// Return error for specific hotel ID
	if params.HotelIDs[0] == 9999 {
		return dto.HotelSearchServiceResponse{}, fmt.Errorf("service error")
	}

	if params.HotelIDs[0] == 1234 {
		return dto.HotelSearchServiceResponse{
			HotelPrices: []dto.HotelPrice{
				{
					HotelID:  "1234",
					Currency: "EUR",
					Price:    199.99,
				},
			},
			SupplierResponse: "response",
			SupplierRequest:  "request",
		}, nil
	}

	return dto.HotelSearchServiceResponse{}, nil
}
