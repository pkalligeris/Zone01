package main

import (
	"html/template"
	"log"
	"os"
)

// tmpl is parsed once at startup so every handler can reuse the same template.
var tmpl = mustLoadTemplate()

// mustLoadTemplate turns a startup configuration problem into an immediate
// process exit, which is appropriate because the web server cannot function
// without its only HTML template.
func mustLoadTemplate() *template.Template {
	tpl, err := loadTemplate()
	if err != nil {
		log.Fatal("Failed to load templates:", err)
	}
	return tpl
}

// loadTemplate tries both the repo-root path and the package-test path.
// The fallback keeps `go test ./cmd/ascii-art-web` working even though tests
// execute with a different current working directory than `go run`.
func loadTemplate() (*template.Template, error) {
	candidates := []string{
		"templates/index.html",
		"../../templates/index.html",
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return template.ParseFiles(path)
		}
	}

	return nil, os.ErrNotExist
}
