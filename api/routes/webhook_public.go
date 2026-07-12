package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

// SetRoutesWebhookPublic — route ນອກ AuthMiddleware ໂດຍຕັ້ງໃຈ: Facebook/Line ຍິງເຂົ້າມາໂດຍກົງ
// ບໍ່ມີ JWT ໃຫ້ໃຊ້, ຄວາມປອດໄພອີງໃສ່ signature verification (X-Hub-Signature-256 / X-Line-Signature) ແທນ
func SetRoutesWebhookPublic(routes *echo.Group) {
	routes.GET("/facebook", controllers.FacebookWebhookVerifyController)
	routes.POST("/facebook", controllers.FacebookWebhookReceiveController)
	routes.POST("/line", controllers.LineWebhookReceiveController)
}
