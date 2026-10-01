package http

import (
	"booking/internal/transport/http/hotel"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	hotelHandler hotel.HotelHandler
}

func NewHandlers(
	hotelHandler hotel.HotelHandler,
) *Handlers {
	return &Handlers{
		hotelHandler: hotelHandler,
	}
}

func (h *Handlers) InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	// Объявить endpoints
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
		return
	})

	api := router.Group("/api")
	{
		hotelInitRouters(
			api.Group("/hotel"),
			h.hotelHandler,
		)
	}

	return router
}

// hotelInitRouters объявляет роутеры для hotel
func hotelInitRouters(api *gin.RouterGroup, h hotel.HotelHandler) {
	api.POST("/create", h.HotelCreate)
}
