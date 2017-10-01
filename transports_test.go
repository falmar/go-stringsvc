package main

import (
	"testing"
	"context"
	"net/http"
	"bytes"
	"strings"
)

func TestDecodeRequestBody(t *testing.T) {
	ctx := context.Background()
	expectedString := "hello world"

	// mock incoming request
	body := bytes.NewReader([]byte(`{"string": "hello world"}`, ))
	httpRequest, _ := http.NewRequest("", "", body)

	// decode
	request, err := decodeHttpRequest(ctx, httpRequest)

	if err != nil {
		t.Fatalf("Unexpected error decoding request: %s", err)
	}

	req, ok := request.(requestBody)

	if !ok {
		t.Fatalf("Expected request type: %T; got: %T", requestBody{}, req)
	}

	if req.S != expectedString {
		t.Fatalf("Expected string: %s; got: %s", expectedString, req.S)
	}
}

func TestEncodeRequestBody(t *testing.T) {
	ctx := context.Background()
	buff := bytes.NewBuffer(make([]byte, 256))
	rw := &mockHttpResponseWriter{buff: buff}
	res := responseBody{V: "hello WORLD", E: ""}
	expectedBody := `{"value":"hello WORLD"}`

	// encode
	encodeHttpResponse(ctx, rw, res)

	body := buff.String()

	if !strings.Contains(body, expectedBody) {
		t.Fatalf("Expected body: %s; got: %s", expectedBody, body)
	}
}
