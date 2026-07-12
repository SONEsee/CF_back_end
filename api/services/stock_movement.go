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

// CreateStockMovementServices ແມ່ນ "ປະຕູດຽວ" ໃນການປ່ຽນ actual_qty:
// ລັອກແຖວ inventories -> ຄິດໄລ່ qty ໃໝ່ -> ກວດບໍ່ໃຫ້ຕິດລົບ -> ອັບເດດ inventories + ບັນທຶກ ledger, ທັງໝົດໃນ transaction ດຽວ
func CreateStockMovementServices(ctx context.Context, req requestbody.StockMovementRequestBody, createdBy *int) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)

		inv, err := dbquery.GetInventoryByVariantIDForUpdate(ctx, db, req.ProductVariantID)
		if err != nil {
			return err
		}

		newActualQty := inv.ActualQty + req.QtyChange
		newAvailableQty := newActualQty - inv.ReservedQty
		if newActualQty < 0 || newAvailableQty < 0 {
			return fmt.Errorf("insufficient stock: actual_qty ຈະຕິດລົບ (ປັດຈຸບັນ %d, ປ່ຽນ %d)", inv.ActualQty, req.QtyChange)
		}

		if err := dbupdate.UpdateInventoryQty(ctx, db, req.ProductVariantID, newActualQty, newAvailableQty); err != nil {
			return err
		}

		return dbinserts.CreateStockMovement(ctx, db, req, newActualQty, createdBy)
	})
}

func GetDataStockMovementServices(ctx context.Context, id *int64, productVariantID *int, page, pageSize int) ([]dbschema.StockMovementDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetStockMovementDataQuery(ctx, id, productVariantID, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}
