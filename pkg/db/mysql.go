package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var _sqlxDB *sqlx.DB

type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func InitSqlxDB(host, username, password, database string) error {
	// "test:test@tcp(127.0.0.1:3306)/abwork?charset=utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, database)
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
