package main

import (
	"context"
	"strings"
	"errors"
)

var EmptyStringErr = errors.New("string is empty")

type StringService interface {
	UpperCase(ctx context.Context, s string) (string, error)
	LowerCase(ctx context.Context, s string) (string, error)
	Count(ctx context.Context, s string) int64
}

type stringService struct{}

func (ss stringService) UpperCase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", EmptyStringErr
	}

	return strings.ToUpper(s), nil
}

func (ss stringService) LowerCase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", EmptyStringErr
	}

	return strings.ToLower(s), nil
}

func (ss stringService) Count(_ context.Context, s string) int64 {
	return int64(len(s))
}
