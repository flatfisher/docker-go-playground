package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestPrepareDB(t *testing.T) {
	t.Logf("setup: %s", time.Now())

	t.Run("Test", func(t *testing.T) {
		t.Run("ParamString", func(t *testing.T) {
			t.Parallel()
			p := "ENV"
			d := "test"
			if getParamString(p, d) != d {
				t.Fatal()
			}
			t.Logf("ParamString: %s", time.Now())
		})
		t.Run("DataSourceName", func(t *testing.T) {
			t.Parallel()
			if dataSourceName() != "root:password@tcp([localhost]:3306)/database" {
				t.Fatal()
			}
			t.Logf("DataSourceName: %s", time.Now())
		})
		t.Run("Connect", func(t *testing.T) {
			t.Parallel()

			db, err := connect()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			t.Logf("Connect: %s", time.Now())
		})
	})

	t.Logf("tear-down: %s", time.Now())
}

func TestQuery(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM sakila.city LIMIT 1;")
	if err != nil {
		t.Skip(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		t.Fatal(err)
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
	}
}
