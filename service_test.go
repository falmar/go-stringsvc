package main

import (
	"testing"
	"context"
)

func TestStringService_Count(t *testing.T) {
	ctx := context.Background()
	s := StringService{}

	i, err := s.Count(ctx, "hello world")

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if i != 11 {
		t.Fatalf("Expected count of: %d; got: %d", 1, i)
	}
}

func TestStringService_Count_EmptyString(t *testing.T) {
	ctx := context.Background()
	s := StringService{}

	_, err := s.Count(ctx, "")

	if err != EmptyStringErr {
		t.Fatalf("Expected error: %s: got: %s", EmptyStringErr, err)
	}
}

func TestStringService_LowerCase(t *testing.T) {
	ctx := context.Background()
	s := StringService{}

	lc, err := s.LowerCase(ctx, "heLLo woRLd")

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	expected := "hello world"

	if lc != expected {
		t.Fatalf("Expected string: %s; got: %s", expected, lc)
	}
}

func TestStringService_LoweCase_EmptyString(t *testing.T) {
	ctx := context.Background()
	s := StringService{}

	_, err := s.LowerCase(ctx, "")

	if err != EmptyStringErr {
		t.Fatalf("Expected error: %s: got: %s", EmptyStringErr, err)
	}
}

func TestStringService_UpperCase(t *testing.T) {
	ctx := context.Background()
	s := StringService{}

	uc, err := s.UpperCase(ctx, "heLLo woRLd")

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	expected := "HELLO WORLD"

	if uc != expected {
		t.Fatalf("Expected string: %s; got: %s", expected, uc)
	}
}

func TestStringService_UpperCase_EmptyString(t *testing.T) {
	ctx := context.Background()
	s := StringService{}

	_, err := s.UpperCase(ctx, "")

	if err != EmptyStringErr {
		t.Fatalf("Expected error: %s: got: %s", EmptyStringErr, err)
	}
}
