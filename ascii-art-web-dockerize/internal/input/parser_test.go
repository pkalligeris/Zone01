package input

import (
	"errors"
	"testing"
)

func TestParseInput_NoArgs(t *testing.T) {
	// Verify that providing no arguments returns a usage error.
	_, err := ParseInput([]string{})
	if !errors.Is(err, ErrUsage) {
		t.Fatalf("expected ErrUsage, got %v", err)
	}
}

func TestParseInput_TooManyArgs(t *testing.T) {
	// Verify that providing more than one argument returns a usage error.
	_, err := ParseInput([]string{"one", "two"})
	if !errors.Is(err, ErrUsage) {
		t.Fatalf("expected ErrUsage, got %v", err)
	}
}

func TestParseInput_ValidString(t *testing.T) {
	// Verify that a valid ASCII string is returned as-is.
	got, err := ParseInput([]string{"Hello 123 !"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "Hello 123 !" {
		t.Fatalf("ParseInput() = %q, want %q", got, "Hello 123 !")
	}
}

func TestParseInput_EscapedNewline(t *testing.T) {
	// Verify that literal "\n" sequences are converted to actual newlines.
	got, err := ParseInput([]string{"Hello\\nWorld"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "Hello\nWorld" {
		t.Fatalf("ParseInput() = %q, want %q", got, "Hello\nWorld")
	}
}

func TestParseInput_InvalidChar(t *testing.T) {
	// Verify that non-ASCII characters trigger an error.
	_, err := ParseInput([]string{"Hell€"})
	if !errors.Is(err, ErrInvalidASCII) {
		t.Fatalf("expected ErrInvalidASCII, got %v", err)
	}
}

func TestParseInput_EmptyString(t *testing.T) {
	// Verify that an empty string argument is handled correctly.
	got, err := ParseInput([]string{""})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "" {
		t.Fatalf("ParseInput() = %q, want empty string", got)
	}
}
