package models

import (
	"time"

	"github.com/google/uuid"
)

// TokenTypeEnum represents the type of token.
type TokenTypeEnum string

const (
	AccessToken            TokenTypeEnum = "access_token"
	RefreshToken           TokenTypeEnum = "refresh_token"
	ForgotPasswordToken    TokenTypeEnum = "forgot_password_token"
	ValidationAccountToken TokenTypeEnum = "validation_account_token"
)

type ValidationToken struct {
	Token     string        `json:"token" gorm:"not null;size:100;index;unique" validate:"required,min=1,max=100"`
	Type      TokenTypeEnum `json:"type" gorm:"not null;size:100" validate:"required,min=1,max=100"`
	Valid     bool          `json:"valid" gorm:"not null;default:true"`
	UserID    uuid.UUID     `json:"user_id" gorm:"not null;type:uuid;index"`
	ExpiresAt time.Time     `json:"expires_at" gorm:"not null"`
	CreatedAt time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}
