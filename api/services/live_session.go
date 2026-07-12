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
)

func CreateLiveSessionServices(ctx context.Context, req requestbody.LiveSessionRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateLiveSession(ctx, db, req)
	})
}

func GetDataLiveSessionServices(ctx context.Context, id *int, page, pageSize int) ([]dbschema.LiveSessionDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetLiveSessionDataQuery(ctx, id, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}

func UpdateLiveSessionServicesPatch(ctx context.Context, id int64, req requestbody.LiveSessionPatchRequest) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateLiveSessionPatch(ctx, db, id, req)
	})
}

// EndLiveSessionServices — ຢຸດ live (STREAMING→ENDED), ຫ້າມຢຸດຊ້ຳ
func EndLiveSessionServices(ctx context.Context, id int64) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		sessions, _, err := dbquery.GetLiveSessionDataQuery(ctx, intPtr(int(id)), nil)
		if err != nil {
			return err
		}
		if sessions[0].Status != "STREAMING" {
			return fmt.Errorf("live session %d already ended", id)
		}
		return dbupdate.EndLiveSession(ctx, db, id)
	})
}

func intPtr(v int) *int { return &v }
