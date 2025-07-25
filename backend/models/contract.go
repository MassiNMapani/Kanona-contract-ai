package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contract struct {
	ID         string    `bson:"_id,omitempty" json:"id"`
	FileName   string    `bson:"file_name" json:"file_name"`
	UploadedAt time.Time `bson:"uploaded_at" json:"uploaded_at"`
	Status     string    `bson:"status" json:"status"`
}

type ContractMetadata struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FileName   string             `bson:"file_name" json:"file_name"`
	UploadedAt time.Time          `bson:"uploaded_at" json:"uploaded_at"`
	Status     string             `bson:"status" json:"status"`
}
