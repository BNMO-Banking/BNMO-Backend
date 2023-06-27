package models

import (
	"BNMO/enum"

	"github.com/google/uuid"
)

type RequestReq struct {
	Id          uuid.UUID    `json:"id"`
	RequestType enum.ReqType `json:"request_type"`
	Currency    string       `json:"currency"`
	Amount      int64        `json:"amount"`
	Pin         string       `json:"pin"`
}
