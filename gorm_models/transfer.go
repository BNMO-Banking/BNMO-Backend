package gormmodels

import "github.com/google/uuid"

type Transfer struct {
	Base
	Currency        string  `gorm:"not null"`
	Amount          int64   `gorm:"not null"`
	ConvertedAmount float32 `gorm:"not null" sql:"type:decimal(12, 2)"`
	Description     string  `sql:"type:text"`
	Source          Customer
	SourceID        uuid.UUID
	Destination     Customer
	DestinationID   uuid.UUID
}
