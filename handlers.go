package main

import (
	"fmt"
	"net/http"
)

// TODO: 意図したHeaderやResponseが返却されるかテストコードを書く
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-type", "text/plain; charset=utf-8'")
	m := ""
	switch {
	case r.Method == http.MethodGet:
		m = "Get"
	case r.Method == http.MethodPost:
		m = "Post"
	case r.Method == http.MethodPut:
		m = "Put"
	case r.Method == http.MethodDelete:
		m = "Delete"
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "405 Method Not Allowed")
		return
	}
	fmt.Fprintf(w, "You requested %s request!", m)
}
