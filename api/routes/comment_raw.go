package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesCommentRaw(routes *echo.Group) {
	routes.POST("/create", controllers.CreateCommentRawController)
	routes.GET("/comment", controllers.GetDataCommentRawController)
}
