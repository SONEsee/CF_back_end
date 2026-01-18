package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesLogin(routes *echo.Group) {
	routes.POST("/login", controllers.LoginController)

}
