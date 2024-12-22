package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/dto"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/service"
)

type HotelsHandler struct {
	hotelService service.HotelService
}

func NewHotelsHandler() *HotelsHandler {
	return &HotelsHandler{
		hotelService: service.NewHotelService(),
	}
}

func (h *HotelsHandler) SearchHotels() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query dto.HotelSearchQueryParams
		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Validate date format and values
		checkIn, err := time.Parse("2006-01-02", query.CheckIn)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "check-in date must be in format YYYY-MM-DD",
			})
			return
		}

		checkOut, err := time.Parse("2006-01-02", query.CheckOut)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "check-out date must be in format YYYY-MM-DD",
			})
			return
		}

		// Check if dates are in the future
		now := time.Now().Truncate(24 * time.Hour)
		if checkIn.Before(now) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "check-in date must be in the future",
			})
			return
		}

		if checkOut.Before(now) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "check-out date must be in the future",
			})
			return
		}

		// Check if check-out is after check-in
		if checkOut.Before(checkIn) || checkOut.Equal(checkIn) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "check-out date must be after check-in date",
			})
			return
		}

		// Get the supplier config from header
		supplierConfig := c.GetHeader("x-liteapi-supplier-config")
		if supplierConfig == "" {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "supplier config is required",
				},
			)
			return
		}

		hotelIds := []int{}
		occupancies := []dto.Occupancy{}

		// Parse hotel IDs
		for _, id := range strings.Split(query.HotelIds, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				c.JSON(
					http.StatusBadRequest,
					gin.H{
						"error": "invalid hotel ID format",
					},
				)
				return
			}
			hotelIds = append(hotelIds, i)
		}

		// Parse occupancies from query params
		if err := json.Unmarshal([]byte(query.Occupancies), &occupancies); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "invalid occupancies format",
				},
			)
			return
		}

		serviceParams := dto.HotelSearchServiceParams{
			CheckIn:     query.CheckIn,
			CheckOut:    query.CheckOut,
			HotelIDs:    hotelIds,
			Occupancies: occupancies,
		}

		serviceResponse, err := h.hotelService.SearchHotels(serviceParams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		response := dto.HotelPriceResponse{
			Data: serviceResponse.Data,
			Supplier: dto.Supplier{
				Request: serviceResponse.SupplierRequest,
			},
		}

		// return response
		c.JSON(
			http.StatusOK,
			response,
		)
	}
}
