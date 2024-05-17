package main

import (
	`log`

	"github.com/labstack/echo/v4"
	`github.com/rhmdnrhuda/simple-rest-api/config`
	`github.com/rhmdnrhuda/simple-rest-api/handlers`
)

func main() {
	r := echo.New()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	hotelHandler := handlers.NewHotelHandler(db)

	r.GET("/hotels", hotelHandler.GetHotels)
	r.POST("/hotel", hotelHandler.CreateHotel)
	r.GET("/hotel/:id", hotelHandler.GetHotelById)
	r.PUT("/hotel/:id", hotelHandler.UpdateHotel)
	r.DELETE("/hotel/:id", hotelHandler.DeleteHotel)

	r.Logger.Fatal(r.Start(":8080"))
}
