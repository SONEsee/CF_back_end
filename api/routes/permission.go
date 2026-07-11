package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesPermission(routes *echo.Group) {
	routes.POST("/create", controllers.CreatePermissionController)
	routes.GET("/permission", controllers.GetDataPermissionController)
	routes.PATCH("/permission/:id/", controllers.UpdatePermissionPatchController)
	routes.DELETE("/permission/:id/", controllers.DeletePermissionController)
}
