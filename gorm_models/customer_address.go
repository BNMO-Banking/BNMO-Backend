package gormmodels

type CustomerAddress struct {
	Base
	AddressLine1 string `gorm:"not null"`
	AddressLine2 string `gorm:"default:null"`
	City         string `gorm:"not null"`
	State        string `gorm:"not null"`
	PostalCode   string `gorm:"not null"`
	Country      string `gorm:"not null"`
}
