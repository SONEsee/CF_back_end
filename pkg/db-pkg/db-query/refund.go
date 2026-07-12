package dbquery

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	"github.com/SONEsee/go-echo/pkg/pagination"
	"github.com/jackc/pgx/v5"
)

var refundColumns = []string{"id", "order_id", "reason", "refund_amount", "status", "created_at"}

func scanRefund(row pgx.Row, item *dbschema.RefundDBSchema) error {
	return row.Scan(&item.ID, &item.OrderID, &item.Reason, &item.RefundAmount, &item.Status, &item.CreatedAt)
}

func GetRefundDataQuery(ctx context.Context, id *int, orderID *int, paginationParams *pagination.PaginationParams) ([]dbschema.RefundDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(refundColumns...).From(`"refunds"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.RefundDBSchema
		if err := scanRefund(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("refund with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.RefundDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"refunds"`)
	listQuery := psql.Select(refundColumns...).From(`"refunds"`)
	if orderID != nil {
		countQuery = countQuery.Where("order_id=?", *orderID)
		listQuery = listQuery.Where("order_id=?", *orderID)
	}

	var totalItem int
	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	listQuery = listQuery.OrderBy("id DESC")
	var paginationResult *pagination.PaginationResult
	if paginationParams != nil && paginationParams.IsValid() {
		listQuery = listQuery.Limit(uint64(paginationParams.GetLimit())).Offset(uint64(paginationParams.GetOffset()))
	}

	sql, args, err := listQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build query: %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query refunds: %w", err)
	}
	defer rows.Close()

	var items []dbschema.RefundDBSchema
	for rows.Next() {
		var item dbschema.RefundDBSchema
		if err := scanRefund(rows, &item); err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if paginationParams != nil && paginationParams.IsValid() {
		paginationResult = paginationParams.CalculatePagination(totalItem, len(items))
	}
	return items, paginationResult, nil
}

// GetRefundByIDForUpdate ອ່ານພ້ອມລັອກແຖວ — ໃຊ້ຕອນປ່ຽນ status ເພື່ອກວດ transition ຖືກຕ້ອງ
func GetRefundByIDForUpdate(ctx context.Context, tx dbpkg.DBTX, id int64) (*dbschema.RefundDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(refundColumns...).From(`"refunds"`).Where("id=?", id).Suffix("FOR UPDATE")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.RefundDBSchema
	if err := scanRefund(tx.QueryRow(ctx, sql, args...), &item); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("refund with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}
