package main

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"errors"
)

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// check request type requestBody
		req, ok := request.(requestBody)

		if !ok {
			return nil, errors.New("Request is not requestBody type")
		}

		// call service method
		count := svc.Count(ctx, req.S)

		return countResponseBody{V: count, E: ""}, nil
	}
}

func makeLowerCaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// check request type requestBody
		req, ok := request.(requestBody)

		if !ok {
			return nil, errors.New("Request is not requestBody type")
		}

		// call service method
		str, err := svc.LowerCase(ctx, req.S)

		if err != nil {
			return responseBody{V: "", E: err.Error()}, nil
		}

		return responseBody{V: str, E: ""}, nil
	}
}

func makeUpperCaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// check request type requestBody
		req, ok := request.(requestBody)

		if !ok {
			return nil, errors.New("Request is not requestBody type")
		}

		// call service method
		str, err := svc.UpperCase(ctx, req.S)

		if err != nil {
			return responseBody{V: "", E: err.Error()}, nil
		}

		return responseBody{V: str, E: ""}, nil
	}
}
