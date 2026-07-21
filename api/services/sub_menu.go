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

func CreateSubMenuServices(ctx context.Context, req requestbody.SubMenuRequesBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateSubMenu(ctx, db, req)
	})
}

func GetDataSubMenuServices(ctx context.Context, id *int, q *string, page, pageSize int) ([]dbschema.SubMenuDBSchema, *pagination.PaginationResult, error) {
    var paginationParam *pagination.PaginationParams
    if page > 0 || pageSize > 0 {
        paginationParam = pagination.NewPaginationParams(page, pageSize)
    }
    
    items, result, err := dbquery.GetSubMenuDataQuery(ctx, id, q, paginationParam)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to get data: %w", err)
    }
    return items, result, nil
}

func UpdateSubMenuPutServices(ctx context.Context, id int64, req requestbody.SubMenuRequesBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateSubMenuPut(ctx, db, id, req)
	})
}

func UpdateSubMenuServicesPatch(ctx context.Context, id int64, req requestbody.SubMenuPatchRequest) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateSubMenuPatch(ctx, db, id, req)
	})
}

func DeleteSubMenuServices(ctx context.Context, id int64) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbdelete.DeleteSubMenu(ctx, db, id)
	})
}
