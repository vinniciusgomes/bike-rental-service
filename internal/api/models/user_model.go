package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	UserStatusActive   = "active"
	UserStatusInactive = "inactive"

	UserRoleAdmin   = "admin"
	UserRoleDefault = "user"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Email     string         `json:"email" gorm:"unique;not null;size:100;" validate:"required,email"`
	Password  string         `json:"password" gorm:"not null;size:100;" validate:"required,min=1,max=100"`
	Name      string         `json:"name" gorm:"not null;size:100;" validate:"required,min=1,max=100"`
	Status    string         `json:"status" gorm:"not null;default:'active'" validate:"required,oneof='active' 'inactive'"`
	Role      string         `json:"role" gorm:"not null;default:'user'" validate:"required,oneof='admin' 'user'"`
	Image     string         `json:"image" gorm:"size:500;"`
	Verified  bool           `json:"verified" gorm:"not null;default:false"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
