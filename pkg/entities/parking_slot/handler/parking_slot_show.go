package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (p *ParkingSlotHandler) ParkingSlotShow(c *gin.Context) {
	parkingSlotID, err := strconv.ParseInt(c.Param("parkingSlotID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	parkingSlot, err := p.parkingSlotService.Show(parkingSlotID)
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
