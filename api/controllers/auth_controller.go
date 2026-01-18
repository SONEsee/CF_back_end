// ໃນໄຟລ໌ controllers/auth_controller.go
package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/api/services"
	"github.com/SONEsee/go-echo/api/validators"
	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	// ✅ Parse & Validate
	var req requestbody.UserLoginRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
			err.Error(),
		))
	}

	// ✅ Login
	ctx := c.Request().Context()
	user, err := services.LoginServices(ctx, req)
	if err != nil {
		errMsg := strings.ToLower(err.Error())

		if strings.Contains(errMsg, "invalid") {
			return c.JSON(http.StatusUnauthorized, presenters.ResponseError(
				"ເຂົ້າສູ່ລະບົບບໍ່ສຳເລັດ",
				"ຊື່ຜູ້ໃຊ້ ຫຼື ລະຫັດຜ່ານບໍ່ຖືກຕ້ອງ",
			))
		}

		if strings.Contains(errMsg, "blocked") {
			return c.JSON(http.StatusForbidden, presenters.ResponseError(
				"ບັນຊີຖືກລ໋ອກ",
				"ບັນຊີຂອງທ່ານຖືກລ໋ອກ, ກະລຸນາຕິດຕໍ່ຜູ້ດູແລລະບົບ",
			))
		}

		log.Printf("Login error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດເຂົ້າສູ່ລະບົບໄດ້",
		))
	}

	// ✅ Success
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData(
		"ເຂົ້າສູ່ລະບົບສຳເລັດ",
		user,
	))
}
