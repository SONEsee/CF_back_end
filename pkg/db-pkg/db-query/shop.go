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

func GetShopDataQuery(ctx context.Context, id *int, paginationParams *pagination.PaginationParams) ([]dbschema.ShopDBSchema, *pagination.PaginationResult, error) {
	psql := db.GetPSQLCommand()

	if id != nil {
		query := psql.Select("id", "shop_name", "owner_user_id", "phone", "status", "timezone", "image_url", "created_at", "updated_at").
			From(`"shops"`).Where("id=?", *id)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build query: %w", err)
		}
		var item dbschema.ShopDBSchema
		err = dbpkg.DB.QueryRow(ctx, sql, args...).Scan(
			&item.ID, &item.ShopName, &item.OwnerUserID, &item.Phone, &item.Status, &item.Timezone, &item.ImageURL, &item.CreatedAt, &item.UpdatedAt,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, fmt.Errorf("shop with id %d not found", *id)
			}
			return nil, nil, fmt.Errorf("failed to execute query: %w", err)
		}
		return []dbschema.ShopDBSchema{item}, nil, nil
	}

	var totalItem int
	countSQL, countArgs, err := psql.Select("COUNT(*)").From(`"shops"`).ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}
	if err := dbpkg.DB.QueryRow(ctx, countSQL, countArgs...).Scan(&totalItem); err != nil {
		return nil, nil, fmt.Errorf("failed to count records: %w", err)
	}

	query := psql.Select("id", "shop_name", "owner_user_id", "phone", "status", "timezone", "image_url", "created_at", "updated_at").
		From(`"shops"`).OrderBy("id ASC")

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
		return nil, nil, fmt.Errorf("failed to query shops: %w", err)
	}
	defer rows.Close()

	var items []dbschema.ShopDBSchema
	for rows.Next() {
		var item dbschema.ShopDBSchema
		if err := rows.Scan(&item.ID, &item.ShopName, &item.OwnerUserID, &item.Phone, &item.Status, &item.Timezone, &item.ImageURL, &item.CreatedAt, &item.UpdatedAt); err != nil {
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

// GetShopOptionsQuery ດຶງ shop ທັງໝົດແບບບໍ່ມີ limit/pagination (ໃຊ້ສຳລັບ dropdown/autocomplete)
// ແຍກອອກຈາກ GetShopDataQuery ໂດຍສະເພາະ ເພື່ອບໍ່ໃຫ້ກະທົບ endpoint ລິດການໃຊ້ງານເກົ່າ
func GetShopOptionsQuery(ctx context.Context) ([]dbschema.ShopOptionDBSchema, error) {
	psql := db.GetPSQLCommand()

	query := psql.Select("id", "shop_name").
		From(`"shops"`).
		Where("status = ?", "ACTIVE").
		OrderBy("shop_name ASC")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query shop options: %w", err)
	}
	defer rows.Close()

	var items []dbschema.ShopOptionDBSchema
	for rows.Next() {
		var item dbschema.ShopOptionDBSchema
		if err := rows.Scan(&item.ID, &item.ShopName); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return items, nil
}
