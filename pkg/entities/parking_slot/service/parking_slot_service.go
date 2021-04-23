package service

import "github.com/indrabay/parking-slots/pkg/entities/parking_slot"

type ParkingSlotService struct {
	repo parking_slot.ParkingSlotRepository
}

func NewParkingSlotService(repo parking_slot.ParkingSlotRepository) *ParkingSlotService {
	return &ParkingSlotService{
		repo: repo,
	}
}
