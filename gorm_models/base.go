package gormmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID 			uuid.UUID 	`gorm:"type:uuid;primary_key;"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql:"index"`
}

func (base *Base) BeforeCreate(db *gorm.DB) {
	uuid := uuid.New()
	db.Statement.SetColumn("ID", uuid)
}