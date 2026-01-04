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

func GetMainMenuByID(ctx context.Context, id int) (*dbschema.MainMenuDGSchema, error) {
	result, err := dbquery.GetMainMenuByID(ctx, id)
	return result, err
}

func GetAllMainMenusService(ctx context.Context) ([]dbschema.MainMenuDGSchema, error) {

	result, err := dbquery.GetAllMainMenus(ctx)
	return result, err
}

func GetMainTester(ctx context.Context) ([]dbschema.MainMenuDGSchema, error) {
	result, err := dbquery.GetTestMainmenu(ctx)
	return result, err
}

func CreateMainMenuServices(ctx context.Context, req requestbody.MainMenuRequesBody) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(context context.Context) error {
		db := dbpkg.GetDBFromContext(context)
		return dbinserts.CreateMainMenu(context, db, req)
	})
	return err
}

func CreateMainMenuServicesTest(ctx context.Context, req requestbody.MainMenuRequesBody) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(context context.Context) error {
		db := dbpkg.GetDBFromContext(context)
		return dbinserts.CreateMainMenuTest(context, db, req)
	})
	return err
}

func GetDataMainMenuServices(ctx context.Context, id *int, page, pageSize int) ([]dbschema.MainMenuDGSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	mainmenu, paginationResult, err := dbquery.GetMainMenuDataQuery(ctx, id, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getdata services %w", err)
	}
	return mainmenu, paginationResult, nil

}

func UpdateMainMenuPutServices(ctx context.Context, id int64, req requestbody.MainMenuRequesBody) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateMainMenuPut(ctx, id, db, req)
	})
	return err
}
func UpdateMainMenuServicesPacth(ctx context.Context, id int64, req requestbody.MainMenuPatchRequest) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateMainMenuPacth(ctx, db, id, req)
	})
	return err
}

func DeletedMainMenuServices(ctx context.Context, id int64, req requestbody.MainMenuPatchRequest) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbdelete.DeletedMainMenu(ctx, db, id)
	})
	return err
}
