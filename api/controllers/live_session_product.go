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

func CreateLiveSessionProductController(c echo.Context) error {
	var req requestbody.LiveSessionProductRequestBody
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.CreateLiveSessionProductServices(c.Request().Context(), req); err != nil {
		log.Printf("create live session product error: %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ບໍ່ສາມາດປັກໝຸດສິນຄ້າໄດ້", err.Error()))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

func GetDataLiveSessionProductController(c echo.Context) error {
	ctx := c.Request().Context()

	var id *int
	if idParam := c.QueryParam("id"); idParam != "" {
		parsedID, err := strconv.Atoi(idParam)
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

	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("limit"))

	items, paginationResult, err := services.GetDataLiveSessionProductServices(ctx, id, liveSessionID, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("get live session product error: %v", err)
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

func UpdateLiveSessionProductPatchController(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາລະບຸ ID ເປັນຕົວເລກ"))
	}
	var req requestbody.LiveSessionProductPatchRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.UpdateLiveSessionProductServicesPatch(c.Request().Context(), id, req); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("patch live session product error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດອັບເດດຂໍ້ມູນໄດ້"))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("ອັບເດດຂໍ້ມູນສຳເລັດ"))
}

// DeleteLiveSessionProductController ລົບແຖວແທ້ (live_session_products ບໍ່ມີ deleted_at)
func DeleteLiveSessionProductController(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ"))
	}
	if err := services.DeleteLiveSessionProductServices(c.Request().Context(), id); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("delete live session product error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດລົບຂໍ້ມູນໄດ້"))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("ລົບຂໍ້ມູນສຳເລັດ"))
}
