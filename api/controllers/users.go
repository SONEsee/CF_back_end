package controllers

import (
	"net/http"
	"strconv"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/api/services"
	"github.com/SONEsee/go-echo/api/validators"

	jwtpkg "github.com/SONEsee/go-echo/pkg/jwt-pkg"
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

func SingTokenController(c echo.Context) error {
	token, err := jwtpkg.SignToken()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, presenters.ResponseSuccess(token))
}

func UserAuthController(c echo.Context) error {
	var user = c.Get("user").(jwt.MapClaims)
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(user))
}
