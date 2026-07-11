package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesShop(routes *echo.Group) {
	routes.POST("/create", controllers.CreateShopController)
	routes.GET("/shop", controllers.GetDataShopController)
	routes.PUT("/shop/:id/", controllers.UpdateShopPutController)
	routes.PATCH("/shop/:id/", controllers.UpdateShopPatchController)
	routes.PATCH("/shop/:id/status", controllers.UpdateShopStatusController)
}
