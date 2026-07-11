package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesModule(routes *echo.Group) {
	routes.POST("/create", controllers.CreateModuleController)
	routes.GET("/module", controllers.GetDataModuleController)
	routes.PUT("/module/:id/", controllers.UpdateModulePutController)
	routes.PATCH("/module/:id/", controllers.UpdateModulePatchController)
}
