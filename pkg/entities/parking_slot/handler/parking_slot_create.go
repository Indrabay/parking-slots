package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	ps "github.com/indrabay/parking-slots/pkg/entities/parking_slot"
)

func (p *ParkingSlotHandler) ParkingSlotCreate(c *gin.Context) {
	var reqParam ps.CreateParam
	err := c.ShouldBindJSON(&reqParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	err = validateParam(reqParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	parkingSlot, err := p.parkingSlotService.Create(reqParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": parkingSlot,
	})
}

func validateParam(param ps.CreateParam) error {
	if param.Name == "" {
		return errors.New("Nama tidak boleh kosong")
	}

	if param.Price <= 0 {
		return errors.New("Price harus lebih dari 0")
	}

	return nil
}
