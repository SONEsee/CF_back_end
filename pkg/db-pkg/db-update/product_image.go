package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateProductImagePatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ProductImagePatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"product_images"`)

	if req.ImageURL != nil {
		query = query.Set("image_url", *req.ImageURL)
	}
	if req.SortOrder != nil {
		query = query.Set("sort_order", *req.SortOrder)
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
		return fmt.Errorf("product image with id %d not found", id)
	}
	return nil
}
