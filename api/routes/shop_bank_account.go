package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesShopBankAccount(routes *echo.Group) {
	routes.POST("/create", controllers.CreateShopBankAccountController)
	routes.GET("/bank-account", controllers.GetDataShopBankAccountController)
	routes.PATCH("/bank-account/:id/", controllers.UpdateShopBankAccountPatchController)
	routes.DELETE("/bank-account/:id/", controllers.DeactivateShopBankAccountController)
}
