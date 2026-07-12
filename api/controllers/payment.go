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

func CreatePaymentController(c echo.Context) error {
	var req requestbody.PaymentRequestBody
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.CreatePaymentServices(c.Request().Context(), req); err != nil {
		log.Printf("create payment error: %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ບໍ່ສາມາດບັນທຶກການຊຳລະໄດ້", err.Error()))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

func GetDataPaymentController(c echo.Context) error {
	ctx := c.Request().Context()

	var id *int
	if idParam := c.QueryParam("id"); idParam != "" {
		parsedID, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ"))
		}
		id = &parsedID
	}

	var orderID *int
	if oidParam := c.QueryParam("order_id"); oidParam != "" {
		parsedOID, err := strconv.Atoi(oidParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ order_id ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນເປັນຕົວເລກ"))
		}
		orderID = &parsedOID
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("limit"))

	items, paginationResult, err := services.GetDataPaymentServices(ctx, id, orderID, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("get payment error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດດຶງຂໍ້ມູນໄດ້"))
	}
	if paginationResult != nil {
		return c.JSON(http.StatusOK, presenters.ResponseSuccessListData(
			items, paginationResult.CurrentPage, paginationResult.CurrentPageTotalItem,
			paginationResult.TotalItems, paginationResult.TotalPage,
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData("ດຶງຂໍ້ມູນສຳເລັດ", items))
}

func VerifyPaymentController(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາລະບຸ ID ເປັນຕົວເລກ"))
	}
	var req requestbody.PaymentVerifyRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.VerifyPaymentServices(c.Request().Context(), id, req); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("verify payment error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດຢືນຢັນການຊຳລະໄດ້"))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("ຢືນຢັນການຊຳລະສຳເລັດ"))
}
