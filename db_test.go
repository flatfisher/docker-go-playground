package main

import (
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
			if dataSourceName() != "user:password@tcp([localhost]:3306)/database" {
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
