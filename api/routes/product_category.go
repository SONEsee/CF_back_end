package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesProductCategory(routes *echo.Group) {
	routes.POST("/create", controllers.CreateProductCategoryController)
	routes.GET("/category", controllers.GetDataProductCategoryController)
	routes.PATCH("/category/:id/", controllers.UpdateProductCategoryPatchController)
	routes.DELETE("/category/:id/", controllers.DeleteProductCategoryController)
}
