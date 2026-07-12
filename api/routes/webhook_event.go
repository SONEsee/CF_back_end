package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesWebhookEvent(routes *echo.Group) {
	routes.POST("/create", controllers.CreateWebhookEventController)
	routes.GET("/event", controllers.GetDataWebhookEventController)
	routes.PATCH("/event/:id/processed", controllers.MarkWebhookEventProcessedController)
}
