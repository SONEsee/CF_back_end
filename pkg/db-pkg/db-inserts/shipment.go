package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateShipment(ctx context.Context, tx dbpkg.DBTX, req requestbody.ShipmentRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"shipments"`).
		Columns("order_id", "courier_name", "tracking_number", "label_pdf_url").
		Values(req.OrderID, nullableStr(req.CourierName), nullableStr(req.TrackingNumber), nullableStr(req.LabelPdfURL))
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
