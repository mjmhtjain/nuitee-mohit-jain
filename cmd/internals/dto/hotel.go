package dto

// HotelSearchQueryParams represents the query params received in Request
type HotelSearchQueryParams struct {
	CheckIn          string `form:"checkin" binding:"required"`
	CheckOut         string `form:"checkout" binding:"required"`
	Currency         string `form:"currency"`
	GuestNationality string `form:"guestNationality"`
	HotelIds         string `form:"hotelIds" binding:"required"`
	Occupancies      string `form:"occupancies" binding:"required"`
}

// HotelSearchServiceParams represents the request structure for HotelSearch Service
type HotelSearchServiceParams struct {
	CheckIn     string
	CheckOut    string
	HotelIDs    []int
	Currency    string
	Occupancies []Occupancy
}

// HotelSearchServiceResponse represents the response for HotelSearch Service
type HotelSearchServiceResponse struct {
	HotelPrices      []HotelPrice
	SupplierResponse string
	SupplierRequest  string
}

// HotelPriceResponse represents the top-level response structure
type HotelPriceResponse struct {
	Data     []HotelPrice `json:"data"`
	Supplier Supplier     `json:"supplier"`
}

// HotelPrice represents individual hotel price information
type HotelPrice struct {
	HotelID  string  `json:"hotelId"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}

// Supplier contains the request and response details
type Supplier struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}
