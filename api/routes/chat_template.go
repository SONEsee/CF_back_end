package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesChatTemplate(routes *echo.Group) {
	routes.POST("/create", controllers.CreateChatTemplateController)
	routes.GET("/template", controllers.GetDataChatTemplateController)
	routes.PATCH("/template/:id/", controllers.UpdateChatTemplatePatchController)
	routes.DELETE("/template/:id/", controllers.DeleteChatTemplateController)
}
