package service

import (
	"encoding/json"
	"fmt"

	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/client"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"
)

type HotelService interface {
	SearchHotels(dto.HotelSearchServiceParams) (dto.HotelSearchServiceResponse, error)
}

type HotelServiceImpl struct {
	client      client.HotelBedsClient
	currService CurrencyService
}

func NewHotelService() HotelService {
	return &HotelServiceImpl{
		client:      client.NewHotelBedsClient(),
		currService: NewCurrencyService(),
	}
}

func (h *HotelServiceImpl) SearchHotels(serviceParams dto.HotelSearchServiceParams) (dto.HotelSearchServiceResponse, error) {
	result := dto.HotelSearchServiceResponse{}

	//create request
	request := dto.HotelBedsSearchRequest{
		Stay: dto.Stay{
			CheckIn:  serviceParams.CheckIn,
			CheckOut: serviceParams.CheckOut,
		},
		Occupancies: serviceParams.Occupancies,
		Hotels: dto.HotelsFilter{
			Hotel: serviceParams.HotelIDs,
		},
	}

	// Convert request to JSON
	byteRequest, err := json.Marshal(request)
	if err != nil {
		return result, fmt.Errorf("failed to marshal request: %w", err)
	}

	// get response from client
	byteResponse, err := h.client.SearchHotels(byteRequest)
	if err != nil {
		return result, fmt.Errorf("failed to marshal response: %w", err)
	}

	// unmarshal response
	response := dto.HotelbedsResponse{}
	err = json.Unmarshal(byteResponse, &response)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// get price for each hotel
	for _, hotel := range response.Hotels.Hotels {
		var price float64 = 0.0

		price, err = hotel.GetPrice()
		if err != nil {
			return result, fmt.Errorf("failed to get Price for Hotel: %v", hotel.Code)
		}

		if hotel.Currency != serviceParams.Currency {
			price, err = h.currService.Convert(price, hotel.Currency, serviceParams.Currency)
			if err != nil {
				return result, fmt.Errorf("failed to convert Currency: %w", err)
			}
		}

		hotelRes := dto.HotelPrice{
			HotelID:  hotel.GetStringifiedHotelCode(),
			Currency: serviceParams.Currency,
			Price:    price,
		}

		result.HotelPrices = append(result.HotelPrices, hotelRes)
	}

	result.SupplierResponse = string(byteResponse)
	result.SupplierRequest = string(byteRequest)

	return result, nil
}
