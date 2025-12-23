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

func GetSubllMenu(c echo.Context) error {
	SubMenu, err := services.GateAllWhitSubmenu(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(SubMenu))
}
func CreatedSubMeNuController(c echo.Context) error {
	var req requestbody.SubMenuRequesBody
	var err error
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	ctx := c.Request().Context()
	err = services.CreatedSubMeNuServiced(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

func GetSubMenuController(c echo.Context) error {
	ctx := c.Request().Context()
	var id *int
	if idParam := c.QueryParam("id"); idParam != "" {
		parsedID, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError(
				"ຮູບແບບ ID ບໍ່ຖືກຕອ້ງ",
				"ກາລຸນາປອ້ນ ID ໃຫ້ເປັນຮູບແບບໂຕເລກ",
			))
		}
		id = &parsedID
	}
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("limit"))
	submenu, paginationResult, err := services.GetSubMenuTotalServices(ctx, id, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}
		log.Printf("failed for get data %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດດຶງຂໍ້ມູນສິດການນຳໃຊ້ໄດ້",
		))
	}
	if paginationResult != nil {
		return c.JSON(http.StatusOK, presenters.ResponseSuccessListData(
			submenu,
			paginationResult.CurrentPage,
			paginationResult.CurrentPageTotalItem,
			paginationResult.TotalItems,
			paginationResult.TotalPage,
		))

	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData(
		"ດືງຂໍ້ມູນສຳເລັດ",
		submenu,
	))
}
