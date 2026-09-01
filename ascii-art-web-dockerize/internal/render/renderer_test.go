package render

import (
	"ascii-art/pkg/model"
	"testing"
)

func TestRender_Simple(t *testing.T) {
	// 1. Mock the Banner (Data Layer)
	// We manually construct a banner with 1 character ('A') and 8 dummy lines.
	mockBanner := model.Banner{
		'A': []string{"Line1", "Line2", "Line3", "Line4", "Line5", "Line6", "Line7", "Line8"},
	}

	// 2. Define Input and Expected Output
	input := "A"
	expected := "Line1\nLine2\nLine3\nLine4\nLine5\nLine6\nLine7\nLine8"

	// 3. Call the function (This will fail to compile initially because Render doesn't exist)
	got, err := Render(&model.Config{Input: input}, mockBanner)
	if err != nil {
		t.Fatalf("Render() unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("Render() = %q, want %q", got, expected)
	}
}

func TestRender_MultiLine(t *testing.T) {
	// Verify that newlines in the input result in vertically stacked ASCII blocks.
	mockBanner := model.Banner{
		'A': []string{"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
		'B': []string{"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8"},
	}

	input := "A\nB"
	expected := "A1\nA2\nA3\nA4\nA5\nA6\nA7\nA8\n" +
		"B1\nB2\nB3\nB4\nB5\nB6\nB7\nB8"

	got, err := Render(&model.Config{Input: input}, mockBanner)
	if err != nil {
		t.Fatalf("Render() unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("Render() = %q, want %q", got, expected)
	}
}

func TestRender_EmptyLines(t *testing.T) {
	mockBanner := model.Banner{}

	// Test case: Single newline should produce a single newline output
	input := "\n"
	expected := "\n"
	got, err := Render(&model.Config{Input: input}, mockBanner)
	if err != nil {
		t.Fatalf("Render() unexpected error: %v", err)
	}

	if got != expected {
		t.Errorf("Render() = %q, want %q", got, expected)
	}
}

func TestRender_InvalidColor(t *testing.T) {
	mockBanner := model.Banner{}
	config := &model.Config{Input: "A", Color: "invalid"}

	_, err := Render(config, mockBanner)
	if err == nil {
		t.Error("Render() expected error for invalid color, got nil")
	} else if err.Error() != ColorUsage {
		t.Errorf("Render() error = %q, want %q", err.Error(), ColorUsage)
	}
}