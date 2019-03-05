package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	t.Run("Request", func(t *testing.T) {
		t.Run("GET", func(t *testing.T) {
			t.Parallel()
			m := "GET"
			res := request(m, t)
			if res.Body.String() != fmt.Sprintf("You requested %s request!", m) {
				t.Fatal(res.Body.String())
			}
		})
		t.Run("POST", func(t *testing.T) {
			t.Parallel()
			m := "POST"
			res := request(m, t)
			if res.Body.String() != fmt.Sprintf("You requested %s request!", m) {
				t.Fatal(res.Body.String())
			}
		})
		t.Run("PUT", func(t *testing.T) {
			t.Parallel()
			m := "PUT"
			res := request(m, t)
			if res.Body.String() != fmt.Sprintf("You requested %s request!", m) {
				t.Fatal(res.Body.String())
			}
		})
		t.Run("DELETE", func(t *testing.T) {
			t.Parallel()
			m := "DELETE"
			res := request(m, t)
			if res.Body.String() != fmt.Sprintf("You requested %s request!", m) {
				t.Fatal(res.Body.String())
			}
		})
	})
}

func request(m string, t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest(m, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
	return rr
}
