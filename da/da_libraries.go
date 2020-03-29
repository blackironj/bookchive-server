package da

import (
	"fmt"
	"strings"

	"github.com/blackironj/bookchive-server/model"
	"github.com/jmoiron/sqlx"
)

func InsertLibraries(tx *sqlx.Tx, libraries []model.Library) error {
	valStrings := []string{}
	valArgs := []interface{}{}

	for _, lib := range libraries {
		valStrings = append(valStrings, "(?, ?, ?)")

		valArgs = append(valArgs, lib.UserUUID)
		valArgs = append(valArgs, lib.BookID)
		valArgs = append(valArgs, lib.AddedDT)
	}

	stmt := `INSERT INTO libraries(user_uuid, book_id, added_dt) VALUES %s`
	stmt = fmt.Sprintf(stmt, strings.Join(valStrings, ","))

	_, err := tx.Exec(stmt, valArgs...)
	if err != nil {
		return err
	}

	return nil
}
