package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesCommentIntent(routes *echo.Group) {
	routes.POST("/create", controllers.CreateCommentIntentController)
	routes.POST("/parse", controllers.ParseCommentController)
	routes.GET("/intent", controllers.GetDataCommentIntentController)
}
