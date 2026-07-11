package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesSubmenu(routes *echo.Group) {
	routes.POST("/create", controllers.CreateSubMenuController)
	routes.GET("/sub-menu", controllers.GetDataSubMenuController)
	routes.PUT("/sub-menu/:id/", controllers.UpdateSubMenuPutController)
	routes.PATCH("/sub-menu/:id/", controllers.UpdateSubMenuPatchController)
	routes.DELETE("/sub-menu/:id/", controllers.DeleteSubMenuController)
}
