package services

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbdelete "github.com/SONEsee/go-echo/pkg/db-pkg/db-delete"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	dbupdate "github.com/SONEsee/go-echo/pkg/db-pkg/db-update"
	"github.com/SONEsee/go-echo/pkg/pagination"
)

func CreateProductImageServices(ctx context.Context, req requestbody.ProductImageRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateProductImage(ctx, db, req)
	})
}

func GetDataProductImageServices(ctx context.Context, id *int, page, pageSize int) ([]dbschema.ProductImageDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetProductImageDataQuery(ctx, id, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}

func UpdateProductImageServicesPatch(ctx context.Context, id int64, req requestbody.ProductImagePatchRequest) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateProductImagePatch(ctx, db, id, req)
	})
}

func DeleteProductImageServices(ctx context.Context, id int64) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbdelete.DeleteProductImage(ctx, db, id)
	})
}
