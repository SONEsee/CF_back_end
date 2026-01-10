package services

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	"github.com/SONEsee/go-echo/pkg/pagination"
)

func CreateTypeMidsine(ctx context.Context, req requestbody.TypeMedicine) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateTypeMidsine(ctx, db, req)
	})
	return err
}
func GetDataTypeMedicineServices(ctx context.Context, id *int, page, pageSize int) ([]dbschema.TypeMedicineDBSchema, *pagination.PaginationResult, error) {
	var paginationParams *pagination.PaginationParams

	if page > 0 || pageSize > 0 {
		params := pagination.NewPaginationParams(page, pageSize)
		paginationParams = params
	}

	result, paginationResult, err := dbquery.GetDataTypeMidsine(ctx, id, paginationParams)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get type medicine data: %w", err)
	}

	return result, paginationResult, nil
}
