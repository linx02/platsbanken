package main

import (
	"log"
	"net/http"
	"platsbanken-api/db"
	"platsbanken-api/server"
	"platsbanken-api/service"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database
	err := db.InitDB("jobs.db")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	service.InitDB(db.GetDB())

	r := mux.NewRouter()

	// CORS middleware handler
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Allow requests from all origins
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// Handle preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	// Define API endpoints using handlers from server package
	r.HandleFunc("/initial-download", server.InitialDownloadHandler).Methods("POST")
	r.HandleFunc("/download-progress", server.GetDownloadProgressHandler).Methods("GET")
	r.HandleFunc("/quick-update", server.QuickUpdateHandler).Methods("POST")
	r.HandleFunc("/search", server.SearchHandler).Methods("POST")
	r.HandleFunc("/job-posting/{id}", server.GetJobPostingHandler).Methods("GET")
	r.HandleFunc("/job-postings", server.GetAmountOfJobPostingsHandler).Methods("GET")
	r.HandleFunc("/job-postings/{amount}", server.GetJobPostingsHandler).Methods("GET")

	// Handle OPTIONS requests for all routes
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
	})

	// Start the server
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
