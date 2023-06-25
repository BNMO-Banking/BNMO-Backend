package gormmodels

import (
	"BNMO/enum"

	"github.com/google/uuid"
)

type Admin struct {
	Base
	Role      enum.AccountRole
	Account   Account
	AccountID uuid.UUID
}
