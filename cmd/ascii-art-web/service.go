package main

import (
	"ascii-art/internal/banner"
	"ascii-art/internal/render"
	"ascii-art/pkg/model"
	"errors"
	"net/http"
	"os"
	"strings"
)

// maxTextLength is the upperlimit on input size to prevent excessively slow
// renders and oversized responses.
const maxTextLength = 2000

// processASCIIArt is used by the JSON API and preserves terminal-style color
// rendering because API clients may want the raw colored output.
func processASCIIArt(req asciiArtRequest) (string, *appError) {
	return processASCIIArtRequest(req, true)
}

// processWebASCIIArt is used by the HTML form route.
// It validates the requested color but leaves the rendered text uncolored so
// the browser can apply color via CSS instead of displaying ANSI escape codes.
func processWebASCIIArt(req asciiArtRequest) (string, *appError) {
	return processASCIIArtRequest(req, false)
}

// processASCIIArtRequest is the shared application flow for the web package:
// normalize inputs, validate them, load the banner, build render config,
// and finally delegate to the core renderer.
func processASCIIArtRequest(req asciiArtRequest, renderColor bool) (string, *appError) {
	text := normalizeNewlines(req.Text)
	bannerName := strings.TrimSpace(req.Banner)
	align := strings.ToLower(strings.TrimSpace(req.Align))
	color := strings.TrimSpace(req.Color)

	if align == "" {
		align = "left"
	}

	if strings.TrimSpace(text) == "" || bannerName == "" {
		return "", &appError{
			Status:  http.StatusBadRequest,
			Message: "Text and banner selection are required",
		}
	}

	if len(text) > maxTextLength {
		return "", &appError{
			Status:  http.StatusBadRequest,
			Message: "Input text is too long (max 2000 characters)",
		}
	}

	if !isValidBanner(bannerName) {
		return "", &appError{
			Status:  http.StatusBadRequest,
			Message: "Invalid banner selection",
		}
	}

	if !isValidASCII(text) {
		return "", &appError{
			Status:  http.StatusBadRequest,
			Message: "Your input contains special characters we don't recognize. Please stick to standard English letters, numbers, and common punctuation marks.",
		}
	}

	if !isValidAlign(align) {
		return "", &appError{
			Status:  http.StatusBadRequest,
			Message: "Invalid alignment option",
		}
	}

	if color != "" && render.GetColorCode(color) == "" {
		return "", &appError{
			Status:  http.StatusBadRequest,
			Message: "Invalid color value",
		}
	}

	bannerPath, err := banner.GetBannerPath(bannerName)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, banner.ErrUnknownBanner) {
			status = http.StatusNotFound
		}
		return "", &appError{
			Status:  status,
			Message: "Banner file not found",
		}
	}

	b, err := loadBannerWithFallbacks(bannerPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", &appError{
				Status:  http.StatusNotFound,
				Message: "Banner file not found",
			}
		}
		return "", &appError{
			Status:  http.StatusInternalServerError,
			Message: "Failed to load banner",
		}
	}

	cfg := &model.Config{
		Input:       text,
		BannerFile:  bannerName,
		Align:       align,
		ColorSubstr: req.ColorSubstring,
	}
	// Color is only injected into the render config for API/terminal output.
	// The HTML route applies the chosen color in the template layer instead.
	if renderColor {
		cfg.Color = color
	}

	result, err := render.Render(cfg, b)
	if err != nil {
		status := http.StatusInternalServerError
		if color != "" {
			status = http.StatusBadRequest
		}
		return "", &appError{
			Status:  status,
			Message: "Failed to render ASCII art",
		}
	}

	return result, nil
}

// loadBannerWithFallbacks mirrors the template-loading strategy so local runs
// and package-local tests can both find the banner assets reliably.
func loadBannerWithFallbacks(path string) (model.Banner, error) {
	candidates := []string{
		path,
		"../../" + path,
	}

	var lastErr error
	for _, candidate := range candidates {
		b, err := banner.LoadBanner(candidate)
		if err == nil {
			return b, nil
		}
		lastErr = err
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	return nil, lastErr
}

// normalizeNewlines makes form and API input consistent before validation and
// rendering by collapsing Windows and legacy Mac line endings into `\n`.
func normalizeNewlines(text string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	return strings.ReplaceAll(text, "\r", "\n")
}

// isValidASCII keeps web input aligned with the banner set, which only supports
// printable ASCII plus newline separators between rendered blocks.
func isValidASCII(text string) bool {
	for _, ch := range text {
		if (ch < 32 || ch > 126) && ch != '\n' {
			return false
		}
	}
	return true
}

// isValidBanner limits requests to the banner assets that ship with the app.
func isValidBanner(name string) bool {
	switch name {
	case "standard", "shadow", "thinkertoy":
		return true
	default:
		return false
	}
}

// isValidAlign keeps alignment options in sync with the renderer contract.
func isValidAlign(align string) bool {
	switch align {
	case "left", "center", "right", "justify":
		return true
	default:
		return false
	}
}
