package service

import ps "github.com/indrabay/parking-slots/pkg/entities/parking_slot"

func (p *ParkingSlotService) Show(parkingSlotID int64) (ps.ParkingSlot, error) {
	var models ps.ParkingSlot
	models, err := p.repo.FindByID(parkingSlotID)
	if err != nil {
		return models, err
	}

	return models, err
}
