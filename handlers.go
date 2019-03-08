package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-type", "text/plain; charset=utf-8'")
	m := ""
	switch {
	case r.Method == http.MethodGet:
		m = "GET"
	case r.Method == http.MethodPost:
		m = "POST"
	case r.Method == http.MethodPut:
		m = "PUT"
	case r.Method == http.MethodDelete:
		m = "DELETE"
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "405 Method Not Allowed")
		return
	}
	fmt.Fprintf(w, "You requested %s request!", m)
}

// JSONRequest is request body from client
type JSONRequest struct {
	Name string
}

// JSONResponse is response body to ciient
type JSONResponse struct {
	Message string
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	body := JSONRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	b, err := json.Marshal(JSONResponse{Message: "Hello " + body.Name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(b)) // fmt.Fprint(w, string(b)) as same response
}
