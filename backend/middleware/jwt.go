// package middleware

// import (
// 	"backend/utils"
// 	"context"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// type contextKey string

// const UserClaimsKey = contextKey("user")

// func JWTMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		publicPaths := []string{"/mock-login", "/health", "/me"}

// 		// Allow unauthenticated access to public paths
// 		for _, path := range publicPaths {
// 			if strings.HasPrefix(r.URL.Path, path) {
// 				next.ServeHTTP(w, r)
// 				return
// 			}
// 		}

// 		// // 🛡️ Require token for all other routes
// 		// cookie, err := r.Cookie("token")
// 		// if err != nil {
// 		// 	http.Error(w, "Unauthorized: missing auth token", http.StatusUnauthorized)
// 		// 	return
// 		// }

// 		// claims, err := utils.ParseJWT(cookie.Value)
// 		// if err != nil || claims == nil {
// 		// 	http.Error(w, "Unauthorized: missing or invalid token", http.StatusUnauthorized)
// 		// 	return
// 		// }

// 		cookie, err := r.Cookie("token")
// 		if err != nil {
// 			log.Println("❌ JWTMiddleware: No token cookie found")
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		log.Println("🔍 JWTMiddleware: Token cookie found")

// 		claims, err := utils.ParseJWT(cookie.Value)
// 		if err != nil {
// 			log.Println("❌ JWTMiddleware: Failed to parse JWT:", err)
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		log.Printf("✅ JWTMiddleware: Parsed token for email=%s role=%s\n", claims.Email, claims.Role)

// 		// ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
// 		// next.ServeHTTP(w, r.WithContext(ctx))

// 		// Add claims to request context
// 		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// package middleware

// import (
// 	"context"
// 	"net/http"

// 	"backend/utils"
// )

// // ✅ Define contextKey and UserClaimsKey in middleware scope
// type contextKey string

// const UserClaimsKey = contextKey("user") // Used to store JWT claims in request context

// func JWTMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("token")
// 		if err != nil {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		claims, err := utils.ParseJWT(cookie.Value)
// 		if err != nil {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		// ✅ Store parsed claims in request context using correct key
// 		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

package middleware

import (
	"context"
	"log"
	"net/http"

	"backend/utils"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			log.Println("❌ JWTMiddleware: No token cookie found")
			next.ServeHTTP(w, r)
			return
		}

		log.Println("🔍 JWTMiddleware: Token cookie found")

		claims, err := utils.ParseJWT(cookie.Value)
		if err != nil {
			log.Println("❌ JWTMiddleware: Failed to parse JWT:", err)
			next.ServeHTTP(w, r)
			return
		}

		log.Printf("✅ JWTMiddleware: Parsed token for email=%s role=%s", claims.Email, claims.Role)

		ctx := context.WithValue(r.Context(), utils.UserClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
