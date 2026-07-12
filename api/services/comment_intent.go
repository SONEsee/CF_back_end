package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	dbupdate "github.com/SONEsee/go-echo/pkg/db-pkg/db-update"
	"github.com/SONEsee/go-echo/pkg/pagination"
)

// CreateCommentIntentServices — "ຢືນຢັນ CF" ຫົວໃຈຂອງລະບົບ live-selling:
//   - ບໍ່ມີ matched_product_variant_id/parsed_qty → INVALID_CODE
//   - ມີ ແຕ່ stock ບໍ່ພຽງພໍ → OUT_OF_STOCK (ບໍ່ຈອງ)
//   - ຈອງ stock ໄດ້ → CF_SUCCESS (ຈອງຈິງຜ່ານ core function ດຽວກັບ Zone 3)
//
// ຕັ້ງ comments_raw.is_processed=true ສະເໝີ, ທັງໝົດຢູ່ໃນ transaction ດຽວ
func CreateCommentIntentServices(ctx context.Context, req requestbody.CommentIntentRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)

		comment, err := dbquery.GetCommentRawByIDForUpdate(ctx, db, req.CommentRawID)
		if err != nil {
			return err
		}
		if comment.IsProcessed {
			return fmt.Errorf("comment %d already processed", req.CommentRawID)
		}

		status := "INVALID_CODE"
		if req.MatchedProductVariantID != nil && req.ParsedQty != nil && *req.ParsedQty > 0 {
			reservationReq := requestbody.StockReservationRequestBody{
				ProductVariantID: *req.MatchedProductVariantID,
				CustomerID:       req.CustomerID,
				ReservedQty:      *req.ParsedQty,
				ExpiresAt:        time.Now().Add(reservationHoldDuration).Format(time.RFC3339),
			}
			_, err := createStockReservationCore(ctx, db, reservationReq)
			if err != nil {
				if isInsufficientStockErr(err) {
					status = "OUT_OF_STOCK"
				} else {
					return err
				}
			} else {
				status = "CF_SUCCESS"
			}
		}

		if err := dbinserts.CreateCommentIntent(ctx, db, req.CommentRawID, req.CustomerID, req.MatchedProductVariantID, req.ParsedQty, status); err != nil {
			return err
		}

		return dbupdate.MarkCommentRawProcessed(ctx, db, req.CommentRawID)
	})
}

func isInsufficientStockErr(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), "insufficient stock")
}

func GetDataCommentIntentServices(ctx context.Context, id *int64, customerID *int, page, pageSize int) ([]dbschema.CommentIntentDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetCommentIntentDataQuery(ctx, id, customerID, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}
