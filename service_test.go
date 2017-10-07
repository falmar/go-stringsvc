package main

import (
	"testing"
	"context"
)

// ------------------- Count ------------------

func TestStringService_Count(t *testing.T) {
	ctx := context.Background()
	s := stringService{}

	i := s.Count(ctx, "hello world")

	if i != 11 {
		t.Fatalf("Expected count of: %d; got: %d", 11, i)
	}
}

// ------------------- Lower Case ------------------

func TestStringService_LowerCase(t *testing.T) {
	ctx := context.Background()
	s := stringService{}

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
	s := stringService{}

	_, err := s.LowerCase(ctx, "")

	if err != EmptyStringErr {
		t.Fatalf("Expected error: %s: got: %s", EmptyStringErr, err)
	}
}

// ------------------- Upper Case ------------------

func TestStringService_UpperCase(t *testing.T) {
	ctx := context.Background()
	s := stringService{}

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
	s := stringService{}

	_, err := s.UpperCase(ctx, "")

	if err != EmptyStringErr {
		t.Fatalf("Expected error: %s: got: %s", EmptyStringErr, err)
	}
}
