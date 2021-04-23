package handler

import "github.com/indrabay/parking-slots/pkg/entities/parking_slot"

type ParkingSlotHandler struct {
	parkingSlotService parking_slot.ParkingSlotService
}

func NewParkingSlotHandler(psService parking_slot.ParkingSlotService) *ParkingSlotHandler {
	return &ParkingSlotHandler{
		parkingSlotService: psService,
	}
}
