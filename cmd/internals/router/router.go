package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/handler"
)

type Router struct {
	engine *gin.Engine
}

func (r *Router) Setup() {
	hotelsHandler := handler.NewHotelsHandler()
	r.engine.GET("/hotels", hotelsHandler.SearchHotels())
}

func NewRouter() *Router {
	return &Router{
		engine: gin.Default(),
	}
}
