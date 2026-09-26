package service

import (
	hotel "booking/internal/domain"
	"booking/internal/repository/memory"
	"fmt"
)

type HotelService interface {
	CreateHotel(hotelData hotel.Hotel) (hotel.Hotel, error)
}
type hotelService struct {
	HotelRepository memory.HotelRepository
}

func NewHotelService(HotelRepository memory.HotelRepository) HotelService {
	return &hotelService{
		HotelRepository: HotelRepository,
	}
}

func (hs *hotelService) CreateHotel(hotelData hotel.Hotel) (hotel.Hotel, error) {
	request, err := hs.HotelRepository.Create(hotelData)
	if err != nil {
		return hotel.Hotel{}, fmt.Errorf("create hotel servise error: %w", err)
	}
	return request, nil
}
