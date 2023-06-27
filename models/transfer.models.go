package models

import "github.com/google/uuid"

type TransferReq struct {
	Id                uuid.UUID `json:"id"`
	DestinationNumber string    `json:"destination_number"`
	Currency          string    `json:"currency"`
	Amount            int64     `json:"amount"`
	Description       string    `json:"description"`
	Pin               string    `json:"pin"`
}
