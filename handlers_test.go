package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	t.Run("Request", func(t *testing.T) {
		t.Run("GET", func(t *testing.T) {
			t.Parallel()
			m := "GET"
			res := request(m, nil, helloHandler, t)
			if res.Body.String() != fmt.Sprintf("You requested %s request!", m) {
				t.Fatal(res.Body.String())
			}
		})
		t.Run("POST", func(t *testing.T) {
			t.Parallel()
			m := "POST"
			res := request(m, nil, helloHandler, t)
			if res.Body.String() != fmt.Sprintf("You requested %s request!", m) {
				t.Fatal(res.Body.String())
			}
		})
		t.Run("PUT", func(t *testing.T) {
			t.Parallel()
			m := "PUT"
			res := request(m, nil, helloHandler, t)
			if res.Body.String() != fmt.Sprintf("You requested %s request!", m) {
				t.Fatal(res.Body.String())
			}
		})
		t.Run("DELETE", func(t *testing.T) {
			t.Parallel()
			m := "DELETE"
			res := request(m, nil, helloHandler, t)
			if res.Body.String() != fmt.Sprintf("You requested %s request!", m) {
				t.Fatal(res.Body.String())
			}
		})
	})
}

func TestJsonHandler(t *testing.T) {
	// Make test body
	n := "Test"
	b, err := json.Marshal(JSONRequest{Name: n})
	if err != nil {
		t.Fatal(err)
	}
	res := request("POST", bytes.NewBuffer(b), jsonHandler, t)
	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	}
	resp := JSONResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		t.Errorf("errpr: %#v, res: %#v", err, res)
	}
	if resp.Message != fmt.Sprintf("Hello %s", n) {
		t.Errorf("invalid response: %#v", resp)
	}

	t.Logf("%#v", resp)
}

func request(m string, b io.Reader, h http.HandlerFunc, t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest(m, "/", b)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(h)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
	return res
}
