package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesLiveSession(routes *echo.Group) {
	routes.POST("/create", controllers.CreateLiveSessionController)
	routes.GET("/session", controllers.GetDataLiveSessionController)
	routes.PATCH("/session/:id/", controllers.UpdateLiveSessionPatchController)
	routes.PATCH("/session/:id/end", controllers.EndLiveSessionController)
}
