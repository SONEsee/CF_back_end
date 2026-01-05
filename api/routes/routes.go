package routes

import "github.com/labstack/echo/v4"

func SetRoutes(echo *echo.Group) {

	userRoutes := echo.Group("/users")
	SetUserRoutes(userRoutes)

	mainMenuRoutes := echo.Group("/main")
	SetRoutesMainMenu(mainMenuRoutes)

	subMenuRoutes := echo.Group("/sub")
	SetRoutesSubmenu(subMenuRoutes)

	taxRoutes := echo.Group("/tax")
	SetRoutesTax(taxRoutes)

	roleRoutes := echo.Group("/role")
	SetRoutesRole(roleRoutes)
	typeMidsine := echo.Group("/type-midsine")
	SetRoutesMidsine(typeMidsine)
}
