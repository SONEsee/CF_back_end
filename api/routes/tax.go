package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"

	"github.com/labstack/echo/v4"
)

func SetRoutesTax(routes *echo.Group) {
	routes.POST("/creat", controllers.CreateTaxControler)
	routes.GET("/taxbyid", controllers.GetTaxDataByidController)
	routes.PUT("/update/:id/", controllers.UpdateTaxController)
	routes.DELETE("/deleted/:id", controllers.DeletedTax)
}
