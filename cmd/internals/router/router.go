package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/handler"
)

type Router struct {
	engine *gin.Engine
}

func (r *Router) Setup() *gin.Engine {

	// Health endpoint
	r.engine.GET("/health", handler.NewHealthHandler().Handle())

	// hotels GET endpoint
	r.engine.GET("/hotels", handler.NewHotelsHandler().SearchHotels())

	return r.engine
}

func NewRouter() *Router {
	return &Router{
		engine: gin.Default(),
	}
}
