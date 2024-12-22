package service

import (
	"fmt"
	"strconv"

	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/client"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"
)

type HotelService interface {
	SearchHotels(dto.HotelSearchServiceParams) ([]dto.HotelPrice, error)
}

type HotelServiceImpl struct {
	client client.HotelBedsClient
}

func NewHotelService() HotelService {
	return &HotelServiceImpl{
		client: client.NewHotelBedsClient(),
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

	for _, h := range response.Hotels.Hotels {
		minRate, err := strconv.ParseFloat(h.MinRate, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse MinRate: %w", err)
		}

		hotelRes := dto.HotelPrice{
			HotelID:  fmt.Sprint(h.Code),
			Currency: h.Currency,
			Price:    minRate,
		}

		res = append(res, hotelRes)
	}

	return res, nil
}
