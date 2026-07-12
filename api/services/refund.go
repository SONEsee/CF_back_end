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

var refundTransitions = map[string][]string{
	"REQUESTED": {"APPROVED", "REJECTED"},
	"APPROVED":  {"DONE"},
	"DONE":      {},
	"REJECTED":  {},
}

func CreateRefundServices(ctx context.Context, req requestbody.RefundRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateRefund(ctx, db, req)
	})
}

func GetDataRefundServices(ctx context.Context, id *int, orderID *int, page, pageSize int) ([]dbschema.RefundDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetRefundDataQuery(ctx, id, orderID, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}

// UpdateRefundStatusServices — REQUESTED→APPROVED→DONE ຫຼື REQUESTED→REJECTED, ຫ້າມຂ້າມຂັ້ນ
func UpdateRefundStatusServices(ctx context.Context, id int64, status string) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		refund, err := dbquery.GetRefundByIDForUpdate(ctx, db, id)
		if err != nil {
			return err
		}
		allowed := false
		for _, s := range refundTransitions[refund.Status] {
			if s == status {
				allowed = true
				break
			}
		}
		if !allowed {
			return fmt.Errorf("invalid refund status transition: %s -> %s", refund.Status, status)
		}
		return dbupdate.UpdateRefundStatus(ctx, db, id, status)
	})
}
