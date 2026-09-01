package banner

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestLoadBanner_MissingFile(t *testing.T) {
	// Ensure that trying to load a non-existent file returns an error.
	_, err := LoadBanner("nonexistent.txt")
	if err == nil {
		t.Fatal("expected an error for missing file, got nil")
	}
}

func TestLoadBanner_EmptyFile(t *testing.T) {
	t.Parallel()

	// Create a temporary empty file
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.txt")
	if err := os.WriteFile(path, []byte{}, 0o644); err != nil {
		t.Fatalf("failed writing empty fixture: %v", err)
	}

	// Ensure that loading an empty file returns a specific error.
	_, err := LoadBanner(path)
	if !errors.Is(err, ErrEmptyBanner) {
		t.Fatalf("expected ErrEmptyBanner, got %v", err)
	}
}

func TestLoadBanner_ValidParsing(t *testing.T) {
	t.Parallel()

	// Create a temporary valid banner file
	dir := t.TempDir()
	path := filepath.Join(dir, "banner.txt")
	if err := os.WriteFile(path, []byte(buildFixtureBanner()), 0o644); err != nil {
		t.Fatalf("failed writing fixture: %v", err)
	}

	// Verify that a correctly formatted banner file is parsed into the expected map structure.
	b, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("LoadBanner returned error: %v", err)
	}

	// Check total glyph count (ASCII 32-126 = 95 chars)
	if len(b) != 95 {
		t.Fatalf("expected 95 glyphs, got %d", len(b))
	}

	// Verify content of specific glyphs (Space)
	spaceGlyph, ok := b[' ']
	if !ok {
		t.Fatal("expected glyph for space rune")
	}
	if len(spaceGlyph) != 8 {
		t.Fatalf("expected 8 rows for space glyph, got %d", len(spaceGlyph))
	}
	if spaceGlyph[0] != "32-0" || spaceGlyph[7] != "32-7" {
		t.Fatalf("unexpected space glyph contents: %v", spaceGlyph)
	}

	// Verify content of specific glyphs (Tilde)
	tildeGlyph, ok := b['~']
	if !ok {
		t.Fatal("expected glyph for '~' rune")
	}
	if tildeGlyph[0] != "126-0" || tildeGlyph[7] != "126-7" {
		t.Fatalf("unexpected '~' glyph contents: %v", tildeGlyph)
	}
}

// buildFixtureBanner generates a deterministic banner file content for testing.
// It creates 8 lines for each character from ASCII 32 to 126.
func buildFixtureBanner() string {
	var sb strings.Builder
	sb.WriteString("\n")
	for ascii := 32; ascii <= 126; ascii++ {
		for row := 0; row < 8; row++ {
			sb.WriteString(strconv.Itoa(ascii))
			sb.WriteString("-")
			sb.WriteString(strconv.Itoa(row))
			sb.WriteString("\n")
		}
		if ascii < 126 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
