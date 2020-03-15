package service

import (
	"time"

	"github.com/blackironj/bookchive-server/da"
	"github.com/blackironj/bookchive-server/model"
	"github.com/jmoiron/sqlx"
)

func AddLib(uuid string, books []model.Books) error {
	txErr := da.DoInTransaction(func(tx *sqlx.Tx) error {
		if err := da.InsertBooks(tx, books); err != nil {
			return err
		}

		currTime := time.Now().Unix()
		libs := []model.Libraries{}
		for _, book := range books {
			libs = append(libs, model.Libraries{
				UserUUID: uuid,
				BookID:   book.ID,
				AddedDT:  &currTime,
			})
		}

		if err := da.InsertLibraries(tx, libs); err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		return txErr
	}

	return nil
}
