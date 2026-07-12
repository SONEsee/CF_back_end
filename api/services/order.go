package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	dbupdate "github.com/SONEsee/go-echo/pkg/db-pkg/db-update"
	"github.com/SONEsee/go-echo/pkg/pagination"
)

// reservationHoldDuration ຄືໄລຍະເວລາຈອງສະຕັອກໄວ້ລໍຖ້າການຊຳລະເງິນ (ຫຼັງໝົດອາຍຸຕ້ອງໄປ resolve ເປັນ EXPIRED ດ້ວຍມື ຫຼື job ຕ່າງຫາກ — ຍັງບໍ່ໄດ້ເຮັດ auto-expire job ໃນຮອບນີ້)
const reservationHoldDuration = 24 * time.Hour

func generateOrderNumber(shopID int) string {
	return fmt.Sprintf("ORD-%d-%s-%04d", shopID, time.Now().Format("20060102150405"), rand.Intn(10000))
}

// CreateOrderServices — transaction ດຽວ: validate stock+discount, ຄິດໄລ່ຍອດເງິນ, ສ້າງ order+order_items, ຈອງ stock ແຕ່ລະ item, ບັນທຶກ status log ທຳອິດ
func CreateOrderServices(ctx context.Context, req requestbody.OrderRequestBody, createdBy *int) (int64, error) {
	tx := dbpkg.GetTransactionManager()
	var orderID int64
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)

		type lineItem struct {
			variantID int
			qty       int
			price     float64
			subtotal  float64
		}
		var lines []lineItem
		var itemsTotal float64

		for _, in := range req.Items {
			variants, _, err := dbquery.GetProductVariantDataQuery(ctx, &in.ProductVariantID, nil)
			if err != nil {
				return err
			}
			price := variants[0].Price
			subtotal := price * float64(in.BuyQuantity)
			lines = append(lines, lineItem{variantID: in.ProductVariantID, qty: in.BuyQuantity, price: price, subtotal: subtotal})
			itemsTotal += subtotal
		}

		var discountAmount float64
		if req.DiscountID != nil {
			discount, err := dbquery.GetDiscountByIDForUpdate(ctx, db, *req.DiscountID)
			if err != nil {
				return err
			}
			if !discount.IsActive {
				return fmt.Errorf("discount is not active")
			}
			now := time.Now()
			if discount.StartAt != nil && now.Before(*discount.StartAt) {
				return fmt.Errorf("discount is not started yet")
			}
			if discount.EndAt != nil && now.After(*discount.EndAt) {
				return fmt.Errorf("discount has expired")
			}
			if discount.UsageLimit != nil && discount.UsedCount >= *discount.UsageLimit {
				return fmt.Errorf("discount usage limit reached")
			}
			if itemsTotal < discount.MinOrder {
				return fmt.Errorf("order total does not meet discount min_order requirement")
			}
			if discount.DiscountType == "PERCENT" {
				discountAmount = itemsTotal * discount.DiscountValue / 100
			} else {
				discountAmount = discount.DiscountValue
			}
			if discountAmount > itemsTotal {
				discountAmount = itemsTotal
			}
			if err := dbupdate.IncrementDiscountUsage(ctx, db, *req.DiscountID); err != nil {
				return err
			}
		}

		netPayable := itemsTotal - discountAmount + req.ShippingFee

		orderNumber := generateOrderNumber(req.ShopID)
		newOrderID, err := dbinserts.CreateOrder(ctx, db, req.ShopID, req.CustomerID, req.LiveSessionID, req.DiscountID, orderNumber, itemsTotal, discountAmount, req.ShippingFee, netPayable, req.Note)
		if err != nil {
			return err
		}
		orderID = newOrderID

		expiresAt := time.Now().Add(reservationHoldDuration).Format(time.RFC3339)
		for _, ln := range lines {
			orderItemID, err := dbinserts.CreateOrderItem(ctx, db, int(newOrderID), ln.variantID, ln.qty, ln.price, ln.subtotal)
			if err != nil {
				return err
			}
			customerID := req.CustomerID
			orderItemIDInt := int(orderItemID)
			reservationReq := requestbody.StockReservationRequestBody{
				ProductVariantID: ln.variantID,
				CustomerID:       &customerID,
				OrderItemID:      &orderItemIDInt,
				ReservedQty:      ln.qty,
				ExpiresAt:        expiresAt,
			}
			if _, err := createStockReservationCore(ctx, db, reservationReq); err != nil {
				return err
			}
		}

		return dbinserts.CreateOrderStatusLog(ctx, db, int(newOrderID), nil, "UNPAID", "STAFF", createdBy, "ສ້າງອໍເດີ")
	})
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func GetDataOrderServices(ctx context.Context, id *int, page, pageSize int) ([]dbschema.OrderDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetOrderDataQuery(ctx, id, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}

func GetDataOrderItemServices(ctx context.Context, id *int, orderID *int, page, pageSize int) ([]dbschema.OrderItemDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetOrderItemDataQuery(ctx, id, orderID, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}
