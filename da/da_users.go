package da

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/blackironj/bookchive-server/model"
)

//GetUsers get user informations from DB
func GetUsers(db *sqlx.DB, where string, condVal []interface{}) ([]model.User, error) {
	users := []model.User{}

	stmt := "SELECT * FROM users " + where
	err := db.Select(&users, stmt, condVal...)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("Cannot find users")
	}

	return users, nil
}

//InsertUser add new user to DB
func InsertUser(tx *sqlx.Tx, user *model.User) error {
	stmt := "INSERT INTO users (uuid, email, name, signin_dt) VALUES (?, ?, ?, ?)"
	res, err := tx.Exec(stmt, user.UUID, user.Email, user.Name, time.Now().Unix())
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n < 1 {
		return errors.New("not affected")
	}
	return nil
}

//UpdateUser update user information
func UpdateUser(tx *sqlx.Tx, setStmt string, val []interface{}) error {
	stmt := "UPDATE users SET " + setStmt
	res, err := tx.Exec(stmt, val...)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n < 1 {
		return errors.New("not affected")
	}
	return nil
}
