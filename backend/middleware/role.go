package middleware

import (
	"log"
	"net/http"
)

// RoleMiddleware checks if the user has the required role(s)
func RoleMiddleware(requiredRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := r.Header.Get("X-User-Role") // This should come from your auth logic

			for _, role := range requiredRoles {
				if userRole == role {
					next.ServeHTTP(w, r)
					return
				}
			}
			log.Printf("🚨 Received request with role: %s", userRole)
			http.Error(w, "Forbidden: insufficient role", http.StatusForbidden)
		})
	}
}
