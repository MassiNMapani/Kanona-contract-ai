// // package middleware

// // import (
// // 	"log"
// // 	"net/http"
// // )

// // // RoleMiddleware checks if the user has the required role(s)
// // func RoleMiddleware(requiredRoles ...string) func(http.Handler) http.Handler {
// // 	return func(next http.Handler) http.Handler {
// // 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// // 			userRole := r.Header.Get("X-User-Role") // This should come from your auth logic

// //				for _, role := range requiredRoles {
// //					if userRole == role {
// //						next.ServeHTTP(w, r)
// //						return
// //					}
// //				}
// //				log.Printf("🚨 Received request with role: %s", userRole)
// //				http.Error(w, "Forbidden: insufficient role", http.StatusForbidden)
// //			})
// //		}
// //	}
// package middleware

// import (
// 	"backend/utils"
// 	"log"
// 	"net/http"
// )

// // ✅ No more local CustomClaims or local contextKey

// // RoleMiddleware checks if the user has the required role(s)
// func RoleMiddleware(requiredRoles ...string) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			log.Printf("🧪 RoleMiddleware: context value = %#v", r.Context().Value(utils.UserClaimsKey))

// 			// ✅ Get claims from context using utils.UserClaimsKey
// 			claims, ok := r.Context().Value(utils.UserClaimsKey).(*utils.Claims)
// 			if !ok || claims == nil {
// 				log.Println("🚫 RoleMiddleware: No JWT claims found in context")
// 				http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 				return
// 			}

// 			log.Printf("🔐 RoleMiddleware: Received request with role=%s\n", claims.Role)

// 			// ✅ Check if role is allowed
// 			for _, role := range requiredRoles {
// 				if claims.Role == role {
// 					next.ServeHTTP(w, r)
// 					return
// 				}
// 			}

// 			log.Printf("🚫 RoleMiddleware: Role %s is not allowed\n", claims.Role)
// 			http.Error(w, "Forbidden", http.StatusForbidden)
// 		})
// 	}
// }

package middleware

import (
	"backend/utils"
	"log"
	"net/http"
)

func RoleMiddleware(requiredRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("🧪 RoleMiddleware: context value = %#v", r.Context().Value(utils.UserClaimsKey))
			claims, ok := r.Context().Value(utils.UserClaimsKey).(*utils.Claims)
			if !ok || claims == nil {
				log.Printf("🚫 RoleMiddleware: No JWT claims found in context")
				http.Error(w, "Forbidden: no valid token", http.StatusForbidden)
				return
			}

			log.Printf("🔐 RoleMiddleware: Received request with role=%s", claims.Role)

			for _, role := range requiredRoles {
				if claims.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Forbidden: insufficient role", http.StatusForbidden)
		})
	}
}
