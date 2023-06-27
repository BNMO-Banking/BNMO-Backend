package gormmodels

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transfer struct {
	Base
	Currency        string          `gorm:"not null"`
	Amount          int64           `gorm:"not null"`
	ConvertedAmount decimal.Decimal `gorm:"not null; type:numeric"`
	Description     string          `sql:"type:text"`
	Source          Customer
	SourceID        uuid.UUID
	Destination     Customer
	DestinationID   uuid.UUID
}
