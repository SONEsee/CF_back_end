package routes

import (
	middleware "github.com/SONEsee/go-echo/pkg/jwt-pkg/middleware"
	"github.com/labstack/echo/v4"
)

func SetRoutes(echo *echo.Group) {
	// ✅ Auth routes - ບໍ່ຕ້ອງການ login
	authRoutes := echo.Group("/auth")
	SetRoutesLogin(authRoutes)

	// ✅ Protected routes - ຕ້ອງການ login
	// ໃຊ້ middleware.AuthMiddleware ກັບທຸກ routes ຂ້າງລຸ່ມນີ້

	userRoutes := echo.Group("/users", middleware.AuthMiddleware)
	SetUserRoutes(userRoutes)

	mainMenuRoutes := echo.Group("/main", middleware.AuthMiddleware)
	SetRoutesMainMenu(mainMenuRoutes)

	subMenuRoutes := echo.Group("/sub", middleware.AuthMiddleware)
	SetRoutesSubmenu(subMenuRoutes)

	taxRoutes := echo.Group("/tax", middleware.AuthMiddleware)
	SetRoutesTax(taxRoutes)

	roleRoutes := echo.Group("/role", middleware.AuthMiddleware)
	SetRoutesRole(roleRoutes)

	typeMidsineRoutes := echo.Group("/type-midsine", middleware.AuthMiddleware)
	SetRoutesMidsine(typeMidsineRoutes)
}
