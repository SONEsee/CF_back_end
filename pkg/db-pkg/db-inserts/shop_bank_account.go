package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateShopBankAccount(ctx context.Context, tx dbpkg.DBTX, req requestbody.ShopBankAccountRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"shop_bank_accounts"`).
		Columns("shop_id", "bank_name", "account_number", "account_name", "promptpay_id").
		Values(req.ShopID, req.BankName, req.AccountNumber, req.AccountName, req.PromptpayID)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
