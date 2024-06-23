package repositories

import (
	"errors"
	"fmt"

	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"gorm.io/gorm"
)

type BikeRepository interface {
	CreateBike(bike *models.Bike) error
	GetAllBikes() (*[]models.Bike, error)
	GetBikeByID(id string) (*models.Bike, error)
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

func (r *bikeRepositoryImp) GetAllBikes() (*[]models.Bike, error) {
	var bikes []models.Bike

	err := r.db.Find(&bikes).Error
	if err != nil {
		return nil, err
	}

	return &bikes, nil
}

func (r *bikeRepositoryImp) GetBikeByID(id string) (*models.Bike, error) {
	var bike models.Bike

	if err := r.db.First(&bike, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("bike not found")
		}

		return nil, err
	}

	return &bike, nil
}
