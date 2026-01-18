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

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	ctx := c.Request().Context()
	var id *int
	if idParma := c.QueryParam("id"); idParma != "" {
		parerdID, err := strconv.Atoi(idParma)
		if err != nil {
			return err
		}
		id = &parerdID
	}
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("limit"))
	result, PaginationResult, err := services.GetUserService(ctx, id, page, pageSize)
	if err != nil {
		return err
	}
	if PaginationResult != nil {
		return c.JSON(http.StatusOK, presenters.ResponseSuccessListData(
			result,
			PaginationResult.CurrentPage,
			PaginationResult.CurrentPageTotalItem,
			PaginationResult.TotalItems,
			PaginationResult.TotalPage,
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData(
		"ດືງຂໍ້ມູນສຳເລັດ",
		result,
	))

}

func CreateUserController(c echo.Context) error {
	var req requestbody.UserRequestBody
	var err error
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = services.CreateUserService(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

// func DeletedesUserController(c echo.Context) error {
// 	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
// 			"ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ",
// 			"ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ",
// 		))
// 	}
// 	ctx := c.Request().Context()
// 	err = services.DeletdedUserServices(ctx, idParam)
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
// 		"ລົບຂໍ້ມູນສຳເລັດ",
// 	))

// }

func DeleteUserController(c echo.Context) error {
	// Parse ID
	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	// Delete User
	ctx := c.Request().Context()
	err = services.DeletdedUserServices(ctx, idParam)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, presenters.ResponseSuccess("DELETE SUCCESS"))
}

func UpdateUserController(c echo.Context) error {
	// ✅ Parse ID
	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ",
			"ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ",
		))
	}

	// ✅ Parse & Validate Body
	var req requestbody.UserRequestBodyPacth
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
			err.Error(),
		))
	}

	// ✅ Update User
	ctx := c.Request().Context()
	err = services.UpdateUserServices(ctx, idParam, req)
	if err != nil {
		// Not found
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				fmt.Sprintf("ບໍ່ພົບຜູ້ໃຊ້ ID: %d", idParam),
			))
		}

		// Duplicate username
		if strings.Contains(strings.ToLower(err.Error()), "already exists") {
			return c.JSON(http.StatusConflict, presenters.ResponseError(
				"Username ຊ້ຳກັນ",
				err.Error(),
			))
		}

		// Other errors
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

// func SingTokenController(c echo.Context) error {
// 	token, err := jwtpkg.SignToken()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, presenters.ResponseSuccess(token))
// }

func UserAuthController(c echo.Context) error {
	var user = c.Get("user").(jwt.MapClaims)
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(user))
}
