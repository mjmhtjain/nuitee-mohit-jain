package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type HotelSearchQuery struct {
	CheckIn          string `form:"checkin" binding:"required"`
	CheckOut         string `form:"checkout" binding:"required"`
	Currency         string `form:"currency" binding:"required"`
	GuestNationality string `form:"guestNationality" binding:"required"`
	HotelIds         string `form:"hotelIds"`
	Occupancies      string `form:"occupancies" binding:"required"`
}

type HotelsHandler struct{}

func NewHotelsHandler() *HotelsHandler {
	return &HotelsHandler{}
}

func (h *HotelsHandler) SearchHotels() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query HotelSearchQuery
		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Get the supplier config from header
		supplierConfig := c.GetHeader("x-liteapi-supplier-config")
		if supplierConfig == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "supplier config is required",
			})
			return
		}

		// Here you would typically:
		// 1. Parse the hotel IDs into a slice
		hotelIds := strings.Split(query.HotelIds, ",")

		// 2. Parse the occupancies JSON string
		// 3. Make API calls to your hotel service
		// 4. Process the response

		// For now, return a mock response
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{
				"checkin":          query.CheckIn,
				"checkout":         query.CheckOut,
				"currency":         query.Currency,
				"guestNationality": query.GuestNationality,
				"hotelIds":         hotelIds,
				"occupancies":      query.Occupancies,
			},
		})
	}
}
