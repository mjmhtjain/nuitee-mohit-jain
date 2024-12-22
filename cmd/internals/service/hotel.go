package service

import "github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"

type HotelService interface {
	SearchHotels() ([]dto.HotelPrice, error)
}

func NewHotelService() HotelService {
	return &HotelServiceImpl{}
}

type HotelServiceImpl struct {
}

func (h *HotelServiceImpl) SearchHotels() ([]dto.HotelPrice, error) {
	return nil, nil
}
