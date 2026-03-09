package main

import (
	"ascii-art/internal/banner"
	"ascii-art/internal/render"
	"ascii-art/pkg/model"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type PageData struct {
	Input  string
	Banner string
	Result string
	Error  string
}

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Failed to load templates:", err)
	}
}

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

	text := r.FormValue("text")
	bannerName := r.FormValue("banner")

	data := PageData{
		Input:  text,
		Banner: bannerName,
	}

	if text == "" || bannerName == "" {
		data.Error = "Text and banner selection are required"
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, data)
		return
	}

	if bannerName != "standard" && bannerName != "shadow" && bannerName != "thinkertoy" {
		data.Error = "Invalid banner selection"
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, data)
		return
	}

	for _, ch := range text {
		if (ch < 32 || ch > 126) && ch != '\n' && ch != '\r' {
			data.Error = "Invalid characters in input. Only ASCII characters (32-126) are allowed"
			w.WriteHeader(http.StatusBadRequest)
			tmpl.Execute(w, data)
			return
		}
	}

	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	bannerPath, err := banner.GetBannerPath(bannerName)
	if err != nil {
		data.Error = "Banner file not found"
		w.WriteHeader(http.StatusNotFound)
		tmpl.Execute(w, data)
		return
	}

	b, err := banner.LoadBanner(bannerPath)
	if err != nil {
		data.Error = "Failed to load banner"
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, data)
		return
	}

	cfg := &model.Config{
		Input:      text,
		BannerFile: bannerName,
		Align:      "left",
	}

	result, err := render.Render(cfg, b)
	if err != nil {
		data.Error = "Failed to render ASCII art"
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, data)
		return
	}

	data.Result = result
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
