package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateShipmentPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ShipmentPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"shipments"`)

	if req.CourierName != nil {
		query = query.Set("courier_name", *req.CourierName)
	}
	if req.TrackingNumber != nil {
		query = query.Set("tracking_number", *req.TrackingNumber)
	}
	if req.LabelPdfURL != nil {
		query = query.Set("label_pdf_url", *req.LabelPdfURL)
	}
	query = query.Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("shipment with id %d not found", id)
	}
	return nil
}

// UpdateShipmentStatus ປ່ຽນ shipping_status — PICKED_UP set shipped_at, DELIVERED set delivered_at ອັດຕະໂນມັດ
func UpdateShipmentStatus(ctx context.Context, tx dbpkg.DBTX, id int64, status string) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"shipments"`).Set("shipping_status", squirrel.Expr("?::shipping_status_enum", status))
	if status == "PICKED_UP" {
		query = query.Set("shipped_at", time.Now())
	} else if status == "DELIVERED" {
		query = query.Set("delivered_at", time.Now())
	}
	query = query.Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("shipment with id %d not found", id)
	}
	return nil
}
