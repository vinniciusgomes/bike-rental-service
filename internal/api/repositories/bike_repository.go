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
	UpdateBike(bike *models.Bike) error
	DeleteBike(id string) error
}

type bikeRepositoryImp struct {
	db *gorm.DB
}

func NewBikeRepository(db *gorm.DB) BikeRepository {
	return &bikeRepositoryImp{
		db: db,
	}
}

// CreateBike creates a new bike in the database.
//
// Parameters:
// - bike: a pointer to a models.Bike object representing the bike to be created.
//
// Returns:
// - error: an error if there was a problem creating the bike, or nil if the bike was created successfully.
func (r *bikeRepositoryImp) CreateBike(bike *models.Bike) error {
	if err := r.db.Create(bike).Error; err != nil {
		return err
	}

	return nil
}

// GetAllBikes retrieves all bikes from the database.
//
// It returns a pointer to a slice of models.Bike and an error if any.
func (r *bikeRepositoryImp) GetAllBikes() (*[]models.Bike, error) {
	var bikes []models.Bike

	err := r.db.Find(&bikes).Error
	if err != nil {
		return nil, err
	}

	return &bikes, nil
}

// GetBikeByID retrieves a bike from the database by its ID.
//
// Parameters:
// - id: the ID of the bike to retrieve.
//
// Returns:
// - *models.Bike: a pointer to the bike model if found, or nil if not found.
// - error: an error if there was a problem retrieving the bike, or nil if successful.
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

// DeleteBike deletes a bike from the database based on its ID.
//
// Parameters:
// - id: the ID of the bike to be deleted.
//
// Returns:
// - error: an error if the bike is not found or if there is an error deleting the bike.
func (r *bikeRepositoryImp) DeleteBike(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Bike{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("bike not found")
		}
		return err
	}
	return nil
}

// UpdateBike updates a bike in the database.
//
// Parameters:
// - bike: a pointer to a models.Bike object representing the bike to be updated.
//
// Returns:
// - error: an error if there was a problem updating the bike, or nil if the bike was updated successfully.
func (r *bikeRepositoryImp) UpdateBike(bike *models.Bike) error {
	result := r.db.Model(&models.Bike{}).Omit("ID", "CreatedAt").Where("id = ?", bike.ID).Updates(bike)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("bike not found")
	}

	return nil
}
