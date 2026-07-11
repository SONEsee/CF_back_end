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

func CreateShopSettingController(c echo.Context) error {
	var req requestbody.ShopSettingRequestBody
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.CreateShopSettingServices(c.Request().Context(), req); err != nil {
		log.Printf("create shop setting error: %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ບໍ່ສາມາດສ້າງຄ່າຕັ້ງໄດ້", err.Error()))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

func GetShopSettingController(c echo.Context) error {
	shopID, err := strconv.Atoi(c.Param("shop_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ shop_id ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ shop_id ເປັນຕົວເລກ"))
	}
	setting, err := services.GetShopSettingServices(c.Request().Context(), shopID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("get shop setting error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດດຶງຂໍ້ມູນໄດ້"))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData("ດຶງຂໍ້ມູນສຳເລັດ", setting))
}

func UpdateShopSettingPatchController(c echo.Context) error {
	shopID, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ shop_id ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ shop_id ເປັນຕົວເລກ"))
	}
	var req requestbody.ShopSettingPatchRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.UpdateShopSettingServicesPatch(c.Request().Context(), shopID, req); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("patch shop setting error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດອັບເດດຂໍ້ມູນໄດ້"))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("ອັບເດດຂໍ້ມູນສຳເລັດ"))
}
