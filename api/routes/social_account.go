package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesSocialAccount(routes *echo.Group) {
	routes.POST("/create", controllers.CreateSocialAccountController)
	routes.GET("/account", controllers.GetDataSocialAccountController)
	routes.PATCH("/account/:id/", controllers.UpdateSocialAccountPatchController)
	routes.DELETE("/account/:id/", controllers.DeactivateSocialAccountController)
}
