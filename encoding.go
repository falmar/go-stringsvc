package main

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
