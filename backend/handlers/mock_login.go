package handlers

import (
	"fmt"
	"net/http"
)

func MockLogin(w http.ResponseWriter, r *http.Request) {
	role := r.URL.Query().Get("role")
	if role == "" {
		role = "viewer"
	}

	// Normally you'd set a session or return a JWT here
	w.Header().Set("X-User-Role", role)
	fmt.Fprintf(w, "Logged in as role: %s", role)
}
