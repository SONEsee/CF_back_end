package services

import (
	"context"
	"fmt"

	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbupdate "github.com/SONEsee/go-echo/pkg/db-pkg/db-update"
)

// orderTransitions ກຳນົດວ່າຈາກ status ໃດ ປ່ຽນໄປ status ໃດໄດ້ແດ່ — ຫ້າມຂ້າມຂັ້ນ (ເຊັ່ນ UNPAID→SHIPPED)
var orderTransitions = map[string][]string{
	"UNPAID":                 {"PAYMENT_PENDING_VERIFY", "CANCELLED"},
	"PAYMENT_PENDING_VERIFY": {"PAID", "CANCELLED"},
	"PAID":                   {"PACKING", "CANCELLED"},
	"PACKING":                {"SHIPPED"},
	"SHIPPED":                {},
	"CANCELLED":              {},
}

func isTransitionAllowed(from, to string) bool {
	for _, s := range orderTransitions[from] {
		if s == to {
			return true
		}
	}
	return false
}

// UpdateOrderStatusServices ປ່ຽນ status ຂອງ order ດ້ວຍ state machine — ຖ້າປ່ຽນເປັນ PAID/CANCELLED ຈະ resolve
// stock_reservations ຂອງທຸກ order_item ນຳ (PAID→COMPLETED ຫັກ actual_qty ຈິງ, CANCELLED→EXPIRED ຄືນ stock),
// ບັນທຶກ order_status_logs ອັດຕະໂນມັດ, ທັງໝົດຢູ່ໃນ transaction ດຽວ
func UpdateOrderStatusServices(ctx context.Context, orderID int64, newStatus string, note string, changedBy *int) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)

		order, err := dbquery.GetOrderByIDForUpdate(ctx, db, orderID)
		if err != nil {
			return err
		}
		if !isTransitionAllowed(order.CurrentStatus, newStatus) {
			return fmt.Errorf("invalid order status transition: %s -> %s", order.CurrentStatus, newStatus)
		}

		if newStatus == "PAID" || newStatus == "CANCELLED" {
			items, _, err := dbquery.GetOrderItemDataQuery(ctx, nil, orderIDIntPtr(orderID), nil)
			if err != nil {
				return err
			}
			resolveStatus := "EXPIRED"
			if newStatus == "PAID" {
				resolveStatus = "COMPLETED"
			}
			for _, item := range items {
				reservation, err := dbquery.GetStockReservationByOrderItemIDForUpdate(ctx, db, item.ID)
				if err != nil {
					continue // ບໍ່ມີການຈອງຜູກກັບ item ນີ້ — ຂ້າມໄປ
				}
				if reservation.Status != "HOLDING" {
					continue // ຖືກ resolve ໄປແລ້ວ (ບໍ່ຄວນເກີດຕາມ flow ປົກກະຕິ, ແຕ່ປ້ອງກັນໄວ້)
				}
				if err := resolveStockReservationCore(ctx, db, reservation, resolveStatus, changedBy); err != nil {
					return err
				}
			}
		}

		fromStatus := order.CurrentStatus
		if err := dbupdate.UpdateOrderStatus(ctx, db, orderID, newStatus); err != nil {
			return err
		}
		return dbinserts.CreateOrderStatusLog(ctx, db, int(orderID), &fromStatus, newStatus, "STAFF", changedBy, note)
	})
}

func orderIDIntPtr(id int64) *int {
	v := int(id)
	return &v
}
