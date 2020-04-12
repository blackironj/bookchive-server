package da

import (
	"github.com/blackironj/bookchive-server/model"
	"github.com/jmoiron/sqlx"
)

func InsertDiary(tx *sqlx.Tx, diary *model.Diary) error {
	stmt := `INSERT INTO diaries(libraries_uk, title, contents, pos_page, added_dt, updated_dt) 
			VALUES (:libraries_uk, :title, :contents, :pos_page, :added_dt, :updated_dt)`

	_, err := tx.NamedExec(stmt, diary)
	if err != nil {
		return err
	}
	return nil
}
