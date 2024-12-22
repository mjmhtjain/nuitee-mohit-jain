package dto

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
