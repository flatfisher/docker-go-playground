package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

type SampleServer struct{}

func (s *SampleServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := http.NewServeMux()
	m.HandleFunc("/", helloHandler)
	m.ServeHTTP(w, r)
}

func TestSampleServer(t *testing.T) {
	ts := httptest.NewServer(&SampleServer{})
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("invalid response: %v", resp)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(body) != "You requested GET request!" {
		t.Errorf("invalid body: %s", body)
	}
}
