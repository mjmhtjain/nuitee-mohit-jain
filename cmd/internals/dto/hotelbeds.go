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
	Code int    `json:"code"`
	Name string `json:"name"`
	// CategoryCode    string `json:"categoryCode"`
	// CategoryName    string `json:"categoryName"`
	// DestinationCode string `json:"destinationCode"`
	// // DestinationName string `json:"destinationName"`
	// ZoneCode        int    `json:"zoneCode"`
	// ZoneName        string `json:"zoneName"`
	// Latitude        string `json:"latitude"`
	// Longitude       string `json:"longitude"`
	// Rooms           []Room `json:"rooms"`
	MinRate  string `json:"minRate"`
	MaxRate  string `json:"maxRate"`
	Currency string `json:"currency"`
}

// type Room struct {
// 	Code  string `json:"code"`
// 	Name  string `json:"name"`
// 	Rates []Rate `json:"rates"`
// }

// type Rate struct {
// 	RateKey              string               `json:"rateKey"`
// 	RateClass            string               `json:"rateClass"`
// 	RateType             string               `json:"rateType"`
// 	Net                  string               `json:"net"`
// 	Allotment            int                  `json:"allotment"`
// 	RateCommentsID       string               `json:"rateCommentsId,omitempty"`
// 	PaymentType          string               `json:"paymentType"`
// 	Packaging            bool                 `json:"packaging"`
// 	BoardCode            string               `json:"boardCode"`
// 	BoardName            string               `json:"boardName"`
// 	CancellationPolicies []CancellationPolicy `json:"cancellationPolicies"`
// 	Taxes                Taxes                `json:"taxes"`
// 	Rooms                int                  `json:"rooms"`
// 	Adults               int                  `json:"adults"`
// 	Children             int                  `json:"children"`
// }

// type CancellationPolicy struct {
// 	Amount string `json:"amount"`
// 	From   string `json:"from"`
// }

// type Taxes struct {
// 	Taxes       []Tax `json:"taxes"`
// 	AllIncluded bool  `json:"allIncluded"`
// }

// type Tax struct {
// 	Included       bool   `json:"included"`
// 	Amount         string `json:"amount"`
// 	Currency       string `json:"currency"`
// 	ClientAmount   string `json:"clientAmount"`
// 	ClientCurrency string `json:"clientCurrency"`
// }

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
