package main

import (
	"net/http"
	"reflect"
	"regexp"
	"testing"
)

func testHandler(w http.ResponseWriter, r *http.Request) {}

func TestFindHandler(t *testing.T) {
	handler := &RegexpHandler{}
	p := "/"
	handler.HandleFunc(regexp.MustCompile(p), testHandler)
	h := handler.FindHandler(p)
	if reflect.ValueOf(h).Pointer() != reflect.ValueOf(testHandler).Pointer() {
		t.Errorf("No matching handler found. %s", p)
	}
}
