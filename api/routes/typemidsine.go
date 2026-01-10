package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesMidsine(routes *echo.Group) {
	routes.POST("/create", controllers.CreateTypeMidsineController)
	routes.GET("/type-midedine", controllers.GetDataTypeMedicineControllers)
}
