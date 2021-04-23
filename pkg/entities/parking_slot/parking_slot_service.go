package parking_slot

type ParkingSlotService interface {
	// Index() ([]ParkingSlot, error)
	// Show(parkingSlotID int64) (ParkingSlot, error)
	// Patch(parkingSlotID int64, params UpdateParam) (ParkingSlot, error)
	Create(params CreateParam) (ParkingSlot, error)
}
