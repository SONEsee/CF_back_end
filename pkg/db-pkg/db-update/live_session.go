package dbupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
)

func UpdateLiveSessionPatch(ctx context.Context, tx dbpkg.DBTX, id int64, req requestbody.LiveSessionPatchRequest) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"live_sessions"`)

	if req.SessionTitle != nil {
		query = query.Set("session_title", *req.SessionTitle)
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
		return fmt.Errorf("live session with id %d not found", id)
	}
	return nil
}

// EndLiveSession ປ່ຽນ status -> ENDED ພ້ອມ set ended_at — caller (service) ຕ້ອງກວດວ່າຍັງ STREAMING ຢູ່ກ່ອນເອີ້ນ
func EndLiveSession(ctx context.Context, tx dbpkg.DBTX, id int64) error {
	psql := db.GetPSQLCommand()
	query := psql.Update(`"live_sessions"`).
		Set("status", squirrel.Expr("?::live_status_enum", "ENDED")).
		Set("ended_at", time.Now()).
		Where("id=?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("live session with id %d not found", id)
	}
	return nil
}
