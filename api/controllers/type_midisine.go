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

func CreateTypeMidsineController(c echo.Context) error {
	var req requestbody.TypeMedicine
	var err error
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err = services.CreateTypeMidsine(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}
func GetDataTypeMedicineControllers(c echo.Context) error {
	ctx := c.Request().Context()
	var id *int
	if idParma := c.QueryParam("id"); idParma != "" {
		parsedID, err := strconv.Atoi(idParma)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError(
				"ID ບໍ່ຖືກຕ້ອງ",
				"ກະລຸນາລະບຸ ID ເປັນຕົວເລກ",
			))
		}
		id = &parsedID
	}
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("limit"))
	result, paginationResult, err := services.GetDataTypeMedicineServices(ctx, id, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "not Found") {
			return c.JSON(http.StatusFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}
		log.Printf("Error Get Type Midesine %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດດຶງຂໍ້ມູນປະເພດຢາໄດ້",
		))
	}
	if paginationResult != nil {
		return c.JSON(http.StatusOK, presenters.ResponseSuccessListData(
			result,
			paginationResult.CurrentPage,
			paginationResult.CurrentPageTotalItem,
			paginationResult.TotalItems,
			paginationResult.TotalPage,
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData(
		"ສຳເລັດດືງຂໍ້ມູນ",
		result,
	))
}

func UpdateTypemidsinePutController(c echo.Context) error {
	idParma, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}
	var req requestbody.TypeMedicine
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err = services.UpdateDateTypemidsinePutServices(ctx, idParma, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))

}

func UdateTypeMididesinePatchController(c echo.Context) error {
	idParma, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}
	var req requestbody.TypeMedisinePatch
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err = services.UpdateDateTypemidsinePatchServices(ctx, idParma, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESSS"))
}

func DeletedTypeMisineController(c echo.Context) error {
	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")

	}
	var req requestbody.TypeMedicine
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err = services.DeletedTypeMisineServices(ctx, idParam, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}
