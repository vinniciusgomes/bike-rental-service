package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MaintenanceStatusEnum represents the status of maintenance.
type MaintenanceStatusEnum string

const (
	PENDING    MaintenanceStatusEnum = "pending"
	INPROGRESS MaintenanceStatusEnum = "in_progress"
	COMPLETED  MaintenanceStatusEnum = "completed"
)

type Maintenance struct {
	ID         uuid.UUID             `json:"id" gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4();index"`
	BikeID     uuid.UUID             `json:"bike_id" gorm:"type:uuid;not null" validate:"required,uuid4"`
	ReportedAt time.Time             `json:"reported_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	ResolvedAt *time.Time            `json:"resolved_at" gorm:"default:null"`
	Status     MaintenanceStatusEnum `json:"status" gorm:"type:enum('pending', 'in_progress', 'completed');not null;default:'pending'" validate:"required,oneof='pending' 'in_progress' 'completed'"`
	Issue      string                `json:"issue" gorm:"type:text;not null" validate:"required"`
	ResolvedBy *uuid.UUID            `json:"resolved_by" gorm:"type:uuid;default:null"`
	CreatedAt  time.Time             `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time             `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt        `json:"deleted_at" gorm:"index"`
}
