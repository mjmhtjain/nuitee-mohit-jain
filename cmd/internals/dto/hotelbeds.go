package dto

import (
	"fmt"
	"strconv"
)

type HotelbedsResponse struct {
	Hotels Hotels `json:"hotels"`
}

type Hotels struct {
	Hotels   []Hotel `json:"hotels"`
	CheckIn  string  `json:"checkIn"`
	Total    int     `json:"total"`
	CheckOut string  `json:"checkOut"`
}

type Hotel struct {
	Code     int    `json:"code"`
	Name     string `json:"name"`
	MinRate  string `json:"minRate"`
	MaxRate  string `json:"maxRate"`
	Currency string `json:"currency"`
}

func (h *Hotel) GetStringifiedHotelCode() string {
	return fmt.Sprint(h.Code)
}

func (h *Hotel) GetPrice() (float64, error) {
	var price float64 = 0
	price, err := strconv.ParseFloat(h.MinRate, 64)
	if err != nil {
		return price, fmt.Errorf("failed to parse MinRate: %w", err)
	}

	return price, nil
}

type HotelBedsSearchRequest struct {
	Stay        Stay         `json:"stay"`
	Occupancies []Occupancy  `json:"occupancies"`
	Hotels      HotelsFilter `json:"hotels"`
}

type Stay struct {
	CheckIn  string `json:"checkIn"`
	CheckOut string `json:"checkOut"`
}

type Occupancy struct {
	Rooms    int `json:"rooms"`
	Adults   int `json:"adults"`
	Children int `json:"children"`
}

type HotelsFilter struct {
	Hotel []int `json:"hotel"`
}
