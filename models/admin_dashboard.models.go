package models

import (
	"BNMO/enum"
	"time"

	"github.com/shopspring/decimal"
)

type PendingAccount struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

type PendingRequest struct {
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	RequestType enum.ReqType `json:"request_type"`
	CreatedAt   time.Time    `json:"created_at"`
}

type NewAccountStatistics struct {
	TotalAccounts   int64   `json:"total_accounts"`
	YearlyAccounts  int64   `json:"yearly_accounts"`
	MonthlyAccounts []int64 `json:"monthly_accounts"`
}

type BankStatistics struct {
	TotalExpense   decimal.Decimal   `json:"total_expense"`
	TotalIncome    decimal.Decimal   `json:"total_income"`
	MonthlyExpense []decimal.Decimal `json:"monthly_expense"`
	MonthlyIncome  []decimal.Decimal `json:"monthly_income"`
}
