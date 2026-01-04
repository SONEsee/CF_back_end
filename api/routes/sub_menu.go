package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesSubmenu(routes *echo.Group) {
	routes.GET("/suball-menu", controllers.GetSubllMenu)
	routes.POST("/sub-menu", controllers.CreatedSubMeNuController)
	routes.GET("/sub-menu", controllers.GetSubMenuController)
	routes.PUT("/sub-menu/:id/", controllers.UpdateSubMenuController)
	routes.PATCH("/sub-menu/:id/", controllers.UpdateSubMenuControllerPut)
	routes.DELETE("/sub-menu/:id/", controllers.DeleteSebMenuControllers)
}
