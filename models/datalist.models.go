package models

import (
	"BNMO/enum"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PageMetadata struct {
	Total    int64   `json:"total"`
	Page     int     `json:"page"`
	LastPage float64 `json:"last_page"`
}

type AccountData struct {
	Id           uuid.UUID          `json:"id"`
	FirstName    string             `json:"first_name"`
	LastName     string             `json:"last_name"`
	PhoneNumber  string             `json:"phone_number"`
	IdCardPath   string             `json:"id_card_path"`
	Status       enum.AccountStatus `json:"status"`
	AddressLine1 string             `json:"address_line_1"`
	AddressLine2 string             `json:"address_line_2"`
	City         string             `json:"city"`
	State        string             `json:"state"`
	PostalCode   string             `json:"postal_code"`
	Country      string             `json:"country"`
}

type AccountDataList struct {
	Data     []AccountData `json:"data"`
	Metadata PageMetadata  `json:"metadata"`
}

type RequestData struct {
	Id              uuid.UUID       `json:"id"`
	RequestType     enum.ReqType    `json:"request_type"`
	Currency        string          `json:"currency"`
	Amount          int64           `json:"amount"`
	ConvertedAmount decimal.Decimal `json:"converted_amount"`
	Status          string          `json:"status"`
	Remarks         string          `json:"remarks"`
	FirstName       string          `json:"first_name"`
	LastName        string          `json:"last_name"`
	AccountNumber   string          `json:"account_number"`
	PhoneNumber     string          `json:"phone_number"`
}

type RequestDataList struct {
	Data     []RequestData `json:"data"`
	Metadata PageMetadata  `json:"metadata"`
}
