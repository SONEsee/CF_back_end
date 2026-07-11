package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateProductCategory(ctx context.Context, tx dbpkg.DBTX, req requestbody.ProductCategoryRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"product_categories"`).
		Columns("shop_id", "parent_id", "name", "sort_order").
		Values(req.ShopID, req.ParentID, req.Name, req.SortOrder)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
