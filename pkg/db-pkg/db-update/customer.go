package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateCustomerPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.CustomerPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"customers"`)

	if req.SocialPlatformID != nil {
		query = query.Set("social_platform_id", *req.SocialPlatformID)
	}
	if req.CustomerName != nil {
		query = query.Set("customer_name", *req.CustomerName)
	}
	if req.ProfilePicURL != nil {
		query = query.Set("profile_pic_url", *req.ProfilePicURL)
	}
	if req.PhoneNumber != nil {
		query = query.Set("phone_number", *req.PhoneNumber)
	}
	if req.Tags != nil {
		query = query.Set("tags", *req.Tags)
	}
	if req.Note != nil {
		query = query.Set("note", *req.Note)
	}
	query = query.Set("updated_at", time.Now()).Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("customer with id %d not found", id)
	}
	return nil
}

// SetCustomerDefaultAddressID ຕັ້ງ/ລົບລ້າງ default_address_id — ໃຊ້ພາຍໃນ transaction ດຽວກັບການ sync is_default ຂອງ customer_addresses
func SetCustomerDefaultAddressID(ctx context.Context, tx dbpkg.DBTX, customerID int, addressID *int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"customers"`).Set("default_address_id", addressID).Set("updated_at", time.Now()).Where("id=?", customerID)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
