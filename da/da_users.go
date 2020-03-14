package da

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

//Users Model
type Users struct {
	UUID     string         `db:"uuid"`
	Email    string         `db:"email"`
	Name     sql.NullString `db:"name"`
	SigninDT sql.NullInt64  `db:"signin_dt"`
}

//GetUsers get user informations from DB
func GetUsers(db *sqlx.DB, where string, condVal []interface{}) ([]Users, error) {
	users := []Users{}

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
func InsertUser(tx *sqlx.Tx, user *Users) error {
	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	stmt := "INSERT INTO users (uuid, email, name, signin_dt) VALUES (?, ?, ?, ?)"
	res := tx.MustExec(stmt, u.String(), user.Email, user.Name, time.Now().Unix())

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
	res := tx.MustExec(stmt, val...)

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n < 1 {
		return errors.New("not affected")
	}
	return nil
}
