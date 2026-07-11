package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func CreateProductImage(ctx context.Context, tx dbpkg.DBTX, req requestbody.ProductImageRequestBody) error {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"product_images"`).
		Columns("product_id", "image_url", "sort_order").
		Values(req.ProductID, req.ImageURL, req.SortOrder)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
