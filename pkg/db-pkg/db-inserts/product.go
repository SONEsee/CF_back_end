package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateProduct(ctx context.Context, tx dbpkg.DBTX, req requestbody.ProductRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"products"`).
		Columns("shop_id", "category_id", "product_name", "description", "image_main_url").
		Values(req.ShopID, req.CategoryID, req.ProductName, req.Description, req.ImageMainURL)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
