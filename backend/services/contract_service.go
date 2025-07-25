package services

import (
	"backend/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"backend/utils"

	// "kanona-contract-ai/backend/database"
	// "kanona-contract-ai/backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// var contractsColl *mongo.Collection = utils.DB.Collection("contracts")
func getContractCollection() *mongo.Collection {
	return utils.DB.Collection("contracts")
}

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

func SaveContractMetadata(fileName string) error {
	collection := utils.DB.Collection("contracts")

	contract := models.Contract{
		FileName:   fileName,
		UploadedAt: time.Now(),
		Status:     "Uploaded",
	}

	_, err := collection.InsertOne(context.Background(), contract)
	return err
}

func AnalyzeContract(filename string) error {
	ocrURL := "http://localhost:5001/analyze"

	body, err := json.Marshal(map[string]string{
		"filename": filename,
	})
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}

	resp, err := http.Post(ocrURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("OCR service request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("OCR service returned status %v", resp.Status)
	}

	var result models.OCRResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return fmt.Errorf("failed to decode OCR response: %v", err)
	}

	// Add filename to result for tracking
	result.Filename = filename

	// Save OCR result to MongoDB
	collection := utils.DB.Collection("ocr_results")
	_, err = collection.InsertOne(context.Background(), result)
	if err != nil {
		return fmt.Errorf("failed to insert OCR result into DB: %v", err)
	}

	fmt.Println("✅ OCR result saved to MongoDB for:", filename)
	return nil
}
