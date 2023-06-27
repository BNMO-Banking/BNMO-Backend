package gormmodels

import (
	"BNMO/enum"

	"github.com/google/uuid"
)

type Request struct {
	Base
	RequestType     enum.ReqType `gorm:"not null"`
	Currency        string       `gorm:"not null"`
	Amount          int64        `gorm:"not null"`
	ConvertedAmount float64      `gorm:"not null" sql:"type:decimal(12, 2)"`
	Status          string       `gorm:"not null; default:'PENDING"`
	Remarks         string       `sql:"type:text"`
	Customer        Customer
	CustomerID      uuid.UUID
}
