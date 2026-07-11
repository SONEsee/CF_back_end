package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateShopBankAccountPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ShopBankAccountPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"shop_bank_accounts"`)

	if req.BankName != nil {
		query = query.Set("bank_name", *req.BankName)
	}
	if req.AccountNumber != nil {
		query = query.Set("account_number", *req.AccountNumber)
	}
	if req.AccountName != nil {
		query = query.Set("account_name", *req.AccountName)
	}
	if req.PromptpayID != nil {
		query = query.Set("promptpay_id", *req.PromptpayID)
	}
	if req.IsActive != nil {
		query = query.Set("is_active", *req.IsActive)
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
		return fmt.Errorf("shop bank account with id %d not found", id)
	}
	return nil
}

// DeactivateShopBankAccount ໃຊ້ແທນການລົບ (shop_bank_accounts ບໍ່ມີ deleted_at) — set is_active = false
func DeactivateShopBankAccount(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"shop_bank_accounts"`).Set("is_active", false).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("shop bank account with id %d not found", id)
	}
	return nil
}
