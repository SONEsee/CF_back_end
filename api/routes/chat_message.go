package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesChatMessage(routes *echo.Group) {
	routes.POST("/send", controllers.SendChatMessageController)
	routes.GET("/message", controllers.GetDataChatMessageController)
}
