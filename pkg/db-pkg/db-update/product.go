package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateProductPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ProductPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"products"`)

	if req.CategoryID != nil {
		query = query.Set("category_id", *req.CategoryID)
	}
	if req.ProductName != nil {
		query = query.Set("product_name", *req.ProductName)
	}
	if req.Description != nil {
		query = query.Set("description", *req.Description)
	}
	if req.ImageMainURL != nil {
		query = query.Set("image_main_url", *req.ImageMainURL)
	}
	if req.IsActive != nil {
		query = query.Set("is_active", *req.IsActive)
	}
	query = query.Set("updated_at", time.Now()).Where("id=?", id).Where("deleted_at IS NULL")

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("product with id %d not found", id)
	}
	return nil
}
