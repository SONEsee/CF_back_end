package jwtpkg

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/SONEsee/go-echo/pkg/jwt-pkg/auth"
)

// AuthMiddleware ກວດສອບວ່າມີການ login ແລ້ວບໍ່
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ✅ ດຶງ Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, presenters.ResponseError(
				"ບໍ່ໄດ້ເຂົ້າສູ່ລະບົບ",
				"ກະລຸນາເຂົ້າສູ່ລະບົບກ່ອນ",
			))
		}

		// ✅ ແຍກ Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, presenters.ResponseError(
				"Token ບໍ່ຖືກຕ້ອງ",
				"ຮູບແບບ token ຜິດພາດ (ຕ້ອງເປັນ Bearer token)",
			))
		}

		tokenString := parts[1]

		// ✅ Validate token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				return c.JSON(http.StatusUnauthorized, presenters.ResponseError(
					"Token ໝົດອາຍຸ",
					"ກະລຸນາເຂົ້າສູ່ລະບົບໃໝ່",
				))
			}
			return c.JSON(http.StatusUnauthorized, presenters.ResponseError(
				"Token ບໍ່ຖືກຕ້ອງ",
				"Token ບໍ່ສາມາດໃຊ້ງານໄດ້",
			))
		}

		// ✅ ເກັບຂໍ້ມູນໄວ້ໃນ context
		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.UserName)
		c.Set("role_id", claims.RoleID)

		// ສືບຕໍ່ໄປຫາ handler ຕໍ່ໄປ
		return next(c)
	}
}
