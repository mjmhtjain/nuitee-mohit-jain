package internals

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/handler"
)

func Router() *gin.Engine {
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/health", handler.NewHealthHandler().Handle())

	// Define a route group for API endpoints
	api := router.Group("/api")
	{
		// GET endpoint with path parameter
		api.GET("/hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, " + name + "!",
			})
		})

		// POST endpoint example
		api.POST("/submit", func(c *gin.Context) {
			var requestBody struct {
				Message string `json:"message"`
			}

			if err := c.BindJSON(&requestBody); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid request body",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": requestBody.Message,
			})
		})
	}

	return router
}
