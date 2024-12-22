package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockService := &mockHotelService{}
	hotelsHandler := NewHotelsHandlerWithService(mockService)
	router.GET("/hotels/search", hotelsHandler.SearchHotels())
	return router
}

func TestSearchHotels(t *testing.T) {
	router := setupRouter()
	validParams := "hotelIds=1234,5678&checkin=2024-12-25&checkout=2024-12-26&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]"
	invalidCheckin := "hotelIds=1234,5678&checkin=asdf&checkout=2024-12-26&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]"
	invalidCheckout := "hotelIds=1234,5678&checkin=2024-12-25&checkout=2024&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]"
	checkoutBeforeCheckin := "hotelIds=1234,5678&checkin=2025-01-16&checkout=2025-01-11&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]"
	invalidHotelID := "hotelIds=asdf,5678&checkin=2024-12-25&checkout=2024-12-26&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]"
	invalidOccupancies := "hotelIds=1234,5678&checkin=2024-12-25&checkout=2024-12-26&occupancies=[{\":[10]}]"

	tests := []struct {
		name           string
		queryParams    string
		supplierConfig string
		expectedCode   int
		expectedError  string
	}{
		{
			name:           "Success case",
			queryParams:    validParams,
			supplierConfig: "test-supplier-config",
			expectedCode:   http.StatusOK,
		},
		{
			name:           "Missing supplier config",
			queryParams:    validParams,
			supplierConfig: "",
			expectedCode:   http.StatusBadRequest,
			expectedError:  "supplier config is required",
		},
		{
			name:           "Invalid check-in date format",
			queryParams:    invalidCheckin,
			supplierConfig: "test-supplier-config",
			expectedCode:   http.StatusBadRequest,
			expectedError:  "check-in date must be in format YYYY-MM-DD",
		},
		{
			name:           "Invalid check-out date format",
			queryParams:    invalidCheckout,
			supplierConfig: "test-supplier-config",
			expectedCode:   http.StatusBadRequest,
			expectedError:  "check-out date must be in format YYYY-MM-DD",
		},
		{
			name:           "Check-out before check-in",
			queryParams:    checkoutBeforeCheckin,
			supplierConfig: "test-supplier-config",
			expectedCode:   http.StatusBadRequest,
			expectedError:  "check-out date must be after check-in date",
		},
		{
			name:           "Invalid hotel ID format",
			queryParams:    invalidHotelID,
			supplierConfig: "test-supplier-config",
			expectedCode:   http.StatusBadRequest,
			expectedError:  "invalid hotel ID format",
		},
		{
			name:           "Invalid occupancies format",
			queryParams:    invalidOccupancies,
			supplierConfig: "test-supplier-config",
			expectedCode:   http.StatusBadRequest,
			expectedError:  "invalid occupancies format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/hotels/search?"+tt.queryParams, nil)
			if tt.supplierConfig != "" {
				req.Header.Set("x-liteapi-supplier-config", tt.supplierConfig)
			}

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedError != "" {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedError, response["error"])
			}
		})
	}
}

// Mock hotel service for testing
type mockHotelService struct{}

func (m *mockHotelService) SearchHotels(params dto.HotelSearchServiceParams) ([]dto.HotelPrice, error) {
	if params.HotelIDs[0] == 1234 {
		return []dto.HotelPrice{
			{
				HotelID:  "1234",
				Currency: "EUR",
				Price:    199.99,
			},
		}, nil
	}

	return nil, nil
}
