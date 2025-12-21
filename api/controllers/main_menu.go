package controllers

import (
	"net/http"
	"strconv"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/api/services"
	"github.com/SONEsee/go-echo/api/validators"
	"github.com/labstack/echo/v4"
)

func GetMainMenuControllers(c echo.Context) error {

	idParam := c.QueryParam("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	mainMenu, err := services.GetMainMenuByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, presenters.ResponseSuccess(mainMenu))
}

func GetMainMenuWhitAll(c echo.Context) error {
	mainMenu, err := services.GetAllMainMenusService(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(mainMenu))
}

func GetMainMenutest(c echo.Context) error {
	mainMenu, err := services.GetMainTester(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(mainMenu))
}

func CreateMainMenuController(c echo.Context) error {
	var req requestbody.MainMenuRequesBody
	var err error
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err = services.CreateMainMenuServices(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

func CreateMainMenuTestController(c echo.Context) error {
	var req requestbody.MainMenuRequesBody
	var err error
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	err = services.CreateMainMenuServicesTest(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}
