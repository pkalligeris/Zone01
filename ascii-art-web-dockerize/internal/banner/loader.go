package banner

import (
	"ascii-art/pkg/model"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	asciiStart  = 32  // First printable ASCII character (space)
	asciiEnd    = 126 // Last printable ASCII character (tilde)
	glyphHeight = 8   // Height of each character in the banner file
)

var (
	ErrEmptyBanner     = errors.New("banner file is empty")
	ErrMalformedBanner = errors.New("banner file format is invalid")
)

// LoadBanner reads a banner file from disk and parses it into a rune->glyph map.
func LoadBanner(filename string) (model.Banner, error) {
	// Read the entire file content
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read banner file %q: %w", filename, err)
	}
	if len(data) == 0 {
		return nil, ErrEmptyBanner
	}

	// Normalize line endings to \n to handle Windows/Mac files correctly
	normalized := strings.ReplaceAll(string(data), "\r\n", "\n")
	normalized = strings.ReplaceAll(normalized, "\r", "\n")
	lines := strings.Split(normalized, "\n")

	b := make(model.Banner, asciiEnd-asciiStart+1)
	index := 0

	// Skip the initial empty line if present (common in standard banner formats)
	// Some banner files start with a newline, others don't. We adjust the index accordingly.
	if len(lines) > 0 && lines[0] == "" {
		index = 1
	}

	// Iterate through the printable ASCII range
	for ascii := asciiStart; ascii <= asciiEnd; ascii++ {
		if index+glyphHeight > len(lines) {
			return nil, ErrMalformedBanner
		}

		glyph := make([]string, glyphHeight)
		copy(glyph, lines[index:index+glyphHeight])
		b[rune(ascii)] = glyph
		index += glyphHeight

		// Skip the empty line separator between glyphs, if it exists
		// Standard banners have an empty line separating each character block.
		if index < len(lines) && lines[index] == "" {
			index++
		}
	}

	// Ensure there are no extra non-empty lines at the end of the file
	for ; index < len(lines); index++ {
		if lines[index] != "" {
			return nil, ErrMalformedBanner
		}
	}

	return b, nil
}
