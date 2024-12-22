package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

		// Here you would typically:
		// 1. Parse the hotel IDs into a slice
		// hotelIds := strings.Split(query.HotelIds, ",")

		// 2. Parse the occupancies JSON string
		// 3. Make API calls to your hotel service
		// 4. Process the response

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

		hotels, err := h.hotelService.SearchHotels(serviceParams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		response := dto.HotelPriceResponse{
			Data:     hotels,
			Supplier: dto.Supplier{},
		}

		// return response
		c.JSON(
			http.StatusOK,
			response,
		)
	}
}
