package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesRefund(routes *echo.Group) {
	routes.POST("/create", controllers.CreateRefundController)
	routes.GET("/refund", controllers.GetDataRefundController)
	routes.PATCH("/refund/:id/status", controllers.UpdateRefundStatusController)
}
