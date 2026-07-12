package dbinserts

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreatePayment(ctx context.Context, tx dbpkg.DBTX, req requestbody.PaymentRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"payments"`).
		Columns("order_id", "shop_bank_account_id", "payment_method", "slip_image_path", "bank_trans_ref_id").
		Values(
			req.OrderID,
			req.ShopBankAccountID,
			squirrel.Expr("?::payment_method_enum", req.PaymentMethod),
			nullableStr(req.SlipImagePath),
			nullableStr(req.BankTransRefID),
		)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
