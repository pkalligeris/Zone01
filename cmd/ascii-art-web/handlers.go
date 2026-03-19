package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

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

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	req := asciiArtRequest{
		Text:   r.FormValue("text"),
		Banner: r.FormValue("banner"),
		Color:  r.FormValue("color"),
		Align:  "left",
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

func apiASCIIArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, asciiArtResponse{Error: "Method not allowed"})
		return
	}

	var req asciiArtRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, asciiArtResponse{Error: "Invalid JSON body"})
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

func writeJSON(w http.ResponseWriter, status int, payload asciiArtResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
