package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/SONEsee/go-echo/api/services"
	"github.com/labstack/echo/v4"
)

// GetDataOrderItemController — ບໍ່ມີ create ແຍກຕ່າງຫາກ (order_items ຖືກສ້າງເປັນສ່ວນໜຶ່ງຂອງ order creation ເທົ່ານັ້ນ ເພື່ອຮັບປະກັນ stock reservation ຄົບ)
func GetDataOrderItemController(c echo.Context) error {
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

	items, paginationResult, err := services.GetDataOrderItemServices(ctx, id, orderID, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("get order item error: %v", err)
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
