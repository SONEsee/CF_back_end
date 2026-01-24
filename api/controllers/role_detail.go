package controllers

import (
	"net/http"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/api/services"
	"github.com/SONEsee/go-echo/api/validators"
	"github.com/labstack/echo/v4"
)

func CreatedRoleDetailController(c echo.Context) error {
	var req requestbody.RoleDetail
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err := services.CreateRoleDetailServices(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}
