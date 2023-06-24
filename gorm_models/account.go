package gormmodels

type Account struct {
	Base
	AccountNumber string `json:"account_number" gorm:"unique; not_null"`
	Email         string `json:"email" gorm:"not_null"`
	Username      string `json:"username" gorm:"not_null"`
	FirstName     string `json:"first_name" gorm:"not_null"`
	LastName      string `json:"last_name"`
	Password      []byte `gorm:"not_null"`
	AccountType   string `gorm:"not_null"`
}