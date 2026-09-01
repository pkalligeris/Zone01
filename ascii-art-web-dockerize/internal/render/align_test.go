package render

import (
	"reflect"
	"testing"
)

func TestCalculatePadding(t *testing.T) {
	tests := []struct {
		align         string
		artWidth      int
		terminalWidth int
		expected      int
	}{
		{"center", 20, 80, 30},
		{"right", 20, 80, 60},
		{"left", 20, 80, 0},
		{"justify", 20, 80, 0}, // Justify handled in renderer
		{"unknown", 20, 80, 0},
		{"center", 100, 80, 0}, // Art wider than terminal
	}

	for _, tt := range tests {
		got := calculatePadding(tt.align, tt.artWidth, tt.terminalWidth)
		if got != tt.expected {
			t.Errorf("calculatePadding(%q, %d, %d) = %d, want %d", tt.align, tt.artWidth, tt.terminalWidth, got, tt.expected)
		}
	}
}

func TestApplyAlign_Right(t *testing.T) {
	lines := []string{"A", "BB"}
	termWidth := 10
	want := []string{"        A", "        BB"}
	got := applyAlign(lines, "right", termWidth)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("applyAlign(right) = %v, want %v", got, want)
	}
}

func TestApplyAlign_Left(t *testing.T) {
	lines := []string{"A", "BB"}
	termWidth := 10
	want := []string{"A", "BB"}
	got := applyAlign(lines, "left", termWidth)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("applyAlign(left) = %v, want %v", got, want)
	}
}