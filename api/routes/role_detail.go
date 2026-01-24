package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesRoleDetail(routes *echo.Group) {
	routes.POST("/create", controllers.CreatedRoleDetailController)
}
