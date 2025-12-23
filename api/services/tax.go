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

func CreateTaxService(ctx context.Context, req requestbody.TaxRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(context context.Context) error {
		db := dbpkg.GetDBFromContext(context)
		return dbinserts.CreateTax(context, db, req)
	})
	return err
}

func GetDAtaTaxServices(ctx context.Context) ([]dbschema.Tax, error) {
	result, err := dbquery.GetDataTax(ctx)
	return result, err
}

func GetDataByidServices(ctx context.Context, id int) (*dbschema.Tax, error) {
	result, err := dbquery.GetByIdTax(ctx, id)
	return result, err
}

func GetTaxService(ctx context.Context, id *int, page, pageSize int) ([]dbschema.Tax, *pagination.PaginationResult, error) {
	var paginationParams *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParams = pagination.NewPaginationParams(page, pageSize)
	}
	taxes, paginationResult, err := dbquery.GetTax(ctx, id, paginationParams)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get tax data: %w", err)
	}
	return taxes, paginationResult, nil
}

func TaxUpdateservices(ctx context.Context, TaxID int64, req requestbody.TaxRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(context context.Context) error {
		db := dbpkg.GetDBFromContext(context)
		return dbupdate.UpdateTax(context, db, TaxID, req)
	})
	return err
}
