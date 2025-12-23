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

func GateAllWhitSubmenu(ctx context.Context) ([]dbschema.SubMenuSchema, error) {
	result, err := dbquery.GetSubmenuWhitAll(ctx)
	return result, err
}

func CreatedSubMeNuServiced(ctx context.Context, req requestbody.SubMenuRequesBody) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateSubMenu(ctx, db, req)
	})
	return err
}

func GetSubMenuTotalServices(ctx context.Context, id *int, page, pageSize int) ([]dbschema.SubMenuSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	subes, PaginationResult, err := dbquery.GetSubMenuByParam(ctx, id, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("faild to get data %w", err)
	}
	return subes, PaginationResult, nil
}
