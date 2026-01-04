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

func UpdateSubMenuController(c echo.Context) error {
	idSubMenu, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ID ບໍ່ຖືກຕ້ອງ",
			"ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ",
		))
	}
	var req requestbody.SubMenuRequesBody
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
			err.Error(),
		))
	}
	err = services.UpdateSubMenuPutServices(c.Request().Context(), idSubMenu, req)
	if err != nil {

		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}
		log.Printf("update submenu error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດອັບເດດຂໍ້ມູນໄດ້",
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ອັບເດດຂໍ້ມູນສຳເລັດ",
	))
}
func UpdateSubMenuControllerPut(c echo.Context) error {
	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຮູບແບບ ID ບໍ່ຖືກຕອ້ງ",
			"ກາລຸນາປອ້ນ ID ເປັນຕົວເລກ",
		))
	}
	var req requestbody.SubMenuRequesBodyPact
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ບໍ່ພົບຂໍ້ມູນ",
			err.Error(),
		))
	}
	err = services.UpdateSubMenuPactServices(c.Request().Context(), idParam, req)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}
		log.Printf("update submenu error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດອັບເດດຂໍ້ມູນໄດ້",
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ອັບເດດຂໍ້ມູນສຳເລັດ",
	))
}

func DeleteSebMenuControllers(c echo.Context) error {
	idSubMenu, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຮູບແບບ id ບໍ່ຖືກຕອ້ງ",
			"ໃຊ້ເປັນຕົວເລກ",
		))
	}
	err = services.DeleteSubMenuServices(c.Request().Context(), idSubMenu)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}

		log.Printf("update submenu error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດລົບຂໍ້ມູນໄດ້",
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ລົບຂໍ້ມູນສຳເລັດ",
	))
}
