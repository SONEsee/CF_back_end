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

func CreateRoleServices(ctx context.Context, req requestbody.Role) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateRole(ctx, db, req)
	})
	return err
}

func GetRoleServices(ctx context.Context, id *int, page, pageSize int) ([]dbschema.Role, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	roles, PaginationResult, err := dbquery.GetRoleDBquery(ctx, id, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("faild to get data %w", err)
	}
	return roles, PaginationResult, nil
}

func UpdatedRolePut(ctx context.Context, roleID int64, req requestbody.Role) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateRolePut(ctx, db, roleID, req)
	})
	return err
}

func UpdatedRolePacth(ctx context.Context, roleID int64, req requestbody.RolePatchRequest) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateRolePacth(ctx, db, roleID, req)
	})
	return err
}

func DeletedRoleServices(ctx context.Context, RoleID int64) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbdelete.DeleteRole(ctx, db, RoleID)
	})
	return err

}
