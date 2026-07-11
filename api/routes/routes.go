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

	// ✅ Zone 1: Tenant/Billing
	subscriptionPlanRoutes := echo.Group("/subscription-plan", middleware.AuthMiddleware)
	SetRoutesSubscriptionPlan(subscriptionPlanRoutes)

	shopRoutes := echo.Group("/shop", middleware.AuthMiddleware)
	SetRoutesShop(shopRoutes)

	shopSubscriptionRoutes := echo.Group("/shop-subscription", middleware.AuthMiddleware)
	SetRoutesShopSubscription(shopSubscriptionRoutes)

	shopBankAccountRoutes := echo.Group("/shop-bank-account", middleware.AuthMiddleware)
	SetRoutesShopBankAccount(shopBankAccountRoutes)

	shopSettingRoutes := echo.Group("/shop-setting", middleware.AuthMiddleware)
	SetRoutesShopSetting(shopSettingRoutes)

	// ✅ Zone 2: Auth/RBAC
	moduleRoutes := echo.Group("/module", middleware.AuthMiddleware)
	SetRoutesModule(moduleRoutes)

	subMenuRoutes := echo.Group("/sub", middleware.AuthMiddleware)
	SetRoutesSubmenu(subMenuRoutes)

	roleRoutes := echo.Group("/role", middleware.AuthMiddleware)
	SetRoutesRole(roleRoutes)
}
