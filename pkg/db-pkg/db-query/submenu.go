package dbquery

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/config/db"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
)

func GetSubmenuWhitAll(ctx context.Context) ([]dbschema.SubMenuSchema, error) {
	var result []dbschema.SubMenuSchema
	psql := db.GetPSQLCommand()
	query := psql.Select("id", "name_submenu", "icon_submenu", "url_submenu", "action", "main_menu_id").From(`"SubMenu"`).OrderBy("id ASC")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("fail convert forn sql %w", err)
	}
	rows, err := dbpkg.DB.Query(ctx, sql, args...)

	if err != nil {
		return nil, fmt.Errorf("fail excue %w ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item dbschema.SubMenuSchema
		err = rows.Scan(
			&item.ID,
			&item.NameSubMenu,
			&item.IconSubMenu,
			&item.URLSubMenu,
			&item.Action,
			&item.MainMenuID,
		)
		if err != nil {
			return nil, fmt.Errorf("fail to scan data %w", err)
		}
		result = append(result, item)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("fail to data %w", err)
	}
	return result, nil

}
