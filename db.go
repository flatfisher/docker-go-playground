package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName())
	if err != nil {
		return nil, err
	}
	return db, err
}

func dataSourceName() string {
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "password")
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	host := getParamString("MYSQL_DB_HOST", "localhost")
	port := getParamString("MYSQL_PORT", "3306")
	db := getParamString("MYSQL_DB", "database")
	args := getParamString("MYSQL_DBARGS", " ")

	if strings.Trim(args, " ") != "" {
		args = "?" + args
	} else {
		args = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", user, pass, protocol, host, port, db, args)
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}
