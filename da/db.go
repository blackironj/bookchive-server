package da

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/blackironj/bookchive-server/env"
)

const componentName = "[DA] "

//DB connection variable
var DB *sqlx.DB

//InitDB initialize a DB connection
func InitDB() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		env.Conf.DB.User,
		env.Conf.DB.Password,
		env.Conf.DB.Host,
		env.Conf.DB.Port,
		env.Conf.DB.Name)

	var err error
	DB, err = sqlx.Connect(env.Conf.DB.Type, connStr)
	if err != nil {
		log.Fatal(componentName, err)
	}

	DB.SetMaxOpenConns(env.Conf.DB.MaxOpenConn)
	DB.SetMaxIdleConns(env.Conf.DB.MaxIdleConn)
	DB.SetConnMaxLifetime(time.Duration(env.Conf.DB.MaxLifeTimeSec) * time.Second)
}

//DoInTransaction uses DB Transaction
func DoInTransaction(fn func(tx *sqlx.Tx) error) error {
	tx := DB.MustBegin()

	execErr := fn(tx)
	if execErr != nil {
		tx.Rollback()
		return execErr
	}

	if commitErr := tx.Commit(); commitErr != nil {
		tx.Rollback()
		return commitErr
	}
	return nil
}
