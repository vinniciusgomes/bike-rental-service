package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BikeStatusEnum represents the status of a bike.
type BikeStatusEnum string

const (
	BIKE_STATUS_AVAILABLE    BikeStatusEnum = "available"
	BIKE_STATUS_NOTAVAILABLE BikeStatusEnum = "notavailable"
	BIKE_STATUS_BOOKED       BikeStatusEnum = "booked"
	BIKE_STATUS_MAINTENANCE  BikeStatusEnum = "maintenance"
)

type Bike struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;not null;index"`
	Name         string         `json:"name" gorm:"not null;size:100;" validate:"required,min=1,max=100"`
	Description  string         `json:"description" gorm:"not null;size:500;" validate:"required,min=1,max=500"`
	PricePerHour float64        `json:"price_per_hour" gorm:"not null;" validate:"required"`
	Location     string         `json:"location" gorm:"not null;size:100;" validate:"required,min=1,max=100"`
	Status       BikeStatusEnum `json:"status" gorm:"not null" validate:"required,oneof='available' 'notavailable' 'booked' 'maintenance'"`
	Image        string         `json:"image" gorm:"not null;size:500;" validate:"required,min=1,max=500"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
