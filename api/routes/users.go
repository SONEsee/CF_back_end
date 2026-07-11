package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	jwtpkg "github.com/SONEsee/go-echo/pkg/jwt-pkg/middleware"

	"github.com/labstack/echo/v4"
)

func SetUserRoutes(router *echo.Group) {
	router.GET("/getData", controllers.GetUserController)
	router.POST("/create", controllers.CreateUserController)
	router.PATCH("/user-update/:id", controllers.UpdateUserController)
	router.DELETE("/deleted/:id", controllers.DeactivateUserController)
	router.GET("/me", controllers.UserAuthController, jwtpkg.AuthMiddleware)

}
