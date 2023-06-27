package gormmodels

import (
	"BNMO/enum"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transfer struct {
	Base
	Currency        string              `gorm:"not null"`
	Amount          int64               `gorm:"not null"`
	ConvertedAmount decimal.Decimal     `gorm:"not null; type:numeric"`
	Status          enum.TransferStatus `gorm:"not null; default:'PENDING"`
	Description     string              `sql:"type:text"`
	Source          Customer            `gorm:"foreignKey:SourceID;references:ID"`
	SourceID        uuid.UUID
	Destination     Customer `gorm:"foreignKey:DestinationID;references:ID"`
	DestinationID   uuid.UUID
}
