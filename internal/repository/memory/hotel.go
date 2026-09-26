package memory

import (
	hotel "booking/internal/domain"
	"time"
)

type HotelRepository interface {
	Create(hotel hotel.Hotel) (hotel.Hotel, error)
}

type hotelRepository struct {
	hotels []hotel.Hotel
}

func NewHotelRepository() HotelRepository {
	hotels := make([]hotel.Hotel, 0)
	return &hotelRepository{
		hotels: hotels,
	}
}

func (hr *hotelRepository) Create(hotel hotel.Hotel) (hotel.Hotel, error) {
	hotel.ID = len(hr.hotels)
	hotel.CreatedAt = time.Now()
	hotel.IsActive = true

	hr.hotels = append(hr.hotels, hotel)

	return hr.hotels[len(hr.hotels)-1], nil
}
