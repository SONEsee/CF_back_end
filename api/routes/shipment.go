package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesShipment(routes *echo.Group) {
	routes.POST("/create", controllers.CreateShipmentController)
	routes.GET("/shipment", controllers.GetDataShipmentController)
	routes.PATCH("/shipment/:id/", controllers.UpdateShipmentPatchController)
	routes.PATCH("/shipment/:id/status", controllers.UpdateShipmentStatusController)
}
