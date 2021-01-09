package main

import (
	"net/http"
)

type HttpHandler struct {
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	Method      string
}

// ServeHTTP calls f(w, r).
func (f HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := f.Method
	if method != r.Method {
		//方请求方法错误
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
		return
	}
	f.HandlerFunc(w, r)
}
