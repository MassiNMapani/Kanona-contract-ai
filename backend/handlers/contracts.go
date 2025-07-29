package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"backend/services"
	"backend/utils"
)

func UploadContract(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data (limit to 10MB)
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		fmt.Println("❌ Error parsing multipart form:", err)
		return
	}

	// Get file from "file" field
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
		fmt.Println("❌ Error retrieving file:", err)
		return
	}
	defer file.Close()

	// Ensure uploads directory exists
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
		fmt.Println("❌ Error creating uploads directory:", err)
		return
	}

	// Create destination file
	destPath := "./uploads/" + handler.Filename
	f, err := os.Create(destPath)
	if err != nil {
		http.Error(w, "Error saving uploaded file", http.StatusInternalServerError)
		fmt.Println("❌ Error creating file:", err)
		return
	}
	defer f.Close()

	// Copy uploaded file to destination
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error writing file", http.StatusInternalServerError)
		fmt.Println("❌ Error copying file contents:", err)
		return
	}

	// Save metadata to MongoDB
	err = services.SaveContractMetadata(handler.Filename)
	if err != nil {
		http.Error(w, "Failed to store contract metadata", http.StatusInternalServerError)
		fmt.Println("❌ Error saving metadata:", err)
		return
	}

	// Trigger OCR analysis via Python
	err = services.AnalyzeContract(handler.Filename)
	if err != nil {
		fmt.Println("⚠️ Failed to analyze contract:", err)
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "✅ File uploaded and analyzed successfully",
		"filename": handler.Filename,
	})
}

func GetAllContracts(w http.ResponseWriter, r *http.Request) {
	// ✅ STEP 1: Import context and JWTClaims
	// Add at the top of the file if not already there:
	// "backend/middleware"
	// "backend/utils"
	// "log"

	// func GetAllContracts(w http.ResponseWriter, r *http.Request) {
	// contracts, err := services.GetAllContracts()
	// if err != nil {
	// 	http.Error(w, "Failed to retrieve contracts", http.StatusInternalServerError)
	// 	fmt.Println("❌ Error fetching contracts:", err)
	// 	return
	// }

	// ✅ STEP 2: Extract user claims from context
	claims, ok := r.Context().Value(utils.UserClaimsKey).(*utils.Claims)
	if !ok || claims == nil {
		http.Error(w, "Missing or invalid JWT claims", http.StatusForbidden)
		log.Println("❌ GetAllContracts: Missing or invalid JWT claims")
		return
	}

	// ✅ STEP 3: Debug log the claims
	log.Printf("🔐 GetAllContracts: Authenticated request by role=%s, email=%s\n", claims.Role, claims.Email)

	// ✅ STEP 4: Return mock contract data
	contracts := []map[string]interface{}{
		{
			"id": "1", "name": "PPA Alpha", "type": "ppa",
			"startDate": "2024-01-01", "endDate": "2027-12-31",
			"tariff": 0.12, "volume": 10000, "renegotiationDate": "2026-01-01",
		},
		{
			"id": "2", "name": "PPA Beta", "type": "psa",
			"startDate": "2024-03-11", "endDate": "2027-12-21",
			"tariff": 1.4, "volume": 15000, "renegotiationDate": "2026-04-01",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contracts)
}
