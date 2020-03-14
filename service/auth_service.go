package service

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/blackironj/bookchive-server/da"
)

//Signin with google
func Signin(signinData *da.Users) error {
	where := "WHERE email = ?"
	condVal := []interface{}{signinData.Email}

	user, _ := da.GetUsers(da.DB, where, condVal)
	if user != nil {
		txErr := da.DoInTransaction(func(tx *sqlx.Tx) error {
			setStmt := "SET signin_dt = ? WHERE = ?"
			val := []interface{}{time.Now().Unix(), user[0].UUID}

			if err := da.UpdateUser(tx, setStmt, val); err != nil {
				return err
			}
			return nil
		})

		if txErr != nil {
			return txErr
		}
	}

	txErr := da.DoInTransaction(func(tx *sqlx.Tx) error {
		newUser := &da.Users{
			Email: signinData.Email,
			Name:  signinData.Name,
		}
		if err := da.InsertUser(tx, newUser); err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		return txErr
	}

	return nil
}
