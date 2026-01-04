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

func GetDataMainMenuControllers(c echo.Context) error {
	ctx := c.Request().Context()

	// 1. Parse ID (optional)
	var id *int
	if idParam := c.QueryParam("id"); idParam != "" {
		parsedID, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError(
				"ID ບໍ່ຖືກຕ້ອງ",
				"ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ",
			))
		}
		id = &parsedID
	}

	// 2. Parse Pagination Params (ຢູ່ນອກ if block)
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// 3. ເອີ້ນ Service
	mainmenus, paginationResult, err := services.GetDataMainMenuServices(ctx, id, page, pageSize)
	if err != nil {
		// ກໍລະນີບໍ່ພົບຂໍ້ມູນ
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}

		// Error ອື່ນໆ
		log.Printf("failed to get mainmenu data: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດດຶງຂໍ້ມູນໄດ້",
		))
	}

	// 4. ສົ່ງຜົນລັບ
	if paginationResult != nil {
		// ມີ Pagination
		return c.JSON(http.StatusOK, presenters.ResponseSuccessListData(
			mainmenus,
			paginationResult.CurrentPage,
			paginationResult.CurrentPageTotalItem,
			paginationResult.TotalItems,
			paginationResult.TotalPage,
		))
	}

	// ບໍ່ມີ Pagination
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData(
		"ດຶງຂໍ້ມູນສຳເລັດ",
		mainmenus,
	))
}

func UpdateMainMenuPutController(c echo.Context) error {
	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຮຸບແບບບໍ່ຖືກຕອ້ງ",
			"ກາລຸນາປອ້ນເປັນ ID ເປັນໂຕເລກ",
		))
	}
	var req requestbody.MainMenuRequesBody
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
			err.Error(),
		))
	}
	err = services.UpdateMainMenuPutServices(c.Request().Context(), idParam, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}
		log.Printf("failed update main menu services %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດອັບເດດຂໍ້ມູນໄດ້",
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ອັບເດດຂໍ້ມູນສຳເລັດ",
	))
}
func UpdateMainMenuPacthController(c echo.Context) error {
	idMenMenu, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຮຸບແບບບໍ່ຖືກຕອ້ງ",
			"ກາລຸນາປອ້ນເປັນ ID ເປັນໂຕເລກ",
		))
	}
	var req requestbody.MainMenuPatchRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
			err.Error(),
		))
	}
	err = services.UpdateMainMenuServicesPacth(c.Request().Context(), idMenMenu, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}
		log.Printf("failed update main menu services %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດອັບເດດຂໍ້ມູນໄດ້",
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ອັບເດດຂໍ້ມູນສຳເລັດ",
	))
}

func DeleteMainMEnuController(c echo.Context) error {
	idMenMenu, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, presenters.ResponseError(
			"ບໍ່ພົບຂໍ້ມູນ",
			err.Error(),
		))

	}
	var req requestbody.MainMenuPatchRequest
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
			err.Error(),
		))
	}
	err = services.DeletedMainMenuServices(c.Request().Context(), idMenMenu, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				err.Error(),
			))
		}
		log.Printf("failed update main menu services %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດລົບຂໍ້ມູນໄດ້",
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ລົບຂໍ້ມູນສຳເລັດ",
	))
}
