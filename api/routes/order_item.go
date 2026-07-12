package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesOrderItem(routes *echo.Group) {
	routes.GET("/order-item", controllers.GetDataOrderItemController)
}
