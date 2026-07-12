package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateRefund(ctx context.Context, tx dbpkg.DBTX, req requestbody.RefundRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"refunds"`).
		Columns("order_id", "reason", "refund_amount").
		Values(req.OrderID, nullableStr(req.Reason), req.RefundAmount)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
