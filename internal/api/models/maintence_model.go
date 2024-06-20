package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaintenanceStatus string

const (
	PENDING    MaintenanceStatus = "pending"
	INPROGRESS MaintenanceStatus = "in_progress"
	COMPLETED  MaintenanceStatus = "completed"
)

type Maintenance struct {
	gorm.Model
	ID         uuid.UUID         `json:"id" gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4();index"`
	BikeID     uuid.UUID         `json:"bike_id" gorm:"type:uuid;not null" validate:"required,uuid4"`
	ReportedAt time.Time         `json:"reported_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	ResolvedAt *time.Time        `json:"resolved_at" gorm:"default:null"`
	Status     MaintenanceStatus `json:"status" gorm:"type:enum('pending', 'in_progress', 'completed');not null;default:'pending'" validate:"required,oneof='pending' 'in_progress' 'completed'"`
	Issue      string            `json:"issue" gorm:"type:text;not null" validate:"required"`
	ResolvedBy *uuid.UUID        `json:"resolved_by" gorm:"type:uuid;default:null"`
}
