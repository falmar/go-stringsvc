package main

import (
	"encoding/json"
	"net/http"
	"context"
)

type requestBody struct {
	S string `json:"string"`
}

type responseBody  struct {
	V string `json:"value"`
	E string `json:"error,omitempty"`
}

type countResponseBody struct {
	V int64 `json:"value"`
	E string `json:"error,omitempty"`
}

func decodeHttpRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requestBody

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeHttpResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
