package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// maxBodyBytes caps API request bodies to 32 KB. Requests larger than this
// are rejected before parsing so the server cannot be forced into allocating
// unbounded memory for a single request.
const maxBodyBytes = 32 * 1024

// RegisterRoutes links standard HTTP endpoints to the internal handlers.
func RegisterRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.HandleFunc("/api/ascii-art", apiASCIIArtHandler)
	http.HandleFunc("/export", exportHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}

// homeHandler serves the initial page and rejects unknown paths so `/`
// behaves like a real route rather than a prefix match.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	data := PageData{Banner: "standard"}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

// asciiArtHandler processes the browser form submission.
// It deliberately uses the web-specific render path so the resulting HTML
// contains browser-friendly text and CSS color, not terminal ANSI escapes.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	req := asciiArtRequest{
		Text:   r.FormValue("text"),
		Banner: r.FormValue("banner"),
		Color:  r.FormValue("color"),
		Align:  r.FormValue("align"),
	}

	data := PageData{
		Input:  req.Text,
		Banner: req.Banner,
		Color:  req.Color,
	}

	result, appErr := processWebASCIIArt(req)
	if appErr != nil {
		data.Error = appErr.Message
		w.WriteHeader(appErr.Status)
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	data.Result = result
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

// apiASCIIArtHandler exposes the same render pipeline over JSON.
// Unknown fields and multiple JSON values are rejected to keep the API contract
// strict and predictable for clients.
func apiASCIIArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, asciiArtResponse{Error: "Method not allowed"})
		return
	}

	var req asciiArtRequest
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		status := http.StatusBadRequest
		if err.Error() == "http: request body too large" {
			status = http.StatusRequestEntityTooLarge
		}
		writeJSON(w, status, asciiArtResponse{Error: "Invalid JSON body"})
		return
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		writeJSON(w, http.StatusBadRequest, asciiArtResponse{Error: "Invalid JSON body"})
		return
	}

	result, appErr := processASCIIArt(req)
	if appErr != nil {
		writeJSON(w, appErr.Status, asciiArtResponse{Error: appErr.Message})
		return
	}

	writeJSON(w, http.StatusOK, asciiArtResponse{Result: result})
}

// exportHandler processes a form POST and returns the rendered ASCII art
// as a downloadable plain text file. The required HTTP headers
// (Content-Type, Content-Length, Content-Disposition) are set so the
// browser triggers a file save dialog instead of rendering the response.
func exportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	req := asciiArtRequest{
		Text:   r.FormValue("text"),
		Banner: r.FormValue("banner"),
		Align:  "left",
	}

	// Render without ANSI color so the exported file contains pure text.
	result, appErr := processWebASCIIArt(req)
	if appErr != nil {
		http.Error(w, fmt.Sprintf("%d %s", appErr.Status, appErr.Message), appErr.Status)
		return
	}

	data := []byte(result)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Content-Disposition", `attachment; filename="ascii_art.txt"`)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// writeJSON centralizes JSON responses so API handlers always emit the same
// content type and encoding behavior.
func writeJSON(w http.ResponseWriter, status int, payload asciiArtResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
