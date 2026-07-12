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

// SetDefaultAddressController ຕັ້ງ default address ຂອງ customer — ຮັບ address_id, validate ວ່າເປັນຂອງ customer ນີ້ແທ້
func SetDefaultAddressController(c echo.Context) error {
	customerID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາລະບຸ ID ເປັນຕົວເລກ"))
	}
	var req requestbody.SetDefaultAddressRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.SetCustomerDefaultAddressServices(c.Request().Context(), customerID, req.AddressID); err != nil {
		if strings.Contains(err.Error(), "does not belong to") {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ທີ່ຢູ່ນີ້ບໍ່ແມ່ນຂອງລູກຄ້ານີ້", err.Error()))
		}
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("set default address error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດຕັ້ງທີ່ຢູ່ default ໄດ້"))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("ຕັ້ງທີ່ຢູ່ default ສຳເລັດ"))
}
