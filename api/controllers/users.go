package controllers

import (
	"fmt"
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

func GetUserController(c echo.Context) error {
	ctx := c.Request().Context()
	var id *int
	if idParam := c.QueryParam("id"); idParam != "" {
		parsedID, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ"))
		}
		id = &parsedID
	}
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("limit"))
	q := c.QueryParam("q")

	result, paginationResult, err := services.GetUserService(ctx, id, page, pageSize, q)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("get user error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດດຶງຂໍ້ມູນໄດ້"))
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
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData("ດືງຂໍ້ມູນສຳເລັດ", result))
}

func CreateUserController(c echo.Context) error {
	var req requestbody.UserRequestBody
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", err.Error()))
	}
	if err := services.CreateUserService(c.Request().Context(), req); err != nil {
		log.Printf("create user error: %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ບໍ່ສາມາດສ້າງຜູ້ໃຊ້ໄດ້", err.Error()))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

// DeactivateUserController ໃຊ້ແທນການລົບ (users ບໍ່ມີ deleted_at) — set is_active = false
func DeactivateUserController(c echo.Context) error {
	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ"))
	}
	if err := services.DeactivateUserServices(c.Request().Context(), idParam); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("deactivate user error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດປິດການໃຊ້ງານໄດ້"))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("ປິດການໃຊ້ງານສຳເລັດ"))
}

func UpdateUserController(c echo.Context) error {
	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ",
			"ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ",
		))
	}

	var req requestbody.UserRequestBodyPacth
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
			err.Error(),
		))
	}

	ctx := c.Request().Context()
	err = services.UpdateUserServices(ctx, idParam, req)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				fmt.Sprintf("ບໍ່ພົບຜູ້ໃຊ້ ID: %d", idParam),
			))
		}
		if strings.Contains(strings.ToLower(err.Error()), "already exists") {
			return c.JSON(http.StatusConflict, presenters.ResponseError(
				"Username ຊ້ຳກັນ",
				err.Error(),
			))
		}
		log.Printf("Update user error (ID: %d): %v", idParam, err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດອັບເດດຂໍ້ມູນໄດ້",
		))
	}

	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ອັບເດດຂໍ້ມູນສຳເລັດ",
	))
}

func UserAuthController(c echo.Context) error {
	claims := map[string]interface{}{
		"user_id":   c.Get("user_id"),
		"user_name": c.Get("user_name"),
		"role_id":   c.Get("role_id"),
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(claims))
}
