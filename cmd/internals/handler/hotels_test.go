package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/handler/mockService"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockService := &mockService.MockHotelService{}
	hotelsHandler := NewHotelsHandlerWithService(mockService)
	router.GET("/hotels/search", hotelsHandler.SearchHotels())
	return router
}

func TestSearchHotels(t *testing.T) {
	router := setupRouter()
	today := time.Now()
	checkinDate := today.AddDate(0, 0, 1).Format("2006-01-02")
	checkoutDate := today.AddDate(0, 0, 2).Format("2006-01-02")
	badCheckoutDate := today.AddDate(0, 0, 0).Format("2006-01-02")

	validParams := fmt.Sprintf("hotelIds=1234,5678&checkin=%s&checkout=%s&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]&currency=EUR", checkinDate, checkoutDate)
	missingCurrency := fmt.Sprintf("hotelIds=1234,5678&checkin=%s&checkout=%s&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]", checkinDate, checkoutDate)
	invalidCheckin := fmt.Sprintf("hotelIds=1234,5678&checkin=asdf&checkout=%s&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]&currency=EUR", checkoutDate)
	invalidCheckout := fmt.Sprintf("hotelIds=1234,5678&checkin=%s&checkout=2024&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]&currency=EUR", checkinDate)
	checkoutBeforeCheckin := fmt.Sprintf("hotelIds=1234,5678&checkin=%s&checkout=%s&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]&currency=EUR", checkinDate, badCheckoutDate)
	invalidHotelID := fmt.Sprintf("hotelIds=asdf,5678&checkin=%s&checkout=%s&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]&currency=EUR", checkinDate, checkoutDate)
	invalidOccupancies := fmt.Sprintf("hotelIds=1234,5678&checkin=%s&checkout=%s&occupancies=[{\":[10]}]&currency=EUR", checkinDate, checkoutDate)
	downstreamErr := fmt.Sprintf("hotelIds=9999,5678&checkin=%s&checkout=%s&occupancies=[{\"adults\":2,\"children\":1,\"childrenAges\":[10]}]&currency=EUR", checkinDate, checkoutDate)

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
			name:           "Missing Currency",
			queryParams:    missingCurrency,
			supplierConfig: "test-supplier-config",
			expectedCode:   http.StatusBadRequest,
			expectedError:  "currency is required",
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
		{
			name:           "Service layer error",
			queryParams:    downstreamErr,
			supplierConfig: "test-supplier-config",
			expectedCode:   http.StatusInternalServerError,
			expectedError:  "service error",
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
