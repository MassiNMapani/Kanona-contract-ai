package handlers

import (
	"fmt"
	"net/http"
	"time"

	"backend/utils"
)

func MockLogin(w http.ResponseWriter, r *http.Request) {
	role := r.URL.Query().Get("role")
	if role == "" {
		role = "viewer"
	}

	// Generate a JWT with that role
	tokenString, err := utils.GenerateJWT("test@example.com", role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Set JWT in cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: false, // you can later switch to true if needed
		Secure:   false, // set true if using HTTPS
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(1 * time.Hour),
	})

	fmt.Fprintf(w, "✅ Mock login successful for role: %s", role)
}
