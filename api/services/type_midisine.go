package services

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
)

func CreateTypeMidsine(ctx context.Context, req requestbody.TypeMedicine) error {
	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateTypeMidsine(ctx, db, req)
	})
	return err
}
