package mysql

import (
	models "github.com/indrabay/parking-slots/pkg/entities/parking_slot"
	"github.com/jinzhu/gorm"
)

type ParkingSlotRepo struct {
	db *gorm.DB
}

func NewParkingSlotRepo(db *gorm.DB) *ParkingSlotRepo {
	return &ParkingSlotRepo{
		db: db,
	}
}

func (p *ParkingSlotRepo) FindAll() ([]models.ParkingSlot, error) {
	return []models.ParkingSlot{}, nil
}

// result := []models.ParkingSlot{}

// db := p.db.Find(&result)
// if db.Error != nil {
// 	return result, db.Error
// }
// return result, nil
// }

func (p *ParkingSlotRepo) FindByID(parkingSlotID int64) (models.ParkingSlot, error) {
	var parkingSlot models.ParkingSlot
	err := p.db.Where("id = ?", parkingSlotID).First(&parkingSlot).Error
	if err != nil {
		return parkingSlot, err
	}

	return parkingSlot, err
}

func (p *ParkingSlotRepo) Insert(parkingSlot *models.ParkingSlot) error {
	return p.db.Create(parkingSlot).Error
}

func (p *ParkingSlotRepo) Update(parkingSlot *models.ParkingSlot) error {
	return nil
}
