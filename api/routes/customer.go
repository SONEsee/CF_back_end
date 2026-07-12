package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesCustomer(routes *echo.Group) {
	routes.POST("/create", controllers.CreateCustomerController)
	routes.GET("/customer", controllers.GetDataCustomerController)
	routes.PATCH("/customer/:id/", controllers.UpdateCustomerPatchController)
	routes.DELETE("/customer/:id/", controllers.DeleteCustomerController)
	routes.PATCH("/customer/:id/default-address", controllers.SetDefaultAddressController)
}
