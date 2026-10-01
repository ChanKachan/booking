package hotel

import (
	hotel "booking/internal/domain"
	"booking/internal/service"
	typesRequest "booking/internal/transport/http"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type HotelHandler interface {
	HotelCreate(c *gin.Context)
}

type hotelHandler struct {
	serviceHotel service.HotelService
}

func NewHotelHandler(serviceHotel service.HotelService) HotelHandler {
	return &hotelHandler{
		serviceHotel: serviceHotel,
	}
}

func (h *hotelHandler) HotelCreate(c *gin.Context) {
	var req hotel.Hotel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			typesRequest.ErrorResponse{
				Error: typesRequest.ApiError{
					Message: "error is parsing body",
					Code:    http.StatusBadRequest,
					Fields:  nil,
				},
			},
		)
		return
	}

	if strings.TrimSpace(req.Name) == "" &&
		strings.TrimSpace(req.Address) == "" &&
		strings.TrimSpace(req.City) == "" {
		c.JSON(
			http.StatusBadRequest,
			typesRequest.ErrorResponse{
				Error: typesRequest.ApiError{
					Message: "hotel request data is required",
					Code:    http.StatusBadRequest,
					Fields:  nil,
				},
			},
		)
		return
	}

	response, err := h.serviceHotel.CreateHotel(req)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			typesRequest.ErrorResponse{
				Error: typesRequest.ApiError{
					Message: "error is parsing body",
					Code:    http.StatusBadRequest,
					Fields:  nil,
				},
			},
		)
		return
	}

	c.JSON(http.StatusOK, response)
	return
}
