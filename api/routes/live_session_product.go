package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesLiveSessionProduct(routes *echo.Group) {
	routes.POST("/create", controllers.CreateLiveSessionProductController)
	routes.GET("/product", controllers.GetDataLiveSessionProductController)
	routes.PATCH("/product/:id/", controllers.UpdateLiveSessionProductPatchController)
	routes.DELETE("/product/:id/", controllers.DeleteLiveSessionProductController)
}
