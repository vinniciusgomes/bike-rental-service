package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RentalStatusEnum string

const (
	RENTAL_STATUS_ACTIVE    RentalStatusEnum = "active"
	RENTAL_STATUS_COMPLETED RentalStatusEnum = "completed"
)

type Rental struct {
	ID        uuid.UUID        `json:"id" gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4();index" validate:"required,uuid4"`
	UserID    uuid.UUID        `json:"user_id" gorm:"type:uuid;not null"`
	BikeID    uuid.UUID        `json:"bike_id" gorm:"type:uuid;not null"`
	StartTime time.Time        `json:"start_time" gorm:"not null"`
	EndTime   time.Time        `json:"end_time" gorm:"not null"`
	Status    RentalStatusEnum `json:"status" gorm:"not null;default:'active'" validate:"required,oneof='active' 'completed'"`
	TotalCost float64          `json:"total_cost" gorm:"not null"`
	CreatedAt time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at" gorm:"index"`
}
