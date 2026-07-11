package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesSubscriptionPlan(routes *echo.Group) {
	routes.POST("/create", controllers.CreateSubscriptionPlanController)
	routes.GET("/plan", controllers.GetDataSubscriptionPlanController)
	routes.PUT("/plan/:id/", controllers.UpdateSubscriptionPlanPutController)
	routes.PATCH("/plan/:id/", controllers.UpdateSubscriptionPlanPatchController)
}
