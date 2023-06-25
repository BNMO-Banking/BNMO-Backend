package gormmodels

import "github.com/google/uuid"

type Transfer struct {
	Base
	Currency        string  `gorm:"not_null"`
	Amount          int64   `gorm:"not_null"`
	ConvertedAmount float32 `gorm:"not_null" sql:"type:decimal(12, 2)"`
	Description     string  `sql:"type:text"`
	Source          Customer
	SourceID        uuid.UUID
	Destination     Customer
	DestinationID   uuid.UUID
}
