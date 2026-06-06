package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ilahazs/yt-webui/backend/api"
	"github.com/ilahazs/yt-webui/backend/config"
)

// corsMiddleware injects CORS headers and handles preflight OPTIONS requests.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// 1. Load configuration
	cfg := config.Load()

	// 2. Ensure directories exist
	if err := os.MkdirAll(cfg.DataDir, 0755); err != nil {
		log.Printf("Warning: Failed to create data directory %s: %v", cfg.DataDir, err)
	}
	if err := os.MkdirAll(cfg.DownloadDir, 0755); err != nil {
		log.Printf("Warning: Failed to create download directory %s: %v", cfg.DownloadDir, err)
	}

	// 3. Register handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		api.HandleHealth(w, r, cfg)
	})

	// 4. Start HTTP server
	log.Printf("Starting backend server on %s...", cfg.BindAddress)
	if err := http.ListenAndServe(cfg.BindAddress, corsMiddleware(mux)); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
