package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesUpload(routes *echo.Group) {
	routes.POST("/image", controllers.UploadImageController)
}
