package gormmodels

import "github.com/google/uuid"

type Transfer struct {
	Base
	Currency        string  `json:"currency" gorm:"not_null"`
	Amount          int64   `json:"amount" gorm:"not_null"`
	ConvertedAmount float32 `json:"converted_amount" gorm:"not_null" sql:"type:decimal(12, 2)"`
	Description     string  `json:"description" sql:"type:text"`
	Source          Customer
	SourceID        uuid.UUID `json:"source_id"`
	Destination     Customer
	DestinationID   uuid.UUID `json:"destination_id"`
}
