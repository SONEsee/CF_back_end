package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesPayment(routes *echo.Group) {
	routes.POST("/create", controllers.CreatePaymentController)
	routes.GET("/payment", controllers.GetDataPaymentController)
	routes.PATCH("/payment/:id/verify", controllers.VerifyPaymentController)
}
