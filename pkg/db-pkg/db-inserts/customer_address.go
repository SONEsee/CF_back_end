package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// CreateCustomerAddress ສ້າງແຖວ (is_default ໃຊ້ default ຂອງ DB ຄື false) ແລະ ຄືນ id ໃໝ່
func CreateCustomerAddress(ctx context.Context, tx dbpkg.DBTX, req requestbody.CustomerAddressRequestBody) (int64, error) {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"customer_addresses"`).
		Columns("customer_id", "recipient_name", "phone", "address", "sub_district", "district", "province", "postal_code").
		Values(req.CustomerID, nullableStr(req.RecipientName), nullableStr(req.Phone), nullableStr(req.Address), nullableStr(req.SubDistrict), nullableStr(req.District), nullableStr(req.Province), nullableStr(req.PostalCode)).
		Suffix("RETURNING id")
	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}
	var id int64
	if err := tx.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
