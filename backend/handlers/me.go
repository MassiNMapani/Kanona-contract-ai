package handlers

import (
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	// Get JWT from cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized: token missing", http.StatusUnauthorized)
		return
	}

	// Parse token
	// claims, err := utils.ParseJWT(cookie.Value)
	// if err != nil {
	// 	http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
	// 	return
	// }

	claims, err := utils.ParseJWT(cookie.Value)
	if err != nil {
		log.Printf("❌ Error decoding token: %v", err)
	} else {
		log.Printf("🧾 Token claims: email=%s | role=%s | exp=%v", claims.Email, claims.Role, claims.ExpiresAt)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"email": claims.Email,
		"role":  claims.Role,
	})

}
