package services

import (
	"backend/models"
	"backend/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 🔄 Returns a handle to the 'contracts' collection
func getContractCollection() *mongo.Collection {
	return utils.DB.Collection("contracts")
}

// 📦 Fetches all contract metadata from MongoDB
func GetAllContracts() ([]models.ContractMetadata, error) {
	collection := getContractCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var contracts []models.ContractMetadata
	if err = cursor.All(ctx, &contracts); err != nil {
		return nil, err
	}

	return contracts, nil
}

// 💾 Saves a minimal contract metadata record (add fields later)
func SaveContractMetadata(fileName string) error {
	collection := getContractCollection()

	contract := models.ContractMetadata{
		FileName:   fileName,
		UploadedAt: time.Now(),
		Status:     "Uploaded",
	}

	_, err := collection.InsertOne(context.Background(), contract)
	if err != nil {
		return fmt.Errorf("failed to save contract metadata: %v", err)
	}
	fmt.Println("✅ Contract metadata saved:", fileName)
	return nil
}

// 🤖 Calls OCR service and stores extracted results
// func AnalyzeContract(filename string) error {
// 	ocrURL := "http://localhost:5001/analyze"

// 	payload, err := json.Marshal(map[string]string{
// 		"filename": filename,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("failed to encode OCR request: %v", err)
// 	}

// 	resp, err := http.Post(ocrURL, "application/json", bytes.NewBuffer(payload))
// 	if err != nil {
// 		return fmt.Errorf("OCR request failed: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("OCR service returned status: %v", resp.Status)
// 	}

// 	var result models.OCRResult
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	if err != nil {
// 		return fmt.Errorf("failed to decode OCR response: %v", err)
// 	}

// 	result.Filename = filename

// 	ocrCollection := utils.DB.Collection("ocr_results")
// 	_, err = ocrCollection.InsertOne(context.Background(), result)
// 	if err != nil {
// 		return fmt.Errorf("failed to store OCR result in DB: %v", err)
// 	}

//		fmt.Println("✅ OCR result saved to MongoDB for:", filename)
//		return nil
//	}
func AnalyzeContract(filename string) error {
	ocrURL := "http://localhost:5001/analyze"

	payload, err := json.Marshal(map[string]string{
		"filename": filename,
	})
	if err != nil {
		return fmt.Errorf("failed to encode OCR request: %v", err)
	}

	resp, err := http.Post(ocrURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("OCR request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("OCR service returned status: %v", resp.Status)
	}

	var result models.OCRResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return fmt.Errorf("failed to decode OCR response: %v", err)
	}

	result.Filename = filename

	// Save OCR result
	ocrCollection := utils.DB.Collection("ocr_results")
	_, err = ocrCollection.InsertOne(context.Background(), result)
	if err != nil {
		return fmt.Errorf("failed to store OCR result: %v", err)
	}
	fmt.Println("✅ OCR result saved for:", filename)

	// Extract metadata to update contract
	metadataUpdate := bson.M{}
	if result.Type != "" {
		metadataUpdate["type"] = result.Type
	}
	if !result.StartDate.IsZero() {
		metadataUpdate["startDate"] = result.StartDate
	}
	if !result.EndDate.IsZero() {
		metadataUpdate["endDate"] = result.EndDate
	}
	if result.Tariff > 0 {
		metadataUpdate["tariff"] = result.Tariff
	}
	if result.Volume > 0 {
		metadataUpdate["volume"] = result.Volume
	}
	if !result.RenegotiationDate.IsZero() {
		metadataUpdate["renegotiationDate"] = result.RenegotiationDate
	}
	metadataUpdate["status"] = "Analyzed"

	// Update contract metadata in contracts collection
	contractsColl := utils.DB.Collection("contracts")
	filter := bson.M{"filename": filename}
	update := bson.M{"$set": metadataUpdate}

	_, err = contractsColl.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update contract metadata: %v", err)
	}
	fmt.Println("📝 Contract metadata updated with OCR info for:", filename)

	return nil
}
