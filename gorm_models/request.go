package gormmodels

import (
	"BNMO/enum"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Request struct {
	Base
	RequestType     enum.ReqType    `gorm:"not null"`
	Currency        string          `gorm:"not null"`
	Amount          int64           `gorm:"not null"`
	ConvertedAmount decimal.Decimal `gorm:"not null; type:numeric"`
	Status          string          `gorm:"not null; default:'PENDING"`
	Remarks         string          `sql:"type:text"`
	Customer        Customer
	CustomerID      uuid.UUID
}
