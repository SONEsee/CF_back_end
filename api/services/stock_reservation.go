package services

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	dbupdate "github.com/SONEsee/go-echo/pkg/db-pkg/db-update"
	"github.com/SONEsee/go-echo/pkg/pagination"
)

// createStockReservationCore ຫັກ available_qty + ບວກ reserved_qty ແລ້ວບັນທຶກແຖວຈອງ — ບໍ່ມີ transaction ຂອງມັນເອງ,
// ໃຫ້ caller (CreateStockReservationServices ຫຼື order creation flow) ຄຸ້ມ transaction ເອງ ເພື່ອໃຫ້ລວມເຂົ້າ transaction ໃຫຍ່ດຽວກັນໄດ້
func createStockReservationCore(ctx context.Context, db dbpkg.DBTX, req requestbody.StockReservationRequestBody) (int64, error) {
	inv, err := dbquery.GetInventoryByVariantIDForUpdate(ctx, db, req.ProductVariantID)
	if err != nil {
		return 0, err
	}

	newAvailableQty := inv.AvailableQty - req.ReservedQty
	if newAvailableQty < 0 {
		return 0, fmt.Errorf("insufficient stock: available_qty ມີພຽງ %d, ຂໍຈອງ %d", inv.AvailableQty, req.ReservedQty)
	}
	newReservedQty := inv.ReservedQty + req.ReservedQty

	if err := dbupdate.UpdateInventoryQty(ctx, db, req.ProductVariantID, inv.ActualQty, newAvailableQty); err != nil {
		return 0, err
	}
	if err := dbupdate.UpdateInventoryReservedQty(ctx, db, req.ProductVariantID, newReservedQty); err != nil {
		return 0, err
	}

	return dbinserts.CreateStockReservation(ctx, db, req)
}

// resolveStockReservationCore ຈັດການ COMPLETED/EXPIRED ຂອງແຖວຈອງດຽວ — ບໍ່ມີ transaction ຂອງມັນເອງ, ໃຊ້ຊ້ຳລະຫວ່າງ
// UpdateStockReservationStatusServices (manual) ແລະ order status-transition flow (Step 3)
func resolveStockReservationCore(ctx context.Context, db dbpkg.DBTX, reservation *dbschema.StockReservationDBSchema, status string, updatedBy *int) error {
	if reservation.Status != "HOLDING" {
		return fmt.Errorf("stock reservation %d already resolved (status=%s)", reservation.ID, reservation.Status)
	}

	inv, err := dbquery.GetInventoryByVariantIDForUpdate(ctx, db, reservation.ProductVariantID)
	if err != nil {
		return err
	}
	newReservedQty := inv.ReservedQty - reservation.ReservedQty
	if newReservedQty < 0 {
		newReservedQty = 0
	}

	if status == "COMPLETED" {
		newActualQty := inv.ActualQty - reservation.ReservedQty
		if newActualQty < 0 {
			return fmt.Errorf("insufficient stock: actual_qty ຈະຕິດລົບ ຕອນຢືນຢັນການຈອງ")
		}
		if err := dbupdate.UpdateInventoryQty(ctx, db, reservation.ProductVariantID, newActualQty, inv.AvailableQty); err != nil {
			return err
		}
		if err := dbupdate.UpdateInventoryReservedQty(ctx, db, reservation.ProductVariantID, newReservedQty); err != nil {
			return err
		}
		reservationID64 := int64(reservation.ID)
		movementReq := requestbody.StockMovementRequestBody{
			ProductVariantID: reservation.ProductVariantID,
			MovementType:     "OUT",
			QtyChange:        -reservation.ReservedQty,
			RefType:          "stock_reservation",
			RefID:            &reservationID64,
			Note:             "ຢືນຢັນການຂາຍຈາກສະຕັອກທີ່ຈອງໄວ້",
		}
		if err := dbinserts.CreateStockMovement(ctx, db, movementReq, newActualQty, updatedBy); err != nil {
			return err
		}
	} else {
		newAvailableQty := inv.AvailableQty + reservation.ReservedQty
		if err := dbupdate.UpdateInventoryQty(ctx, db, reservation.ProductVariantID, inv.ActualQty, newAvailableQty); err != nil {
			return err
		}
		if err := dbupdate.UpdateInventoryReservedQty(ctx, db, reservation.ProductVariantID, newReservedQty); err != nil {
			return err
		}
	}

	return dbupdate.UpdateStockReservationStatus(ctx, db, int64(reservation.ID), status)
}

// CreateStockReservationServices — entry point ສຳລັບ endpoint /stock-reservation/create ໂດຍກົງ (ຄຸ້ມ transaction ເອງ)
func CreateStockReservationServices(ctx context.Context, req requestbody.StockReservationRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		_, err := createStockReservationCore(ctx, db, req)
		return err
	})
}

func GetDataStockReservationServices(ctx context.Context, id *int, productVariantID *int, page, pageSize int) ([]dbschema.StockReservationDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetStockReservationDataQuery(ctx, id, productVariantID, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}

// UpdateStockReservationStatusServices — entry point ສຳລັບ endpoint /stock-reservation/:id/status ໂດຍກົງ (ຄຸ້ມ transaction ເອງ)
func UpdateStockReservationStatusServices(ctx context.Context, id int64, status string, updatedBy *int) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		reservation, err := dbquery.GetStockReservationByIDForUpdate(ctx, db, id)
		if err != nil {
			return err
		}
		return resolveStockReservationCore(ctx, db, reservation, status, updatedBy)
	})
}
