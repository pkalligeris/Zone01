package input

import (
	"errors"
	"strings"
)

var ErrUsage = errors.New("usage: go run . [STRING]")
var ErrInvalidASCII = errors.New("input contains non-ascii characters")

// ParseInput validates CLI arguments, converts escaped newlines, and enforces ASCII range.
func ParseInput(args []string) (string, error) {
	// Ensure exactly one argument is provided
	if len(args) != 1 {
		return "", ErrUsage
	}

	// Replace literal "\n" strings with actual newline characters
	parsed := strings.ReplaceAll(args[0], "\\n", "\n")

	// Validate that all characters are within the printable ASCII range (32-126)
	for _, r := range parsed {
		if r == '\n' {
			continue
		}
		if r < 32 || r > 126 {
			return "", ErrInvalidASCII
		}
	}

	return parsed, nil
}
