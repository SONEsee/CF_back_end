package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/api/services"
	"github.com/SONEsee/go-echo/api/validators"
	"github.com/labstack/echo/v4"
)

// UpdateOrderStatusController ປ່ຽນ status ອໍເດີ — validate transition path, resolve stock reservation ຖ້າ PAID/CANCELLED
func UpdateOrderStatusController(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາລະບຸ ID ເປັນຕົວເລກ"))
	}
	var req requestbody.OrderStatusRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}

	var changedBy *int
	if uid, ok := c.Get("user_id").(int64); ok {
		v := int(uid)
		changedBy = &v
	}

	if err := services.UpdateOrderStatusServices(c.Request().Context(), id, req.Status, req.Note, changedBy); err != nil {
		if strings.Contains(err.Error(), "invalid order status transition") {
			return c.JSON(http.StatusConflict, presenters.ResponseError("ປ່ຽນສະຖານະບໍ່ໄດ້", err.Error()))
		}
		if strings.Contains(err.Error(), "insufficient stock") {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ສິນຄ້າບໍ່ພຽງພໍ", err.Error()))
		}
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("update order status error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດປ່ຽນສະຖານະໄດ້"))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("ປ່ຽນສະຖານະສຳເລັດ"))
}
