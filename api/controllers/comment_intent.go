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

// CreateCommentIntentController — "ຢືນຢັນ CF": ກວດ+ຈອງ stock ອັດຕະໂນມັດ, ຄິດ intent_status ເອງ
func CreateCommentIntentController(c echo.Context) error {
	var req requestbody.CommentIntentRequestBody
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.CreateCommentIntentServices(c.Request().Context(), req); err != nil {
		if strings.Contains(err.Error(), "already processed") {
			return c.JSON(http.StatusConflict, presenters.ResponseError("Comment ນີ້ຖືກປະມວນຜົນໄປແລ້ວ", err.Error()))
		}
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("create comment intent error: %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ບໍ່ສາມາດປະມວນຜົນ comment ໄດ້", err.Error()))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

func GetDataCommentIntentController(c echo.Context) error {
	ctx := c.Request().Context()

	var id *int64
	if idParam := c.QueryParam("id"); idParam != "" {
		parsedID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ"))
		}
		id = &parsedID
	}

	var customerID *int
	if cidParam := c.QueryParam("customer_id"); cidParam != "" {
		parsedCID, err := strconv.Atoi(cidParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ customer_id ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນເປັນຕົວເລກ"))
		}
		customerID = &parsedCID
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("limit"))

	items, paginationResult, err := services.GetDataCommentIntentServices(ctx, id, customerID, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("get comment intent error: %v", err)
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
