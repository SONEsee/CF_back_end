package services

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	dbupdate "github.com/SONEsee/go-echo/pkg/db-pkg/db-update"
	"github.com/SONEsee/go-echo/pkg/pagination"
)

// CreateProductVariantServices ສ້າງ variant ພ້ອມ inventories ເລີ່ມຕົ້ນ (qty=0) ໃນ transaction ດຽວກັນ
func CreateProductVariantServices(ctx context.Context, req requestbody.ProductVariantRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		variantID, err := dbinserts.CreateProductVariant(ctx, db, req)
		if err != nil {
			return err
		}
		return dbinserts.CreateInventoryForVariant(ctx, db, variantID)
	})
}

func GetDataProductVariantServices(ctx context.Context, id *int, page, pageSize int) ([]dbschema.ProductVariantDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetProductVariantDataQuery(ctx, id, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}

func UpdateProductVariantServicesPatch(ctx context.Context, id int64, req requestbody.ProductVariantPatchRequest) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateProductVariantPatch(ctx, db, id, req)
	})
}

func DeactivateProductVariantServices(ctx context.Context, id int64) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.DeactivateProductVariant(ctx, db, id)
	})
}
