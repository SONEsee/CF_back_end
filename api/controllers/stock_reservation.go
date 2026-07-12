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

func CreateStockReservationController(c echo.Context) error {
	var req requestbody.StockReservationRequestBody
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.CreateStockReservationServices(c.Request().Context(), req); err != nil {
		if strings.Contains(err.Error(), "insufficient stock") {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ສິນຄ້າບໍ່ພຽງພໍ", err.Error()))
		}
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("create stock reservation error: %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ບໍ່ສາມາດຈອງສິນຄ້າໄດ້", err.Error()))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

func GetDataStockReservationController(c echo.Context) error {
	ctx := c.Request().Context()

	var id *int
	if idParam := c.QueryParam("id"); idParam != "" {
		parsedID, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ"))
		}
		id = &parsedID
	}

	var productVariantID *int
	if vidParam := c.QueryParam("product_variant_id"); vidParam != "" {
		parsedVID, err := strconv.Atoi(vidParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ product_variant_id ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນເປັນຕົວເລກ"))
		}
		productVariantID = &parsedVID
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("limit"))

	items, paginationResult, err := services.GetDataStockReservationServices(ctx, id, productVariantID, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("get stock reservation error: %v", err)
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

// UpdateStockReservationStatusController ຢືນຢັນຂາຍ (COMPLETED) ຫຼືປົດການຈອງ (EXPIRED)
func UpdateStockReservationStatusController(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາລະບຸ ID ເປັນຕົວເລກ"))
	}
	var req requestbody.StockReservationStatusRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}

	var updatedBy *int
	if uid, ok := c.Get("user_id").(int64); ok {
		v := int(uid)
		updatedBy = &v
	}

	if err := services.UpdateStockReservationStatusServices(c.Request().Context(), id, req.Status, updatedBy); err != nil {
		if strings.Contains(err.Error(), "already resolved") {
			return c.JSON(http.StatusConflict, presenters.ResponseError("ການຈອງນີ້ຖືກແກ້ໄຂໄປແລ້ວ", err.Error()))
		}
		if strings.Contains(err.Error(), "insufficient stock") {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ສິນຄ້າບໍ່ພຽງພໍ", err.Error()))
		}
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("update stock reservation status error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດປ່ຽນສະຖານະໄດ້"))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("ປ່ຽນສະຖານະສຳເລັດ"))
}
