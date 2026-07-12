package dbinserts

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// CreateCommentIntent ບັນທຶກຜົນການປະມວນຜົນ comment — status ຄິດໄລ່ແລ້ວຈາກ service (CF_SUCCESS/OUT_OF_STOCK/INVALID_CODE)
func CreateCommentIntent(ctx context.Context, tx dbpkg.DBTX, commentRawID int64, customerID, matchedProductVariantID, parsedQty *int, status string) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"comment_intents"`).
		Columns("comment_raw_id", "customer_id", "matched_product_variant_id", "parsed_qty", "intent_status").
		Values(commentRawID, customerID, matchedProductVariantID, parsedQty, squirrel.Expr("?::intent_status_enum", status))
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
