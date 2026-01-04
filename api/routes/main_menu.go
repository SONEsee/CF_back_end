package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesMainMenu(router *echo.Group) {
	router.POST("/create", controllers.CreateMainMenuController)
	router.GET("/main-menu", controllers.GetDataMainMenuControllers)
	router.PUT("/main-menu/:id/", controllers.UpdateMainMenuPutController)
	router.PATCH("/main-menu/:id/", controllers.UpdateMainMenuPacthController)
	router.DELETE("/main-menu/:id/", controllers.DeleteMainMEnuController)

}
