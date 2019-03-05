package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp([localhost]:3306)/database ")
	if err != nil {
		return nil, err
	}
	return db, err
}
