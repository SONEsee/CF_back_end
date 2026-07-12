package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesProductVariant(routes *echo.Group) {
	routes.POST("/create", controllers.CreateProductVariantController)
	routes.GET("/variant", controllers.GetDataProductVariantController)
	routes.PATCH("/variant/:id/", controllers.UpdateProductVariantPatchController)
	routes.DELETE("/variant/:id/", controllers.DeactivateProductVariantController)
}
