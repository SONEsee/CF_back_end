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

	permissionRoutes := echo.Group("/permission", middleware.AuthMiddleware)
	SetRoutesPermission(permissionRoutes)

	// ✅ Zone 3: Product catalog
	productCategoryRoutes := echo.Group("/product-category", middleware.AuthMiddleware)
	SetRoutesProductCategory(productCategoryRoutes)

	productRoutes := echo.Group("/product", middleware.AuthMiddleware)
	SetRoutesProduct(productRoutes)

	productImageRoutes := echo.Group("/product-image", middleware.AuthMiddleware)
	SetRoutesProductImage(productImageRoutes)

	productVariantRoutes := echo.Group("/product-variant", middleware.AuthMiddleware)
	SetRoutesProductVariant(productVariantRoutes)

	inventoryRoutes := echo.Group("/inventory", middleware.AuthMiddleware)
	SetRoutesInventory(inventoryRoutes)

	stockMovementRoutes := echo.Group("/stock-movement", middleware.AuthMiddleware)
	SetRoutesStockMovement(stockMovementRoutes)

	stockReservationRoutes := echo.Group("/stock-reservation", middleware.AuthMiddleware)
	SetRoutesStockReservation(stockReservationRoutes)

	// ✅ Zone 4: Customers
	customerRoutes := echo.Group("/customer", middleware.AuthMiddleware)
	SetRoutesCustomer(customerRoutes)

	customerAddressRoutes := echo.Group("/customer-address", middleware.AuthMiddleware)
	SetRoutesCustomerAddress(customerAddressRoutes)

	// ✅ Zone 5: Order/Payment
	discountRoutes := echo.Group("/discount", middleware.AuthMiddleware)
	SetRoutesDiscount(discountRoutes)

	orderRoutes := echo.Group("/order", middleware.AuthMiddleware)
	SetRoutesOrder(orderRoutes)

	orderItemRoutes := echo.Group("/order-item", middleware.AuthMiddleware)
	SetRoutesOrderItem(orderItemRoutes)

	paymentRoutes := echo.Group("/payment", middleware.AuthMiddleware)
	SetRoutesPayment(paymentRoutes)

	shipmentRoutes := echo.Group("/shipment", middleware.AuthMiddleware)
	SetRoutesShipment(shipmentRoutes)

	refundRoutes := echo.Group("/refund", middleware.AuthMiddleware)
	SetRoutesRefund(refundRoutes)

	// ✅ Zone 6: Social
	socialAccountRoutes := echo.Group("/social-account", middleware.AuthMiddleware)
	SetRoutesSocialAccount(socialAccountRoutes)

	chatConversationRoutes := echo.Group("/chat-conversation", middleware.AuthMiddleware)
	SetRoutesChatConversation(chatConversationRoutes)

	chatMessageRoutes := echo.Group("/chat-message", middleware.AuthMiddleware)
	SetRoutesChatMessage(chatMessageRoutes)

	chatTemplateRoutes := echo.Group("/chat-template", middleware.AuthMiddleware)
	SetRoutesChatTemplate(chatTemplateRoutes)

	webhookEventRoutes := echo.Group("/webhook-event", middleware.AuthMiddleware)
	SetRoutesWebhookEvent(webhookEventRoutes)

	// ✅ Zone 7: Live-session
	liveSessionRoutes := echo.Group("/live-session", middleware.AuthMiddleware)
	SetRoutesLiveSession(liveSessionRoutes)

	liveSessionProductRoutes := echo.Group("/live-session-product", middleware.AuthMiddleware)
	SetRoutesLiveSessionProduct(liveSessionProductRoutes)

	commentRawRoutes := echo.Group("/comment-raw", middleware.AuthMiddleware)
	SetRoutesCommentRaw(commentRawRoutes)

	commentIntentRoutes := echo.Group("/comment-intent", middleware.AuthMiddleware)
	SetRoutesCommentIntent(commentIntentRoutes)
}
