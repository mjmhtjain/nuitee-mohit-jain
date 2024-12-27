package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/util"
	"github.com/stretchr/testify/assert"
)

func TestNewHotelBedsClient(t *testing.T) {
	// Setup test environment variables
	t.Setenv("HOTEL_BEDS_BASE_URL", "http://test.com")
	t.Setenv("HOTEL_BEDS_API_KEY", "test-key")
	t.Setenv("HOTEL_BEDS_SECRET", "test-secret")

	client := NewHotelBedsClient()
	impl, ok := client.(*HotelBedsClientImpl)
	assert.True(t, ok)
	assert.Equal(t, "http://test.com", impl.baseURL)
	assert.Equal(t, "test-key", impl.apiKey)
	assert.Equal(t, "test-secret", impl.apiSecret)
	assert.Equal(t, time.Second*10, impl.httpClient.Timeout)
}

func TestSearchHotels(t *testing.T) {
	// Setup mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request headers
		assert.Equal(t, util.ValueApplicationJSON, r.Header.Get(util.HeaderContentType))
		assert.NotEmpty(t, r.Header.Get(util.HeaderApiKey))
		assert.NotEmpty(t, r.Header.Get(util.HeaderSignature))
		assert.Equal(t, util.ValueApplicationJSON, r.Header.Get(util.HeaderAccept))
		assert.Equal(t, util.ValueGzip, r.Header.Get(util.HeaderAcceptEncoding))

		// Mock response
		response := dto.HotelbedsResponse{
			Hotels: dto.Hotels{
				Hotels: []dto.Hotel{
					{
						Code:     123,
						Name:     "Test Hotel",
						MinRate:  "100.00",
						MaxRate:  "200.00",
						Currency: "EUR",
					},
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	// Setup test client
	t.Setenv("HOTEL_BEDS_BASE_URL", mockServer.URL)
	t.Setenv("HOTEL_BEDS_API_KEY", "test-key")
	t.Setenv("HOTEL_BEDS_SECRET", "test-secret")

	client := NewHotelBedsClient()

	// Test search request
	request := &dto.HotelBedsSearchRequest{
		Stay: dto.Stay{
			CheckIn:  "2024-01-01",
			CheckOut: "2024-01-05",
		},
		Occupancies: []dto.Occupancy{
			{
				Rooms:    1,
				Adults:   2,
				Children: 0,
			},
		},
	}

	requestBytes, err := json.Marshal(request)
	assert.NoError(t, err)

	response, err := client.SearchHotels(requestBytes)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	responseData := dto.HotelbedsResponse{}
	err = json.Unmarshal(response, &responseData)
	assert.NoError(t, err)

	assert.Len(t, responseData.Hotels.Hotels, 1)
	assert.Equal(t, 123, responseData.Hotels.Hotels[0].Code)
	assert.Equal(t, "Test Hotel", responseData.Hotels.Hotels[0].Name)
	assert.Equal(t, "100.00", responseData.Hotels.Hotels[0].MinRate)
	assert.Equal(t, "200.00", responseData.Hotels.Hotels[0].MaxRate)
	assert.Equal(t, "EUR", responseData.Hotels.Hotels[0].Currency)
}

func TestSearchHotels_Error(t *testing.T) {
	// Setup mock server that returns an error
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer mockServer.Close()

	// Setup test client
	t.Setenv("HOTEL_BEDS_BASE_URL", mockServer.URL)
	t.Setenv("HOTEL_BEDS_API_KEY", "test-key")
	t.Setenv("HOTEL_BEDS_SECRET", "test-secret")

	client := NewHotelBedsClient()

	requestBytes, err := json.Marshal(&dto.HotelBedsSearchRequest{
		Stay: dto.Stay{
			CheckIn:  "2024-01-01",
			CheckOut: "2024-01-05",
		},
	})
	assert.NoError(t, err)

	response, err := client.SearchHotels(requestBytes)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "API returned non-200 status code")
}

func TestGenerateSignature(t *testing.T) {
	tests := []struct {
		name        string
		client      *HotelBedsClientImpl
		wantErr     bool
		errContains string
	}{
		{
			name: "Success case",
			client: &HotelBedsClientImpl{
				apiKey:    "test-key",
				apiSecret: "test-secret",
			},
			wantErr: false,
		},
		{
			name:        "Missing credentials",
			client:      &HotelBedsClientImpl{},
			wantErr:     true,
			errContains: "environment variables",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signature, err := tt.client.generateSignature()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, signature)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, signature)
			}
		})
	}
}
