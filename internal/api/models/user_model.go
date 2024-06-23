package models

import (
	"time"

	"github.com/google/uuid"
)

// UserStatusEnum represents the status of a user.
type UserStatusEnum string

const (
	UserStatusActive   UserStatusEnum = "active"
	UserStatusInactive UserStatusEnum = "inactive"
)

// UserRoleEnum represents the role of a user.
type UserRoleEnum string

const (
	UserRoleAdmin   UserRoleEnum = "admin"
	UserRoleDefault UserRoleEnum = "user"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Email     string         `json:"email" gorm:"unique;not null;size:100;" validate:"required,email"`
	Password  string         `json:"password" gorm:"not null;size:100;" validate:"required,min=1,max=100"`
	Name      string         `json:"name" gorm:"not null;size:100;" validate:"required,min=1,max=100"`
	Phone     string         `json:"phone" gorm:"not null;size:100;"`
	Status    UserStatusEnum `json:"status" gorm:"not null;default:'active'" validate:"required,oneof='active' 'inactive'"`
	Role      UserRoleEnum   `json:"role" gorm:"not null;default:'user'" validate:"required,oneof='admin' 'user'"`
	Image     string         `json:"image" gorm:"size:500;"`
	Verified  bool           `json:"verified" gorm:"not null;default:false"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}
