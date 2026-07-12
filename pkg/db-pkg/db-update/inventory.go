package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// UpdateInventoryPatch ອະນຸຍາດແກ້ໄດ້ສະເພາະ reorder_level — qty ທັງໝົດຕ້ອງປ່ຽນຜ່ານ stock_movements ເທົ່ານັ້ນ ເພື່ອຮັກສາ audit trail
func UpdateInventoryPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.InventoryPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"inventories"`)

	if req.ReorderLevel != nil {
		query = query.Set("reorder_level", *req.ReorderLevel)
	}
	query = query.Set("last_updated", time.Now()).Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("inventory with id %d not found", id)
	}
	return nil
}

// UpdateInventoryQty ອັບເດດ actual_qty/available_qty — ເອີ້ນຈາກ CreateStockMovementServices ເທົ່ານັ້ນ, ຢູ່ໃນ transaction ດຽວກັບການລັອກແຖວ (GetInventoryByVariantIDForUpdate)
func UpdateInventoryQty(ctx context.Context, tx dbpkg.DBTX, productVariantID int, newActualQty, newAvailableQty int) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"inventories"`).
		Set("actual_qty", newActualQty).
		Set("available_qty", newAvailableQty).
		Set("last_updated", time.Now()).
		Where("product_variant_id=?", productVariantID)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("inventory for product_variant_id %d not found", productVariantID)
	}
	return nil
}

// UpdateInventoryReservedQty ອັບເດດ reserved_qty ຢ່າງດຽວ — ເອີ້ນຄຽງຄູ່ກັບ UpdateInventoryQty ຈາກ stock_reservations flow (ຢູ່ໃນ transaction ດຽວກັນ)
func UpdateInventoryReservedQty(ctx context.Context, tx dbpkg.DBTX, productVariantID int, newReservedQty int) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"inventories"`).
		Set("reserved_qty", newReservedQty).
		Set("last_updated", time.Now()).
		Where("product_variant_id=?", productVariantID)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("inventory for product_variant_id %d not found", productVariantID)
	}
	return nil
}
