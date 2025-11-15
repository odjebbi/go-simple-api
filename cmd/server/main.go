package main

import (
	"fmt"
	"go-simple-api/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// API routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", handlers.StaticHandler()))
	http.HandleFunc("/api/hello", handlers.HelloHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/api/info", handlers.InfoHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
