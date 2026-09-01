package render

import (
	"ascii-art/pkg/model"
	"fmt"
	"strings"
)

// Render converts the input string into an ASCII art string using the provided banner.
func Render(config *model.Config, banner model.Banner) (string, error) {
	input := config.Input
	// Split input into lines to handle newlines correctly
	inputLines := strings.Split(input, "\n")
	var sb strings.Builder
	terminalWidth := terminalWidthProvider()

	colorCode := GetColorCode(config.Color)
	if config.Color != "" && colorCode == "" {
		return "", fmt.Errorf(ColorUsage)
	}
	// Pre-calculate which indices in the input string need coloring
	indicesToColor := IdentifyColorIndices(input, config.ColorSubstr)
	currentIndex := 0

	for i, line := range inputLines {
		// Add a newline between blocks of text, but not before the first one
		if i > 0 {
			sb.WriteByte('\n')
			currentIndex++ // Treat the newline as a character index for coloring purposes
		}
		// If the line is empty, we just needed the newline added above (if i > 0)
		if line == "" {
			continue
		}

		var block []string

		// Handle justify alignment: render words separately and distribute spaces
		if config.Align == "justify" {
			savedIndex := currentIndex
			var renderedWords []struct {
				lines [8]string
				width int
			}
			wordWidths := 0
			var currentWordLines [8]strings.Builder
			inWord := false
			lineIdx := 0

			// Parse and render words manually to maintain currentIndex sync
			for lineIdx < len(line) {
				char := rune(line[lineIdx])
				if char == ' ' {
					if inWord {
						var blk [8]string
						for r := 0; r < 8; r++ {
							blk[r] = currentWordLines[r].String()
							currentWordLines[r].Reset()
						}
						w := visibleWidth(blk[0])
						renderedWords = append(renderedWords, struct {
							lines [8]string
							width int
						}{blk, w})
						wordWidths += w
						inWord = false
					}
					currentIndex++
					lineIdx++
					continue
				}

				inWord = true
				if asciiLines, ok := banner[char]; ok {
					applyColor := false
					if config.Color != "" && indicesToColor[currentIndex] {
						applyColor = true
					}
					for row := 0; row < 8; row++ {
						if applyColor {
							currentWordLines[row].WriteString(ApplyColor(asciiLines[row], colorCode))
						} else {
							currentWordLines[row].WriteString(asciiLines[row])
						}
					}
				}
				currentIndex++
				lineIdx++
			}
			if inWord {
				var blk [8]string
				for r := 0; r < 8; r++ {
					blk[r] = currentWordLines[r].String()
				}
				w := visibleWidth(blk[0])
				renderedWords = append(renderedWords, struct {
					lines [8]string
					width int
				}{blk, w})
				wordWidths += w
			}

			spacesNeeded := terminalWidth - wordWidths
			if spacesNeeded >= 0 && len(renderedWords) > 1 {
				block = make([]string, 8)
				gaps := len(renderedWords) - 1
				perGap := spacesNeeded / gaps
				rem := spacesNeeded % gaps

				for row := 0; row < 8; row++ {
					var sbRow strings.Builder
					for i, word := range renderedWords {
						sbRow.WriteString(word.lines[row])
						if i < gaps {
							sp := perGap
							if i < rem {
								sp++
							}
							sbRow.WriteString(strings.Repeat(" ", sp))
						}
					}
					block[row] = sbRow.String()
				}
			} else {
				// Fallback to normal render if it doesn't fit or single word
				currentIndex = savedIndex
			}
		}

		if block == nil {
			var lines [8]strings.Builder
			// Iterate over each character in the current line of text
			for _, char := range line {
				if asciiLines, ok := banner[char]; ok {
					// Check if the current character index is marked for coloring
					applyColor := false
					if config.Color != "" && indicesToColor[currentIndex] {
						applyColor = true
					}
					// Append each of the 8 rows of the character to the corresponding builder
					for row := 0; row < 8; row++ {
						if applyColor {
							lines[row].WriteString(ApplyColor(asciiLines[row], colorCode))
						} else {
							lines[row].WriteString(asciiLines[row])
						}
					}
				}
				currentIndex++
			}
			// Combine the 8 rows into the final output builder
			block = make([]string, 8)
			for row := 0; row < 8; row++ {
				block[row] = lines[row].String()
			}
		}
		block = applyAlign(block, config.Align, terminalWidth)
		for row := 0; row < 8; row++ {
			sb.WriteString(block[row])
			if row < 7 {
				sb.WriteByte('\n')
			}
		}
	}

	return sb.String(), nil
}
