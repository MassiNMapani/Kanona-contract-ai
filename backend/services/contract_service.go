package services

import (
	"backend/models"
	"context"
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
