package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesDiscount(routes *echo.Group) {
	routes.POST("/create", controllers.CreateDiscountController)
	routes.GET("/discount", controllers.GetDataDiscountController)
	routes.PATCH("/discount/:id/", controllers.UpdateDiscountPatchController)
	routes.DELETE("/discount/:id/", controllers.DeactivateDiscountController)
}
