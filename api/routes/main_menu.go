package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesMainMenu(router *echo.Group) {
	router.GET("/data", controllers.GetMainMenuControllers)
	router.GET("/data/all", controllers.GetMainMenuWhitAll)
	router.GET("/test/funsun", controllers.GetMainMenutest)
	router.GET("/suball-menu", controllers.GetSubllMenu)
	router.POST("/create", controllers.CreateMainMenuController)
}
