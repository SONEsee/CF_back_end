package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesOrder(routes *echo.Group) {
	routes.POST("/create", controllers.CreateOrderController)
	routes.GET("/order", controllers.GetDataOrderController)
	routes.PATCH("/order/:id/status", controllers.UpdateOrderStatusController)
}
