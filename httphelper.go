package main

import (
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

// RegexpHandler manages array route
type RegexpHandler struct {
	routes []*route
}

// Handler adds handler and path as route
func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

// HandleFunc adds handler func and path as route
func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := h.FindHandler(r.URL.Path); nil != handler {
		handler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

// FindHandler looks for a handler from path
func (h *RegexpHandler) FindHandler(path string) http.Handler {
	for _, route := range h.routes {
		if route.pattern.MatchString(path) {
			return route.handler
		}
	}
	return nil
}
