package main

import (
	"log"
	"net/http"

	"ascii-art/internal/web"
)

// main keeps the entrypoint intentionally small: route wiring lives here,
// while validation and render behavior stay in the helper files beside it.
func main() {
	web.RegisterRoutes()

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
