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

func CreateRoleServices(ctx context.Context, req requestbody.RoleRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateRole(ctx, db, req)
	})
}

func GetDataRoleServices(ctx context.Context, id *int, page, pageSize int, q string) ([]dbschema.RoleDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetRoleDataQuery(ctx, id, paginationParam, q)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}

func UpdateRolePutServices(ctx context.Context, id int64, req requestbody.RoleRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateRolePut(ctx, db, id, req)
	})
}

func UpdateRoleServicesPatch(ctx context.Context, id int64, req requestbody.RolePatchRequest) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateRolePatch(ctx, db, id, req)
	})
}

func DeleteRoleServices(ctx context.Context, id int64) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbdelete.DeleteRole(ctx, db, id)
	})
}

// GetRoleOptionsServices ດຶງ role ທັງໝົດແບບບໍ່ມີ pagination — ໃຊ້ສຳລັບ dropdown/autocomplete
func GetRoleOptionsServices(ctx context.Context) ([]dbschema.RoleOptionDBSchema, error) {
	items, err := dbquery.GetRoleOptionsQuery(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get role options: %w", err)
	}
	return items, nil
}
