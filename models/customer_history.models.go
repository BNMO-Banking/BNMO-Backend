package models

import (
	"BNMO/enum"
	"time"

	"github.com/shopspring/decimal"
)

type RequestHistory struct {
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	RequestType     enum.ReqType    `json:"request_type"`
	Currency        string          `json:"currency"`
	Amount          int64           `json:"amount"`
	ConvertedAmount decimal.Decimal `json:"converted_amount"`
	Status          string          `json:"status"`
	Remarks         string          `json:"remarks"`
}

type RequestHistoryList struct {
	Data     []RequestHistory `json:"data"`
	Metadata PageMetadata     `json:"metadata"`
}

type TransferHistory struct {
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	Currency        string          `json:"currency"`
	Amount          int64           `json:"amount"`
	ConvertedAmount decimal.Decimal `json:"converted_amount"`
	Status          string          `json:"status"`
	Description     string          `json:"description"`
	AccountNumber   string          `json:"account_number"`
	FirstName       string          `json:"first_name"`
	LastName        string          `json:"last_name"`
}

type TransferTypes struct {
	InboundTransfers  []TransferHistory `json:"inbound_list"`
	OutboundTransfers []TransferHistory `json:"outbound_list"`
}

type TransferHistoryList struct {
	Data     TransferTypes `json:"data"`
	Metadata PageMetadata  `json:"metadata"`
}
