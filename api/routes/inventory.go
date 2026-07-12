package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesInventory(routes *echo.Group) {
	routes.GET("/inventory", controllers.GetDataInventoryController)
	routes.PATCH("/inventory/:id/", controllers.UpdateInventoryPatchController)
}
