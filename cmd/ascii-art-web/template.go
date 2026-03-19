package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl = mustLoadTemplate()

func mustLoadTemplate() *template.Template {
	tpl, err := loadTemplate()
	if err != nil {
		log.Fatal("Failed to load templates:", err)
	}
	return tpl
}

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
