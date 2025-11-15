package handlers

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	 "embed"
)
//go:embed ../../static/*
var staticFiles embed.FS

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    index, _ := staticFiles.Open("static/index.html")
    http.ServeContent(w, r, "index.html", time.Now(), index)
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
