package dbinserts

import (
	"context"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

// CreateChatConversation ສ້າງ conversation ໃໝ່ (status ໃຊ້ default OPEN ຂອງ DB) ແລະ ຄືນ id — ໃຊ້ພາຍໃນ find-or-create flow
func CreateChatConversation(ctx context.Context, tx dbpkg.DBTX, socialAccountID, customerID int) (int64, error) {
	psql := db.GetPSQLCommand()
	query := psql.Insert(`"chat_conversations"`).
		Columns("social_account_id", "customer_id").
		Values(socialAccountID, customerID).
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
