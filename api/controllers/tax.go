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

func CreateTaxControler(c echo.Context) error {
	var req requestbody.TaxRequestBody
	var err error
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err = services.CreateTaxService(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

func GetDataTaxControllers(c echo.Context) error {
	result, err := services.GetDAtaTaxServices(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(result))
}

func GetDataByidController(c echo.Context) error {
	idParam := c.QueryParam("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	result_id, err := services.GetDataByidServices(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(result_id))
}

func GetTaxDataByidController(c echo.Context) error {
	ctx := c.Request().Context()

	var id *int
	if idParam := c.QueryParam("id"); idParam != "" {
		parsedID, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError(
				"ID ບໍ່ຖືກຕ້ອງ",
				"ກະລຸນາລະບຸ ID ເປັນຕົວເລກ",
			))
		}
		id = &parsedID
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	taxes, paginationResult, err := services.GetTaxService(ctx, id, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}

		log.Printf("Error getting tax: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດດຶງຂໍ້ມູນພາສີໄດ້",
		))
	}

	if paginationResult != nil {

		return c.JSON(http.StatusOK, presenters.ResponseSuccessListData(
			taxes,
			paginationResult.CurrentPage,
			paginationResult.CurrentPageTotalItem,
			paginationResult.TotalPage,
			paginationResult.TotalItems,
		))
	}

	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData(
		"ດຶງຂໍ້ມູນສຳເລັດ",
		taxes,
	))
}

func UpdateTaxController(ctx echo.Context) error {

	taxID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID",
		})
	}

	var req requestbody.TaxRequestBody
	if err := validators.ParseAndValidateBody(ctx, &req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err = services.TaxUpdateservices(ctx.Request().Context(), taxID, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated successfully",
		"data":    req,
	})
}
