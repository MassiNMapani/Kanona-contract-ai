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

	// Prepare the JSON body
	body, err := json.Marshal(map[string]string{
		"filename": filename,
	})
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}

	// Send POST request to OCR Python service
	resp, err := http.Post(ocrURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("OCR service request failed: %v", err)
	}
	defer resp.Body.Close()

	// Optional: print or parse response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("OCR service returned status %v", resp.Status)
	}

	fmt.Println("✅ OCR analysis triggered successfully for:", filename)
	return nil
}
