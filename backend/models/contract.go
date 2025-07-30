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
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FileName          string             `bson:"filename" json:"filename"`
	Status            string             `bson:"status" json:"status"`
	UploadedAt        time.Time          `bson:"uploadedAt" json:"uploadedAt"`
	Type              string             `bson:"type,omitempty" json:"type,omitempty"`
	StartDate         time.Time          `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate           time.Time          `bson:"endDate,omitempty" json:"endDate,omitempty"`
	Tariff            float64            `bson:"tariff,omitempty" json:"tariff,omitempty"`
	Volume            int                `bson:"volume,omitempty" json:"volume,omitempty"`
	RenegotiationDate time.Time          `bson:"renegotiationDate,omitempty" json:"renegotiationDate,omitempty"`
}
