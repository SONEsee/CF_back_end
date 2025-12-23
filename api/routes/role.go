package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesRole(routes *echo.Group) {
	routes.POST("/create", controllers.CreateRoleControllers)
	routes.GET("/role", controllers.GetRoleControllers)
	routes.PATCH("/role/:id/", controllers.UpdatedRolrePactController)
	routes.PUT("/role/:id/", controllers.UpdatedRolePutControllers)
	routes.DELETE("/role/:id/", controllers.DeleteRoleController)
}
