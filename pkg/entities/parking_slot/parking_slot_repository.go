package parking_slot

type ParkingSlotRepository interface {
	FindAll() ([]ParkingSlot, error)
	FindByID(id int64) (ParkingSlot, error)
	Insert(*ParkingSlot) error
	Update(*ParkingSlot) error
}
