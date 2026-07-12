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

var customerColumns = []string{"id", "shop_id", "social_platform_id", "customer_name", "profile_pic_url", "phone_number", "default_address_id", "tags", "note", "created_at", "updated_at"}

func scanCustomer(row pgx.Row, item *dbschema.CustomerDBSchema) error {
	return row.Scan(
		&item.ID, &item.ShopID, &item.SocialPlatformID, &item.CustomerName, &item.ProfilePicURL,
		&item.PhoneNumber, &item.DefaultAddressID, &item.Tags, &item.Note, &item.CreatedAt, &item.UpdatedAt,
	)
}

func GetCustomerDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.CustomerDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select(customerColumns...).From(`"customers"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.CustomerDBSchema
		if err := scanCustomer(dbpkg.DB.QueryRow(ctx, sql, args...), &item); err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("customer with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.CustomerDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"customers"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select(customerColumns...).From(`"customers"`).OrderBy("id ASC")

	var paginationResult *pagination.PaginationResult
	if paginationParams != nil && paginationParams.IsValid() {
		query = query.Limit(uint64(paginationParams.GetLimit())).Offset(uint64(paginationParams.GetOffset()))
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build query: %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query customers: %w", err)
	}
	defer rows.Close()

	var items []dbschema.CustomerDBSchema
	for rows.Next() {
		var item dbschema.CustomerDBSchema
		if err := scanCustomer(rows, &item); err != nil {
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

// GetCustomerBySocialPlatformID ຄົ້ນຫາ customer ດ້ວຍ (shop_id, social_platform_id) — ໃຊ້ຕອນ parse comment ອັດຕະໂນມັດ
// ເພື່ອຮູ້ວ່າ fb_user_id ນີ້ແມ່ນ customer ຄົນໃດ. ຄືນ nil, nil ຖ້າບໍ່ພົບ (ບໍ່ແມ່ນ error — customer ອາດຍັງບໍ່ເຄີຍລົງທະບຽນ)
func GetCustomerBySocialPlatformID(ctx context.Context, shopID int, socialPlatformID string) (*dbschema.CustomerDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(customerColumns...).From(`"customers"`).
		Where("shop_id=?", shopID).
		Where("social_platform_id=?", socialPlatformID)
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.CustomerDBSchema
	err = scanCustomer(dbpkg.DB.QueryRow(ctx, sql, args...), &item)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}

// GetCustomerByIDForUpdate ອ່ານພ້ອມລັອກແຖວ — ໃຊ້ຕອນສ້າງ/ຕັ້ງ default address ໃນ transaction ດຽວກັນ
func GetCustomerByIDForUpdate(ctx context.Context, tx dbpkg.DBTX, id int) (*dbschema.CustomerDBSchema, error) {
	psql := db.GetPSQLCommand()
	query := psql.Select(customerColumns...).From(`"customers"`).Where("id=?", id).Suffix("FOR UPDATE")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	var item dbschema.CustomerDBSchema
	if err := scanCustomer(tx.QueryRow(ctx, sql, args...), &item); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("customer with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &item, nil
}
