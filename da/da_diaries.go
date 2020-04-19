package da

import (
	"errors"

	"github.com/blackironj/bookchive-server/model"
	"github.com/jmoiron/sqlx"
)

func GetDiaries(db *sqlx.DB, userUUID string, libUK int) ([]model.Diary, error) {
	diaries := []model.Diary{}

	stmt := `SELECT 
			diaries.uk,
			diaries.libraries_uk,
			diaries.title,
			diaries.contents, 
			diaries.pos_page, 
			diaries.added_dt, 
			diaries.updated_dt 
			FROM diaries
				LEFT JOIN libraries ON libraries.uk = diaries.libraries_uk
			WHERE libraries.user_uuid = ? AND diaries.libraries_uk = ?`

	err := db.Select(&diaries, stmt, userUUID, libUK)
	if err != nil {
		return nil, err
	}

	if len(diaries) == 0 {
		return nil, errors.New("no datas")
	}

	return diaries, nil
}

func InsertDiary(tx *sqlx.Tx, diary *model.Diary) error {
	stmt := `INSERT INTO diaries(libraries_uk, title, contents, pos_page, added_dt, updated_dt) 
			VALUES (:libraries_uk, :title, :contents, :pos_page, :added_dt, :updated_dt)`

	_, err := tx.NamedExec(stmt, diary)
	if err != nil {
		return err
	}
	return nil
}
