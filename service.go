package main

import (
	"context"
	"strings"
	"errors"
)

var EmptyStringErr = errors.New("String is empty")

type StringServiceInterface interface {
	UpperCase(ctx context.Context, s string) (string, error)
	LowerCase(ctx context.Context, s string) (string, error)
	Count(ctx context.Context, s string) int
}

type StringService struct {

}

func (ss *StringService) UpperCase(_ context.Context, s string) (string, error) {
	if (s == "") {
		return "", EmptyStringErr
	}

	return strings.ToUpper(s), nil
}

func (ss *StringService) LowerCase(_ context.Context, s string) (string, error) {
	if (s == "") {
		return "", EmptyStringErr
	}

	return strings.ToLower(s), nil
}

func (ss *StringService) Count(_ context.Context, s string) (int64, error) {
	if (s == "") {
		return 0, EmptyStringErr
	}

	return int64(len(s)), nil
}
