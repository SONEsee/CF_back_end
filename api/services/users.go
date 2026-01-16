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
	"golang.org/x/crypto/bcrypt"
)

// func CreateUserService(ctx context.Context, req requestbody.UserRequestBody) error {
// 	tx := dbpkg.GetTransactionManager()
// 	err := tx.WithTransaction(ctx, func(context context.Context) error {
// 		db := dbpkg.GetDBFromContext(context)
// 		return dbinserts.InsertNewUserTx(context, db, req)

// 	})

// 	return err
// }
// ໃນໄຟລ໌ service/user_service.go

func CreateUserService(ctx context.Context, req requestbody.UserRequestBody) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	req.Password = string(hashedPassword)

	tx := dbpkg.GetTransactionManager()
	err = tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.InsertNewUserTx(ctx, db, req)
	})

	return err
}

func GetUserService(ctx context.Context, id *int, page, pageSize int) ([]dbschema.GetUserDataDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	result, PaginationParam, err := dbquery.GetUserDataDBQuery(ctx, id, paginationParam)
	if err != nil {
		return nil, nil, err
	}
	return result, PaginationParam, nil
}

func UpdateUserServices(ctx context.Context, id int64, req requestbody.UserRequestBodyPacth) error {

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		req.Password = string(hashedPassword)
	}

	tx := dbpkg.GetTransactionManager()
	err := tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateUser(ctx, db, id, req)
	})

	return err
}
