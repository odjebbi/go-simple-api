package handlers

import (
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

// HomeHandler sert index.html depuis embed
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	subFS, _ := fs.Sub(staticFiles, "static")
	http.FileServer(http.FS(subFS)).ServeHTTP(w, r)
}

func StaticHandler() http.Handler {
	fsys, _ := fs.Sub(staticFiles, "static")
	return http.FileServer(http.FS(fsys))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Hello from Go!",
	}
	json.NewEncoder(w).Encode(response)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"status":  "ok",
		"service": "go-api",
	}
	json.NewEncoder(w).Encode(response)
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"language": "Go",
		"version":  "1.0.0",
	}
	json.NewEncoder(w).Encode(response)
}
