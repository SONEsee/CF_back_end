package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateCustomer(ctx context.Context, tx dbpkg.DBTX, req requestbody.CustomerRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"customers"`).
		Columns("shop_id", "social_platform_id", "customer_name", "profile_pic_url", "phone_number", "tags", "note").
		Values(req.ShopID, nullableStr(req.SocialPlatformID), nullableStr(req.CustomerName), nullableStr(req.ProfilePicURL), nullableStr(req.PhoneNumber), nullableStr(req.Tags), nullableStr(req.Note))
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}

// nullableStr ປ່ຽນ empty string ໃຫ້ເປັນ nil ຕອນ insert — ຫຼີກລ່ຽງເກັບ "" ໃນຄໍລຳທີ່ nullable ແທນ NULL ແທ້
func nullableStr(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}
