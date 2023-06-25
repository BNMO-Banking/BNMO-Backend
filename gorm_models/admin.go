package gormmodels

import (
	"BNMO/enum"

	"github.com/google/uuid"
)

type Admin struct {
	Base
	Role      enum.AccountRole `json:"role"`
	Account   Account          `json:"account"`
	AccountID uuid.UUID
}
