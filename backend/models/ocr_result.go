package models

import "time"

type OCREntity struct {
	Label string `json:"label" bson:"label"`
	Text  string `json:"text" bson:"text"`
}

type OCRResult struct {
	Filename           string      `json:"filename" bson:"filename"`
	TextPreview        string      `json:"text" bson:"text"`
	Entities           []OCREntity `json:"entities" bson:"entities"`
	Tariffs            []string    `json:"tariffs" bson:"tariffs"`
	RenegotiationTerms string      `json:"renegotiation_terms,omitempty" bson:"renegotiation_terms,omitempty"`
	PageCount          int         `json:"pages" bson:"pages"`

	// ✅ New fields to support contract metadata update
	Type              string    `json:"type,omitempty" bson:"type,omitempty"`
	StartDate         time.Time `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate           time.Time `json:"endDate,omitempty" bson:"endDate,omitempty"`
	Tariff            float64   `json:"tariff,omitempty" bson:"tariff,omitempty"`
	Volume            float64   `json:"volume,omitempty" bson:"volume,omitempty"`
	RenegotiationDate time.Time `json:"renegotiationDate,omitempty" bson:"renegotiationDate,omitempty"`
}
