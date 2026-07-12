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

func CreateWebhookEventServices(ctx context.Context, req requestbody.WebhookEventRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateWebhookEvent(ctx, db, req)
	})
}

func GetDataWebhookEventServices(ctx context.Context, id *int, processed *bool, page, pageSize int) ([]dbschema.WebhookEventDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetWebhookEventDataQuery(ctx, id, processed, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}

// IngestWebhookEventServices ບັນທຶກ event ຈາກ webhook ຈິງ (ຫຼັງກວດ signature ຜ່ານແລ້ວ) — ຄົ້ນຫາ social_account
// ດ້ວຍ (platform, platformAccountID), ຖ້າບໍ່ພົບ (ຍັງບໍ່ໄດ້ເຊື່ອມຕໍ່ page/channel ນີ້ໄວ້ໃນລະບົບ) ຈະ return error ໃຫ້ caller ຕັດສິນໃຈ
func IngestWebhookEventServices(ctx context.Context, platform, platformAccountID, eventType, rawPayload string) error {
	account, err := dbquery.GetSocialAccountByPlatformAccount(ctx, platform, platformAccountID)
	if err != nil {
		return err
	}

	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbinserts.CreateWebhookEvent(ctx, db, requestbody.WebhookEventRequestBody{
			SocialAccountID: account.ID,
			EventType:       eventType,
			RawPayload:      rawPayload,
		})
	})
}

func MarkWebhookEventProcessedServices(ctx context.Context, id int64) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.MarkWebhookEventProcessed(ctx, db, id)
	})
}
