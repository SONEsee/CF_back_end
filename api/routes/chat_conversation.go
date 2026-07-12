package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesChatConversation(routes *echo.Group) {
	routes.GET("/conversation", controllers.GetDataChatConversationController)
	routes.PATCH("/conversation/:id/", controllers.UpdateChatConversationPatchController)
	routes.PATCH("/conversation/:id/read", controllers.MarkChatConversationReadController)
}
