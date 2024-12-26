package service

import (
	"fmt"

	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/client"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"
)

type HotelService interface {
	SearchHotels(dto.HotelSearchServiceParams) ([]dto.HotelPrice, error)
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

func (h *HotelServiceImpl) SearchHotels(serviceParams dto.HotelSearchServiceParams) ([]dto.HotelPrice, error) {
	res := []dto.HotelPrice{}
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

	response, err := h.client.SearchHotels(&request)
	if err != nil {
		return nil, err
	}

	for _, hotel := range response.Hotels.Hotels {
		var price float64 = 0.0

		price, err = hotel.GetPrice()
		if err != nil {
			return nil, fmt.Errorf("failed to get Price for Hotel: %v", hotel.Code)
		}

		if hotel.Currency != serviceParams.Currency {
			price, err = h.currService.Convert(price, hotel.Currency, serviceParams.Currency)
			if err != nil {
				return nil, fmt.Errorf("failed to convert Currency: %w", err)
			}
		}

		hotelRes := dto.HotelPrice{
			HotelID:  hotel.GetStringifiedHotelCode(),
			Currency: serviceParams.Currency,
			Price:    price,
		}

		res = append(res, hotelRes)
	}

	return res, nil
}
