package services

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"

	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
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
