package render

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ColorUsage provides the usage message for the color flag.
const ColorUsage = "Usage: go run . --color=<color> <substring> \"string\"\n\nSupported formats: ANSI standard colors (red, green, blue...), Hex (#RRGGBB), RGB (rgb(r,g,b)), HSL (hsl(h,s,l))"

// standardColors maps common color names to their ANSI 3-bit/4-bit escape codes.
var standardColors = map[string]string{
	"reset":   "\033[0m",
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"orange":  "\033[38;5;208m", // Extended 256-color code for orange
}

// GetColorCode converts a color string (name, hex, rgb, hsl) to an ANSI escape sequence.
func GetColorCode(color string) string {
	color = strings.TrimSpace(color)
	if color == "" {
		return ""
	}

	// 1. Standard Named Colors
	if code, ok := standardColors[strings.ToLower(color)]; ok {
		return code
	}

	// 2. Hex Format: #RRGGBB or #RGB
	if strings.HasPrefix(color, "#") {
		return hexToANSI(color)
	}

	lower := strings.ToLower(color)

	// 3. RGB Format: rgb(255, 0, 0)
	if strings.HasPrefix(lower, "rgb(") && strings.HasSuffix(lower, ")") {
		return rgbToANSI(lower)
	}

	// 4. HSL Format: hsl(0, 100%, 50%)
	if strings.HasPrefix(lower, "hsl(") && strings.HasSuffix(lower, ")") {
		return hslToANSI(lower)
	}

	return ""
}

func hexToANSI(hex string) string {
	hex = strings.TrimPrefix(hex, "#")
	var r, g, b uint64
	var err error

	if len(hex) == 3 {
		// Expand #RGB to #RRGGBB
		if r, err = strconv.ParseUint(string(hex[0])+string(hex[0]), 16, 8); err != nil {
			return ""
		}
		if g, err = strconv.ParseUint(string(hex[1])+string(hex[1]), 16, 8); err != nil {
			return ""
		}
		if b, err = strconv.ParseUint(string(hex[2])+string(hex[2]), 16, 8); err != nil {
			return ""
		}
	} else if len(hex) == 6 {
		if r, err = strconv.ParseUint(hex[0:2], 16, 8); err != nil {
			return ""
		}
		if g, err = strconv.ParseUint(hex[2:4], 16, 8); err != nil {
			return ""
		}
		if b, err = strconv.ParseUint(hex[4:6], 16, 8); err != nil {
			return ""
		}
	} else {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func rgbToANSI(rgbStr string) string {
	// Remove "rgb(" and ")"
	content := strings.TrimSuffix(strings.TrimPrefix(rgbStr, "rgb("), ")")
	parts := strings.Split(content, ",")
	if len(parts) != 3 {
		return ""
	}

	var r, g, b int
	// Parse R, G, B
	if _, err := fmt.Sscanf(strings.TrimSpace(parts[0]), "%d", &r); err != nil {
		return ""
	}
	if _, err := fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &g); err != nil {
		return ""
	}
	if _, err := fmt.Sscanf(strings.TrimSpace(parts[2]), "%d", &b); err != nil {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func hslToANSI(hslStr string) string {
	// Remove "hsl(" and ")"
	content := strings.TrimSuffix(strings.TrimPrefix(hslStr, "hsl("), ")")
	parts := strings.Split(content, ",")
	if len(parts) != 3 {
		return ""
	}

	var h, s, l float64
	// Parse H, S, L (handling % signs)
	if _, err := fmt.Sscanf(strings.TrimSpace(parts[0]), "%f", &h); err != nil {
		return ""
	}
	sStr := strings.TrimSuffix(strings.TrimSpace(parts[1]), "%")
	if _, err := fmt.Sscanf(sStr, "%f", &s); err != nil {
		return ""
	}
	lStr := strings.TrimSuffix(strings.TrimSpace(parts[2]), "%")
	if _, err := fmt.Sscanf(lStr, "%f", &l); err != nil {
		return ""
	}

	r, g, b := hslToRGB(h, s/100.0, l/100.0)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func hslToRGB(h, s, l float64) (int, int, int) {
	var r, g, b float64

	if s == 0 {
		r, g, b = l, l, l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q
		r = hueToRGB(p, q, h/360.0+1.0/3.0)
		g = hueToRGB(p, q, h/360.0)
		b = hueToRGB(p, q, h/360.0-1.0/3.0)
	}

	return int(math.Round(r * 255)), int(math.Round(g * 255)), int(math.Round(b * 255))
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

// IdentifyColorIndices determines which characters in the input string should be colored
// based on the substring. It returns a slice of booleans corresponding to the input runes.
func IdentifyColorIndices(input, sub string) []bool {
	inputRunes := []rune(input)
	subRunes := []rune(sub)
	toColor := make([]bool, len(inputRunes))

	// If no substring is provided, color the entire string.
	if sub == "" {
		for i := range toColor {
			toColor[i] = true
		}
		return toColor
	}

	if len(subRunes) == 0 {
		return toColor
	}

	// Find all occurrences of the substring
	for i := 0; i <= len(inputRunes)-len(subRunes); i++ {
		match := true
		for j := 0; j < len(subRunes); j++ {
			if inputRunes[i+j] != subRunes[j] {
				match = false
				break
			}
		}
		if match {
			for j := 0; j < len(subRunes); j++ {
				toColor[i+j] = true
			}
		}
	}

	return toColor
}

// ApplyColor wraps a string with the given ANSI color code and resets it afterwards.
func ApplyColor(str, colorCode string) string {
	if colorCode == "" {
		return str
	}
	return colorCode + str + "\033[0m"
}
