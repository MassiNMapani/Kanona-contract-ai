package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"backend/handlers"
	"backend/middleware"
	"backend/utils"
)

func main() {
	// Connect to MongoDB
	utils.ConnectDB()

	// Set up router
	router := mux.NewRouter()

	// Apply middleware
	router.Use(middleware.JWTMiddleware)

	// Dev mock login
	router.HandleFunc("/mock-login", handlers.MockLogin).Methods("GET")
	// /me route
	router.HandleFunc("/me", handlers.Me).Methods("GET")
	// Health check
	router.HandleFunc("/health", HealthCheck).Methods("GET")

	// Protected contract routes
	router.Handle("/contracts", middleware.RoleMiddleware("admin", "ceo", "hod")(http.HandlerFunc(handlers.GetAllContracts))).Methods("GET")
	router.Handle("/upload", middleware.RoleMiddleware("admin", "ppa-user", "psa-user")(http.HandlerFunc(handlers.UploadContract))).Methods("POST")

	// ✅ Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-User-Role"},
	})

	handler := c.Handler(router)

	// ✅ Log available routes BEFORE starting the server
	log.Println("✅ Available routes:")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		log.Println(" -", path)
		return nil
	})

	// ✅ Start server (only once)
	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// Health check
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Kanona Backend API is running"))
}
