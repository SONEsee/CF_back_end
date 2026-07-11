package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesProduct(routes *echo.Group) {
	routes.POST("/create", controllers.CreateProductController)
	routes.GET("/product", controllers.GetDataProductController)
	routes.PATCH("/product/:id/", controllers.UpdateProductPatchController)
	routes.DELETE("/product/:id/", controllers.DeleteProductController)
}
