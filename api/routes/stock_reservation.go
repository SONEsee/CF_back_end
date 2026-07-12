package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesStockReservation(routes *echo.Group) {
	routes.POST("/create", controllers.CreateStockReservationController)
	routes.GET("/reservation", controllers.GetDataStockReservationController)
	routes.PATCH("/reservation/:id/status", controllers.UpdateStockReservationStatusController)
}
