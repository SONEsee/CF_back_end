package dbinserts

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// CreateStockMovement ບັນທຶກ ledger row — ຮັບ balanceAfter/createdBy ທີ່ຄິດໄລ່ແລ້ວຈາກ service (ບໍ່ຮັບ balance_after ຈາກ client)
func CreateStockMovement(ctx context.Context, tx dbpkg.DBTX, req requestbody.StockMovementRequestBody, balanceAfter int, createdBy *int) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"stock_movements"`).
		Columns("product_variant_id", "movement_type", "qty_change", "balance_after", "ref_type", "ref_id", "note", "created_by").
		Values(
			req.ProductVariantID,
			squirrel.Expr("?::movement_type_enum", req.MovementType),
			req.QtyChange,
			balanceAfter,
			req.RefType,
			req.RefID,
			req.Note,
			createdBy,
		)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
