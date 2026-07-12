package dbupdate

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateCustomerAddressPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.CustomerAddressPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"customer_addresses"`)

	if req.RecipientName != nil {
		query = query.Set("recipient_name", *req.RecipientName)
	}
	if req.Phone != nil {
		query = query.Set("phone", *req.Phone)
	}
	if req.Address != nil {
		query = query.Set("address", *req.Address)
	}
	if req.SubDistrict != nil {
		query = query.Set("sub_district", *req.SubDistrict)
	}
	if req.District != nil {
		query = query.Set("district", *req.District)
	}
	if req.Province != nil {
		query = query.Set("province", *req.Province)
	}
	if req.PostalCode != nil {
		query = query.Set("postal_code", *req.PostalCode)
	}
	query = query.Where("id=?", id)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("customer address with id %d not found", id)
	}
	return nil
}

// SetCustomerAddressDefault ຕັ້ງ is_default ໃຫ້ແຖວດຽວ — ໃຊ້ຮ່ວມກັບ ClearCustomerAddressDefaults ໃນ transaction ດຽວກັນ
func SetCustomerAddressDefault(ctx context.Context, tx dbpkg.DBTX, id int64, isDefault bool) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"customer_addresses"`).Set("is_default", isDefault).Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}

// ClearCustomerAddressDefaults set is_default=false ໃຫ້ທຸກແຖວຂອງ customer ນີ້ ຍົກເວັ້ນ excludeID — ໃຊ້ sync ກ່ອນຕັ້ງ default ໃໝ່
func ClearCustomerAddressDefaults(ctx context.Context, tx dbpkg.DBTX, customerID int, excludeID int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"customer_addresses"`).
		Set("is_default", false).
		Where("customer_id=?", customerID).
		Where("id<>?", excludeID)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql, args...)
	return err
}
