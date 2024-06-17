package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BikeStatus string

const (
	AVAILABLE    BikeStatus = "available"
	NOTAVAILABLE BikeStatus = "notavailable"
	BOOKED       BikeStatus = "booked"
	MAINTENANCE  BikeStatus = "maintenance"
)

type BikeModel struct {
	gorm.Model
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4();index" validate:"required,uuid4"`
	Name         string     `json:"name" gorm:"not null;size:100;" validate:"required,min=1,max=100"`
	Description  string     `json:"description" gorm:"not null;size:500;" validate:"required,min=1,max=500"`
	PricePerHour float64    `json:"price_per_hour" gorm:"not null;" validate:"required"`
	Location     string     `json:"location" gorm:"not null;size:100;" validate:"required,min=1,max=100"`
	Status       BikeStatus `json:"status" gorm:"type:enum('available', 'notavailable', 'booked', 'maintenance');default:'available';not null" validate:"required,oneof='available' 'notavailable' 'booked' 'maintenance'"`
	Image        string     `json:"image" gorm:"not null;size:500;" validate:"required,min=1,max=500"`
}
