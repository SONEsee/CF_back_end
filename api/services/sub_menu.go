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

func UpdateSubMenuPutServices(ctx context.Context, id int64, req requestbody.SubMenuRequesBody) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateSubMenuPut(ctx, db, id, req)
	})
	return err
}

func UpdateSubMenuPactServices(ctx context.Context, id int64, req requestbody.SubMenuRequesBodyPact) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateSubMenuPacth(ctx, db, id, req)
	})
	return err
}

func DeleteSubMenuServices(ctx context.Context, id int64) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbdelete.DeLeteSubMenu(ctx, db, id)
	})
	return err
}
