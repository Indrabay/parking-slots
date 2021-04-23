package main

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabay/parking-slots/config/postgres"
	handlerParking "github.com/indrabay/parking-slots/pkg/entities/parking_slot/handler"
	repoParking "github.com/indrabay/parking-slots/pkg/entities/parking_slot/repository/mysql"
	serviceParking "github.com/indrabay/parking-slots/pkg/entities/parking_slot/service"
	"github.com/subosito/gotenv"
)

func main() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}

	postgresDb, err := postgres.NewPostgres()
	if err != nil {
		panic(err)
	}

	parkingRepo := repoParking.NewParkingSlotRepo(postgresDb.DB)
	parkingService := serviceParking.NewParkingSlotService(parkingRepo)
	parkingHandler := handlerParking.NewParkingSlotHandler(parkingService)

	r := gin.Default()
	r.POST("/parking-slots", parkingHandler.ParkingSlotCreate)
	r.GET("/parking-slots/:parkingSlotID", parkingHandler.ParkingSlotShow)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
