package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// VerifyPayment ຢືນຢັນ/ປະຕິເສດສະລິບ — set is_valid_slip, verified_amount, paid_at (ຖ້າ valid)
func VerifyPayment(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.PaymentVerifyRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"payments"`).
		Set("is_valid_slip", req.IsValidSlip).
		Set("verified_amount", req.VerifiedAmount)
	if req.IsValidSlip {
		query = query.Set("paid_at", time.Now())
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
		return fmt.Errorf("payment with id %d not found", id)
	}
	return nil
}
