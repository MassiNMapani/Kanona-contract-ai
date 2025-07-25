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
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	f, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	fmt.Fprintf(w, "Uploaded file: %s\n", handler.Filename)

	// Save to MongoDB
	err = services.SaveContractMetadata(handler.Filename)
	if err != nil {
		http.Error(w, "Failed to store metadata", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Uploaded and stored metadata for file: %s\n", handler.Filename)
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
