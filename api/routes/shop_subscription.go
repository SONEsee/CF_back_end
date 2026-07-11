package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesShopSubscription(routes *echo.Group) {
	routes.POST("/create", controllers.CreateShopSubscriptionController)
	routes.GET("/subscription", controllers.GetDataShopSubscriptionController)
	routes.PATCH("/subscription/:id/", controllers.UpdateShopSubscriptionPatchController)
	routes.PATCH("/subscription/:id/status", controllers.UpdateShopSubscriptionStatusController)
}
