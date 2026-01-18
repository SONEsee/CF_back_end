package jwtpkg

import (
	"net/http"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/labstack/echo/v4"
)

// RequireRole middleware ກວດສອບ role
func RequireRole(allowedRoles ...int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ✅ ດຶງ role_id ຈາກ context (ມາຈາກ AuthMiddleware)
			roleID, ok := c.Get("role_id").(int)
			if !ok {
				return c.JSON(http.StatusUnauthorized, presenters.ResponseError(
					"ບໍ່ໄດ້ເຂົ້າສູ່ລະບົບ",
					"ກະລຸນາເຂົ້າສູ່ລະບົບກ່ອນ",
				))
			}

			// ✅ ເຊັກວ່າມີສິດບໍ່
			hasPermission := false
			for _, allowedRole := range allowedRoles {
				if roleID == allowedRole {
					hasPermission = true
					break
				}
			}

			if !hasPermission {
				return c.JSON(http.StatusForbidden, presenters.ResponseError(
					"ບໍ່ມີສິດເຂົ້າເຖິງ",
					"ທ່ານບໍ່ມີສິດໃຊ້ງານຟັງຊັ່ນນີ້",
				))
			}

			return next(c)
		}
	}
}

// RequireOwnerOrAdmin middleware ກວດສອບວ່າເປັນເຈົ້າຂອງ ຫຼື Admin
func RequireOwnerOrAdmin(getUserIDFromParam func(echo.Context) (int64, error)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			currentUserID := c.Get("user_id").(int64)
			currentRoleID := c.Get("role_id").(int)

			// ✅ ຖ້າເປັນ admin (role_id = 1) ຜ່ານເລີຍ
			if currentRoleID == 1 {
				return next(c)
			}

			// ✅ ເຊັກວ່າເປັນເຈົ້າຂອງບໍ່
			targetUserID, err := getUserIDFromParam(c)
			if err != nil {
				return c.JSON(http.StatusBadRequest, presenters.ResponseError(
					"ID ບໍ່ຖືກຕ້ອງ",
					err.Error(),
				))
			}

			if currentUserID != targetUserID {
				return c.JSON(http.StatusForbidden, presenters.ResponseError(
					"ບໍ່ມີສິດເຂົ້າເຖິງ",
					"ທ່ານສາມາດແກ້ໄຂເຉພາະຂໍ້ມູນຂອງທ່ານເອງເທົ່ານັ້ນ",
				))
			}

			return next(c)
		}
	}
}
