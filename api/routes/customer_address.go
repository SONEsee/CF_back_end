package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesCustomerAddress(routes *echo.Group) {
	routes.POST("/create", controllers.CreateCustomerAddressController)
	routes.GET("/address", controllers.GetDataCustomerAddressController)
	routes.PATCH("/address/:id/", controllers.UpdateCustomerAddressPatchController)
	routes.DELETE("/address/:id/", controllers.DeleteCustomerAddressController)
}
