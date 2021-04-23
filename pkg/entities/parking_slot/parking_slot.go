package parking_slot

type ParkingSlot struct {
	ID    int64  `json:"id"`
	Price int64  `json:"price"`
	Name  string `json:"name"`
}

type UpdateParam struct {
	Price *int64  `json:"price"`
	Name  *string `json:"name"`
}

type CreateParam struct {
	Price int64  `json:"price"`
	Name  string `json:"name"`
}
