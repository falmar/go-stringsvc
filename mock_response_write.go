package main

import (
	"net/http"
	"io"
)

// implement
// http.ResponseWriter interface {
//    Header() http.Header
//    Write(b []byte) (int, error)
//    WriteHeader(s int)
// }

type mockHttpResponseWriter struct {
	buff io.Writer
}

func (mk *mockHttpResponseWriter) Header() http.Header {
	return http.Header{}
}

func (mk *mockHttpResponseWriter) Write(b []byte) (int, error) {
	return mk.buff.Write(b)
}

func (mk *mockHttpResponseWriter) WriteHeader(s int) {

}


