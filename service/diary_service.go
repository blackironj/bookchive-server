package service

import (
	"time"

	"github.com/blackironj/bookchive-server/da"
	"github.com/blackironj/bookchive-server/model"
	"github.com/jmoiron/sqlx"
)

func GetDiaries(userUUID string, libUK int) ([]model.Diary, error) {
	diareis, getErr := da.GetDiaries(da.DB, userUUID, libUK)
	if getErr != nil {
		return nil, getErr
	}

	return diareis, nil
}

func AddDiary(diary *model.Diary) error {
	currTime := time.Now().Unix()
	diary.AddedDT = &currTime
	diary.UpdatedDT = &currTime

	txErr := da.DoInTransaction(func(tx *sqlx.Tx) error {
		if err := da.InsertDiary(tx, diary); err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		return txErr
	}

	return nil
}
