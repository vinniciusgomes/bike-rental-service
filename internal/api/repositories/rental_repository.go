package repositories

import (
	"errors"
	"fmt"

	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"gorm.io/gorm"
)

type RentalRepository interface {
	CreateRental(rental *models.Rental) error
	GetBikeByID(id string) (*models.Bike, error)
	GetAllRentals() (*[]models.Rental, error)
	GetRentalByUserID(id string) (*[]models.Rental, error)
	GetRentalByID(id string) (*models.Rental, error)
	UpdateBikeStatus(bikeID string, status models.BikeStatusEnum) error
	UpdateRental(rental *models.Rental) error
}

type rentalRepositoryImp struct {
	db *gorm.DB
}

func NewRentalRepository(db *gorm.DB) RentalRepository {
	return &rentalRepositoryImp{
		db: db,
	}
}

func (r *rentalRepositoryImp) CreateRental(rental *models.Rental) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(rental).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&models.Bike{}).Where("id = ?", rental.BikeID).Update("status", models.BIKE_STATUS_BOOKED).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *rentalRepositoryImp) GetBikeByID(id string) (*models.Bike, error) {
	var bike models.Bike

	if err := r.db.Where("id = ?", id).First(&bike).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("bike not found")
		}
		return nil, err
	}

	return &bike, nil
}

func (r *rentalRepositoryImp) GetAllRentals() (*[]models.Rental, error) {
	var rentals []models.Rental

	err := r.db.Find(&rentals).Error
	if err != nil {
		return nil, err
	}

	return &rentals, nil
}

func (r *rentalRepositoryImp) GetRentalByUserID(id string) (*[]models.Rental, error) {
	var rentals []models.Rental

	err := r.db.Find(&rentals).Where("user_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &rentals, nil
}

func (r *rentalRepositoryImp) GetRentalByID(id string) (*models.Rental, error) {
	var rental models.Rental
	if err := r.db.Where("id = ?", id).First(&rental).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("rental not found")
		}
		return nil, err
	}
	return &rental, nil
}

func (r *rentalRepositoryImp) UpdateBikeStatus(bikeID string, status models.BikeStatusEnum) error {
	return r.db.Model(&models.Bike{}).Where("id = ?", bikeID).Update("status", status).Error
}

func (r *rentalRepositoryImp) UpdateRental(rental *models.Rental) error {
	return r.db.Save(rental).Error
}
