package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID             uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4();index" validate:"required,uuid4"`
	UserID         uuid.UUID      `json:"user_id" gorm:"type:uuid;not null"`
	CardNumber     string         `json:"card_number" gorm:"not null;size:16" validate:"required,min=16,max=16"`
	CardExpiry     string         `json:"card_expiry" gorm:"not null;size:5" validate:"required,min=5,max=5"`
	CardCVV        string         `json:"card_cvv" gorm:"not null;size:3" validate:"required,min=3,max=3"`
	BillingAddress string         `json:"billing_address" gorm:"not null;size:255" validate:"required,min=1,max=255"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
