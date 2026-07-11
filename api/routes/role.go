package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesRole(routes *echo.Group) {
	routes.POST("/create", controllers.CreateRoleController)
	routes.GET("/role", controllers.GetDataRoleController)
	routes.PUT("/role/:id/", controllers.UpdateRolePutController)
	routes.PATCH("/role/:id/", controllers.UpdateRolePatchController)
	routes.DELETE("/role/:id/", controllers.DeleteRoleController)
}
