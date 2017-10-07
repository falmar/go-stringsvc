package main

import (
	"testing"
	"context"
)

// ------------------- Count ------------------

func TestCountEndpoint(t *testing.T) {
	req := requestBody{S: "WorLD HELLO"}
	svc := stringService{}
	ctx := context.Background()
	var expectedResponseValue int64 = 11
	expectedResponseError := ""

	response, err := makeCountEndpoint(svc)(ctx, req)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// check response type countResponseBody
	res, ok := response.(countResponseBody)

	if !ok {
		t.Fatalf("Expected response type: %T; got: %T", countResponseBody{}, response)
	}

	if res.E != expectedResponseError {
		t.Fatalf("Unexpected error on response: %s", res.E)
	}

	if res.V != expectedResponseValue {
		t.Fatalf("Expected value on response: %d: got: %d", expectedResponseValue, res.V)
	}
}

// ------------------- Lower Case ------------------

func TestLowerCaseEndpoint(t *testing.T) {
	req := requestBody{S: "WorLD HELLO"}
	svc := stringService{}
	ctx := context.Background()
	expectedResponseValue := "world hello"
	expectedResponseError := ""

	// execute endpoint
	response, err := makeLowerCaseEndpoint(svc)(ctx, req)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// check response type requestBody
	res, ok := response.(responseBody)

	if !ok {
		t.Fatalf("Expected response type: %T; got: %T", responseBody{}, response)
	}

	if res.E != expectedResponseError {
		t.Fatalf("Unexpected error on response: %s", res.E)
	}

	if res.V != expectedResponseValue {
		t.Fatalf("Expected value on response: %s: got: %s", expectedResponseValue, res.V)
	}
}

func TestLowerCaseEndpoint_EmptyString(t *testing.T) {
	req := requestBody{S: ""}
	svc := stringService{}
	ctx := context.Background()
	expectedResponseValue := ""
	expectedResponseError := EmptyStringErr.Error()

	// execute endpoint
	response, err := makeLowerCaseEndpoint(svc)(ctx, req)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// check response type requestBody
	res, ok := response.(responseBody)

	if !ok {
		t.Fatalf("Expected response type: %T; got: %T", responseBody{}, response)
	}

	if res.E != expectedResponseError {
		t.Fatalf("Unexpected error on response: %s", res.E)
	}

	if res.V != expectedResponseValue {
		t.Fatalf("Expected value on response: %s: got: %s", expectedResponseValue, res.V)
	}
}

// ------------------- Upper Case ------------------

func TestUpperCaseEndpoint(t *testing.T) {
	req := requestBody{S: "hello world"}
	svc := stringService{}
	ctx := context.Background()
	expectedResponseValue := "HELLO WORLD"
	expectedResponseError := ""

	// execute endpoint
	response, err := makeUpperCaseEndpoint(svc)(ctx, req)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// check response type requestBody
	res, ok := response.(responseBody)

	if !ok {
		t.Fatalf("Expected response type: %T; got: %T", responseBody{}, response)
	}

	if res.E != expectedResponseError {
		t.Fatalf("Unexpected error on response: %s", res.E)
	}

	if res.V != expectedResponseValue {
		t.Fatalf("Expected value on response: %s: got: %s", expectedResponseValue, res.V)
	}
}

func TestUpperCaseEndpoint_EmptyString(t *testing.T) {
	req := requestBody{S: ""}
	svc := stringService{}
	ctx := context.Background()
	expectedResponseValue := ""
	expectedResponseError := EmptyStringErr.Error()

	// execute endpoint
	response, err := makeUpperCaseEndpoint(svc)(ctx, req)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// check response type requestBody
	res, ok := response.(responseBody)

	if !ok {
		t.Fatalf("Expected response type: %T; got: %T", responseBody{}, response)
	}

	if res.E != expectedResponseError {
		t.Fatalf("Unexpected error on response: %s", res.E)
	}

	if res.V != expectedResponseValue {
		t.Fatalf("Expected value on response: %s: got: %s", expectedResponseValue, res.V)
	}
}
