package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/api/services"
	"github.com/SONEsee/go-echo/api/validators"
	"github.com/labstack/echo/v4"
)

// ParseCommentController — ຢືນຢັນ CF ອັດຕະໂນມັດ: parse ຂໍ້ຄວາມຄອມເມັນ, ຈັບຄູ່ສິນຄ້າ+customer ເອງ, ຈອງ stock
// ຖ້າສຳເລັດ (ໃຊ້ແທນ /comment-intent/create ຕອນທີ່ຢາກໃຫ້ລະບົບຈັບຄູ່ອັດຕະໂນມັດ ບໍ່ແມ່ນ staff ເລືອກເອງ)
func ParseCommentController(c echo.Context) error {
	var req requestbody.ParseCommentRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.ParseAndCreateCommentIntentServices(c.Request().Context(), req.CommentRawID); err != nil {
		if strings.Contains(err.Error(), "already processed") {
			return c.JSON(http.StatusConflict, presenters.ResponseError("Comment ນີ້ຖືກປະມວນຜົນໄປແລ້ວ", err.Error()))
		}
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("parse comment error: %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ບໍ່ສາມາດປະມວນຜົນ comment ໄດ້", err.Error()))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}
