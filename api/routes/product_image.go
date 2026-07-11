package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesProductImage(routes *echo.Group) {
	routes.POST("/create", controllers.CreateProductImageController)
	routes.GET("/image", controllers.GetDataProductImageController)
	routes.PATCH("/image/:id/", controllers.UpdateProductImagePatchController)
	routes.DELETE("/image/:id/", controllers.DeleteProductImageController)
}
