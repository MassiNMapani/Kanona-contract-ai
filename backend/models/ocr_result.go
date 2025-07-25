package models

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
}
