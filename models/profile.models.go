package models

import (
	"BNMO/enum"
	"mime/multipart"

	"github.com/shopspring/decimal"
)

type ProfileRes struct {
	AccountNumber      string           `json:"account_number"`
	AccountType        enum.AccountType `json:"account_type"`
	Email              string           `json:"email"`
	Username           string           `json:"username"`
	FirstName          string           `json:"first_name"`
	LastName           string           `json:"last_name"`
	CardNumber         string           `json:"card_number"`
	Balance            decimal.Decimal  `json:"balance"`
	PhoneNumber        string           `json:"phone_number"`
	ProfilePicturePath string           `json:"profile_pic_path"`
	AddressLine1       string           `json:"address_line_1"`
	AddressLine2       string           `json:"address_line_2"`
	City               string           `json:"city"`
	State              string           `json:"state"`
	PostalCode         string           `json:"postal_code"`
	Country            string           `json:"country"`
}

type EditProfileReq struct {
	PhoneNumber    string                `form:"phone_number"`
	ProfilePicture *multipart.FileHeader `form:"profile_pic"`
	AddressLine1   string                `form:"address_line_1"`
	AddressLine2   string                `form:"address_line_2"`
	City           string                `form:"city"`
	State          string                `form:"state"`
	PostalCode     string                `form:"postal_code"`
	Country        string                `form:"country"`
}

type ProfileStatistics struct {
	Balance         decimal.Decimal   `json:"balance"`
	TotalReceived   decimal.Decimal   `json:"total_received"`
	TotalSpent      decimal.Decimal   `json:"total_spent"`
	MonthlyReceived []decimal.Decimal `json:"monthly_received"`
	MonthlySpending []decimal.Decimal `json:"monthly_spending"`
}
