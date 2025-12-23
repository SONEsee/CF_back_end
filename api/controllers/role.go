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

func CreateRoleControllers(c echo.Context) error {
	var req requestbody.Role
	var err error
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	ctx := c.Request().Context()
	err = services.CreateRoleServices(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}
func GetRoleControllers(c echo.Context) error {
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
	roles, paginationResult, err := services.GetRoleServices(ctx, id, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}
		log.Printf("not get data for role %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດດຶງຂໍ້ມູນສິດການນຳໃຊ້ໄດ້",
		))
	}
	if paginationResult != nil {
		return c.JSON(http.StatusOK, presenters.ResponseSuccessListData(
			roles,
			paginationResult.CurrentPage,
			paginationResult.CurrentPageTotalItem,
			paginationResult.TotalItems,
			paginationResult.TotalPage,
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData(
		"ດືງຂໍ້ມູນສຳເລັດ",
		roles,
	))

}

func UpdatedRolePutControllers(c echo.Context) error {
	roleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາລະບຸ ID ເປັນຕົວເລກ",
		))

	}
	var req requestbody.Role
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕອ້ງ", err.Error(),
		))
	}
	err = services.UpdatedRolePut(c.Request().Context(), roleID, req)
	if err != nil {
		if strings.Contains(err.Error(), "not Found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຊໍ້ມູນ", err.Error(),
			))
		}
		log.Printf("fail not found data :%v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດອັບເດດຂໍ້ມູນໄດ້",
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ອັບເດດຂໍ້ມູນສຳເລັດ",
	))
}

func UpdatedRolrePactController(c echo.Context) error {
	roleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຮູບແບບ ID ບໍ່ຖືກ",
			"ກາລຸນາປອ້ນ ID ເປັນຕົວເລກ",
		))
	}
	var req requestbody.RolePatchRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕອ້ງ", err.Error(),
		))
	}
	err = services.UpdatedRolePacth(c.Request().Context(), roleID, req)
	if err != nil {
		if strings.Contains(err.Error(), "not Found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຊໍ້ມູນ", err.Error(),
			))
		}
		log.Printf("failed not data role %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດອັບເດດຂໍ້ມູນໄດ້",
		))
	}

	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ອັບເດດຂໍ້ມູນສຳເລັດ",
	))
}
func DeleteRoleController(c echo.Context) error {
	roleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຮູບແບບ ID ບໍ່ຖືກຕອ້ງ",
			"ກາລູນາປອ້ນ ID ເປັນໂຕເລກ",
		))
	}
	err = services.DeletedRoleServices(c.Request().Context(), roleID)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}
		log.Printf("deleted error %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດລຶບຂໍ້ມູນໄດ້",
		))
	}

	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ລືບຊໍ້ມູນສຳເລັດ",
	))
}
