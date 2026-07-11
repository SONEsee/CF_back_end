package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateProductCategoryPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.ProductCategoryPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"product_categories"`)

	if req.ParentID != nil {
		query = query.Set("parent_id", *req.ParentID)
	}
	if req.Name != nil {
		query = query.Set("name", *req.Name)
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
		return fmt.Errorf("product category with id %d not found", id)
	}
	return nil
}
