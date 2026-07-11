package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesShopSetting(routes *echo.Group) {
	routes.POST("/create", controllers.CreateShopSettingController)
	routes.GET("/setting/:shop_id/", controllers.GetShopSettingController)
	routes.PATCH("/setting/:shop_id/", controllers.UpdateShopSettingPatchController)
}
