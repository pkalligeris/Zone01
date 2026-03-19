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

func processASCIIArt(req asciiArtRequest) (string, *appError) {
	text := normalizeNewlines(req.Text)
	bannerName := strings.TrimSpace(req.Banner)
	align := strings.ToLower(strings.TrimSpace(req.Align))
	color := strings.TrimSpace(req.Color)

	if align == "" {
		align = "left"
	}

	if text == "" || bannerName == "" {
		return "", &appError{
			Status:  http.StatusBadRequest,
			Message: "Text and banner selection are required",
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
			Message: "Invalid characters in input. Only ASCII characters (32-126) are allowed",
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
		Color:       color,
		ColorSubstr: req.ColorSubstring,
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

func normalizeNewlines(text string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	return strings.ReplaceAll(text, "\r", "\n")
}

func isValidASCII(text string) bool {
	for _, ch := range text {
		if (ch < 32 || ch > 126) && ch != '\n' {
			return false
		}
	}
	return true
}

func isValidBanner(name string) bool {
	switch name {
	case "standard", "shadow", "thinkertoy":
		return true
	default:
		return false
	}
}

func isValidAlign(align string) bool {
	switch align {
	case "left", "center", "right", "justify":
		return true
	default:
		return false
	}
}
