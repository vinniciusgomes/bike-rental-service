package repositories

import (
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"gorm.io/gorm"
)

type BikeRepository interface {
	CreateBike(bike *models.Bike) error
}

type bikeRepositoryImp struct {
	db *gorm.DB
}

func NewBikeRepository(db *gorm.DB) BikeRepository {
	return &bikeRepositoryImp{
		db: db,
	}
}

func (r *bikeRepositoryImp) CreateBike(bike *models.Bike) error {
	if err := r.db.Create(bike).Error; err != nil {
		return err
	}

	return nil
}
