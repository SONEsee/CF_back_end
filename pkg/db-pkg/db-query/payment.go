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

var paymentColumns = []string{"id", "order_id", "shop_bank_account_id", "payment_method", "slip_image_path", "bank_trans_ref_id", "verified_amount", "is_valid_slip", "paid_at", "created_at"}

func scanPayment(row pgx.Row, item *dbschema.PaymentDBSchema) error {
	return row.Scan(
		&item.ID, &item.OrderID, &item.ShopBankAccountID, &item.PaymentMethod, &item.SlipImagePath,
		&item.BankTransRefID, &item.VerifiedAmount, &item.IsValidSlip, &item.PaidAt, &item.CreatedAt,
	)
}

func GetPaymentDataQuery(ctx context.Context, id *int, orderID *int, paginationParams *pagination.PaginationParams) ([]dbschema.PaymentDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(paymentColumns...).From(`"payments"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.PaymentDBSchema
		if err := scanPayment(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("payment with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.PaymentDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"payments"`)
	listQuery := psql.Select(paymentColumns...).From(`"payments"`)
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
		return nil, nil, fmt.Errorf("failed to query payments: %w", err)
	}
	defer rows.Close()

	var items []dbschema.PaymentDBSchema
	for rows.Next() {
		var item dbschema.PaymentDBSchema
		if err := scanPayment(rows, &item); err != nil {
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
