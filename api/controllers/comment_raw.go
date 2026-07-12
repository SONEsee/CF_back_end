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

func CreateCommentRawController(c echo.Context) error {
	var req requestbody.CommentRawRequestBody
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.CreateCommentRawServices(c.Request().Context(), req); err != nil {
		log.Printf("create comment raw error: %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ບໍ່ສາມາດບັນທຶກ comment ໄດ້", err.Error()))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

func GetDataCommentRawController(c echo.Context) error {
	ctx := c.Request().Context()

	var id *int64
	if idParam := c.QueryParam("id"); idParam != "" {
		parsedID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ"))
		}
		id = &parsedID
	}

	var liveSessionID *int
	if lsidParam := c.QueryParam("live_session_id"); lsidParam != "" {
		parsedLSID, err := strconv.Atoi(lsidParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ live_session_id ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນເປັນຕົວເລກ"))
		}
		liveSessionID = &parsedLSID
	}

	var processed *bool
	if pParam := c.QueryParam("processed"); pParam != "" {
		parsedP, err := strconv.ParseBool(pParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ processed ບໍ່ຖືກຕ້ອງ", "ຕ້ອງເປັນ true ຫຼື false"))
		}
		processed = &parsedP
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("limit"))

	items, paginationResult, err := services.GetDataCommentRawServices(ctx, id, liveSessionID, processed, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("get comment raw error: %v", err)
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
