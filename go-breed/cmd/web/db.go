package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	maxOpenDbConn = 25
	maxIdleDbConn = 25
	maxDbLifetime = 5 * time.Minute
)

func initMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(maxDbLifetime)
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)

	return db, nil
}
