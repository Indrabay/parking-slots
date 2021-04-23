package service

import "github.com/indrabay/parking-slots/pkg/entities/parking_slot"

func (p *ParkingSlotService) Create(params parking_slot.CreateParam) (parking_slot.ParkingSlot, error) {
	parkingSlot := parking_slot.ParkingSlot{}
	parkingSlot.Name = params.Name
	parkingSlot.Price = params.Price

	err := p.repo.Insert(&parkingSlot)
	if err != nil {
		return parkingSlot, err
	}

	return parkingSlot, err
}
