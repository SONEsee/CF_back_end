package services

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	dbupdate "github.com/SONEsee/go-echo/pkg/db-pkg/db-update"
)

func CreateShopSettingServices(ctx context.Context, req requestbody.ShopSettingRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateShopSetting(ctx, db, req)
	})
}

func GetShopSettingServices(ctx context.Context, shopID int) (*dbschema.ShopSettingDBSchema, error) {
	return dbquery.GetShopSettingByShopID(ctx, shopID)
}

func UpdateShopSettingServicesPatch(ctx context.Context, shopID int64, req requestbody.ShopSettingPatchRequest) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateShopSettingPatch(ctx, db, shopID, req)
	})
}
