package database

import (
	"database/sql"
	"fmt"
)

type DBType string

const (
	MySQL = "mysql"
	Postgres = "postgres"
)

func NewDbConnection(dbType DBType,  connString string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch dbType {
	case MySQL:
		db, err = sql.Open("mysql", connString)
	case Postgres:
		db, err = sql.Open("postgres", connString)
	default: 
		return nil, fmt.Errorf("unsupported database type: %v", dbType)
	}

	if err != nil {
		return nil, fmt.Errorf("error openning database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}
	
	return db, err
}
