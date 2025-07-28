package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"backend/services"
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
	contracts, err := services.GetAllContracts()
	if err != nil {
		http.Error(w, "Failed to retrieve contracts", http.StatusInternalServerError)
		fmt.Println("❌ Error fetching contracts:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contracts)
}
