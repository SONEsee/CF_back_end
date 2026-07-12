package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

// SetRoutesStockMovement — immutable ledger: ມີແຕ່ create+list, ບໍ່ມີ update/delete
func SetRoutesStockMovement(routes *echo.Group) {
	routes.POST("/create", controllers.CreateStockMovementController)
	routes.GET("/movement", controllers.GetDataStockMovementController)
}
