package render

import (
	"testing"
)

func TestGetColorCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Standard Colors
		{"Red", "red", "\033[31m"},
		{"Blue", "blue", "\033[34m"},
		{"Reset", "reset", "\033[0m"},
		{"Case Insensitive", "ReD", "\033[31m"},

		// Hex Colors
		{"Hex Red", "#ff0000", "\033[38;2;255;0;0m"},
		{"Hex Green", "#00ff00", "\033[38;2;0;255;0m"},
		{"Hex Short", "#f00", "\033[38;2;255;0;0m"},

		// RGB Colors
		{"RGB Red", "rgb(255, 0, 0)", "\033[38;2;255;0;0m"},
		{"RGB Spacing", "rgb( 0 , 255 , 0 )", "\033[38;2;0;255;0m"},

		// HSL Colors
		{"HSL Red", "hsl(0, 100%, 50%)", "\033[38;2;255;0;0m"},
		{"HSL Green", "hsl(120, 100%, 50%)", "\033[38;2;0;255;0m"},
		{"HSL Blue", "hsl(240, 100%, 50%)", "\033[38;2;0;0;255m"},
		{"HSL White", "hsl(0, 0%, 100%)", "\033[38;2;255;255;255m"},

		// Invalid/Empty
		{"Empty", "", ""},
		{"Invalid Name", "potato", ""},
		{"Invalid Hex", "#zzzzzz", ""},
		{"Invalid RGB", "rgb(255, 0)", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetColorCode(tt.input)
			if got != tt.expected {
				t.Errorf("GetColorCode(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestIdentifyColorIndices(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sub      string
		expected []bool
	}{
		{
			name:     "Full String",
			input:    "ABC",
			sub:      "",
			expected: []bool{true, true, true},
		},
		{
			name:     "Substring Match",
			input:    "Hello",
			sub:      "ll",
			expected: []bool{false, false, true, true, false},
		},
		{
			name:     "No Match",
			input:    "Hello",
			sub:      "x",
			expected: []bool{false, false, false, false, false},
		},
		{
			name:     "Multiple Matches",
			input:    "banana",
			sub:      "a",
			expected: []bool{false, true, false, true, false, true},
		},
		{
			name:     "Unicode Support",
			input:    "Héllo",
			sub:      "é",
			expected: []bool{false, true, false, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IdentifyColorIndices(tt.input, tt.sub)
			if len(got) != len(tt.expected) {
				t.Fatalf("IdentifyColorIndices() length = %d, want %d", len(got), len(tt.expected))
			}
			for i, v := range got {
				if v != tt.expected[i] {
					t.Errorf("IdentifyColorIndices()[%d] = %v, want %v", i, v, tt.expected[i])
				}
			}
		})
	}
}

func TestApplyColor(t *testing.T) {
	got := ApplyColor("A", "\033[31m")
	want := "\033[31mA\033[0m"
	if got != want {
		t.Errorf("ApplyColor() = %q, want %q", got, want)
	}

	gotEmpty := ApplyColor("A", "")
	if gotEmpty != "A" {
		t.Errorf("ApplyColor(empty) = %q, want %q", gotEmpty, "A")
	}
}
