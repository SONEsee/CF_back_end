package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateShop(ctx context.Context, tx dbpkg.DBTX, req requestbody.ShopRequestBody) error {
	psql := db.GetPSQLCommand()
	timezone := req.Timezone
	if timezone == "" {
		timezone = "Asia/Bangkok"
	}
	var imageURL *string
	if req.ImageURL != "" {
		imageURL = &req.ImageURL
	}
	query := psql.Insert(`"shops"`).
		Columns("shop_name", "owner_user_id", "phone", "timezone", "image_url").
		Values(req.ShopName, req.OwnerUserID, req.Phone, timezone, imageURL)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
