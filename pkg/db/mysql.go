package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var _sqlxDB *sqlx.DB

type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func InitSqlxDB(dsn string) error {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(32)
	db.SetMaxIdleConns(2)
	// https://github.com/go-sql-driver/mysql/issues/446
	db.SetConnMaxLifetime(time.Second * 14400)
	_sqlxDB = db
	return nil
}

func GetSqlxDB() *sqlx.DB {
	return _sqlxDB
}
