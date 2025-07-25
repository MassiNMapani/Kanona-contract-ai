package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"backend/handlers"
	"backend/utils"
)

func main() {
	// Connect to MongoDB
	utils.ConnectDB()

	// Set up router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/upload", handlers.UploadContract).Methods("POST")
	router.HandleFunc("/contracts", handlers.GetAllContracts).Methods("GET")

	// Log available routes (for debug)
	log.Println("✅ Available routes:")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		log.Println(" -", path)
		return nil
	})

	// Start server
	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Simple GET /health endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Kanona Backend API is running"))
}
