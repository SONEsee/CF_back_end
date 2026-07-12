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

var customerAddressColumns = []string{"id", "customer_id", "recipient_name", "phone", "address", "sub_district", "district", "province", "postal_code", "is_default"}

func scanCustomerAddress(row pgx.Row, item *dbschema.CustomerAddressDBSchema) error {
	return row.Scan(
		&item.ID, &item.CustomerID, &item.RecipientName, &item.Phone, &item.Address,
		&item.SubDistrict, &item.District, &item.Province, &item.PostalCode, &item.IsDefault,
	)
}

func GetCustomerAddressDataQuery(ctx context.Context, id *int, customerID *int, paginationParams *pagination.PaginationParams) ([]dbschema.CustomerAddressDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(customerAddressColumns...).From(`"customer_addresses"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.CustomerAddressDBSchema
		if err := scanCustomerAddress(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("customer address with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.CustomerAddressDBSchema{item}, nil, nil
	}

	countQuery := psql.Select("COUNT(*)").From(`"customer_addresses"`)
	listQuery := psql.Select(customerAddressColumns...).From(`"customer_addresses"`)
	if customerID != nil {
		countQuery = countQuery.Where("customer_id=?", *customerID)
		listQuery = listQuery.Where("customer_id=?", *customerID)
	}

	var totalItem int
	countSQL, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	listQuery = listQuery.OrderBy("id ASC")
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
		return nil, nil, fmt.Errorf("failed to query customer addresses: %w", err)
	}
	defer rows.Close()

	var items []dbschema.CustomerAddressDBSchema
	for rows.Next() {
		var item dbschema.CustomerAddressDBSchema
		if err := scanCustomerAddress(rows, &item); err != nil {
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

// GetCustomerAddressByIDForUpdate ອ່ານພ້ອມລັອກແຖວ — ໃຊ້ຕອນລົບ (ກວດ is_default) ຫຼືຕັ້ງ default ໃນ transaction ດຽວກັນ
func GetCustomerAddressByIDForUpdate(ctx context.Context, tx dbpkg.DBTX, id int64) (*dbschema.CustomerAddressDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(customerAddressColumns...).From(`"customer_addresses"`).Where("id=?", id).Suffix("FOR UPDATE")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.CustomerAddressDBSchema
	if err := scanCustomerAddress(tx.QueryRow(ctx, sql, args...), &item); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("customer address with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}
