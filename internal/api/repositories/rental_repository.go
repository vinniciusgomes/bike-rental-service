package repositories

import (
	"errors"
	"fmt"

	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"github.com/vinniciusgomes/ebike-rental-service/pkg"
	"gorm.io/gorm"
)

type RentalRepository interface {
	CreateRental(rental *models.Rental) error
	GetBikeByID(id string) (*models.Bike, error)
	GetAllRentals(pagination pkg.Pagination) (*[]models.Rental, *pkg.Pagination, error)
	GetRentalByUserID(id string) (*[]models.Rental, error)
	GetRentalByID(id string) (*models.Rental, error)
	UpdateBikeStatus(bikeID string, status models.BikeStatusEnum) error
	UpdateRental(rental *models.Rental) error
}

type rentalRepositoryImp struct {
	db *gorm.DB
}

// NewRentalRepository creates a new rental repository instance.
//
// Parameters:
// - db: a pointer to a gorm.DB object representing the database connection.
// Returns:
// - RentalRepository: an implementation of the RentalRepository interface.
func NewRentalRepository(db *gorm.DB) RentalRepository {
	return &rentalRepositoryImp{
		db: db,
	}
}

// CreateRental creates a new rental in the database.
//
// It takes a pointer to a models.Rental struct as a parameter, which represents the rental to be created.
// The function returns an error if there was an issue creating the rental or updating the bike's status.
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

// GetBikeByID retrieves a bike from the database by its ID.
//
// Parameters:
// - id: the ID of the bike to retrieve.
//
// Returns:
// - *models.Bike: a pointer to the bike model if found, or nil if not found.
// - error: an error if there was a problem retrieving the bike, or nil if successful.
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

// GetAllRentals retrieves all rentals from the rental repository.
//
// It takes a pagination parameter of type pkg.Pagination and returns a pointer to a slice of models.Rental, a pointer to the pagination parameter, and an error if any.
func (r *rentalRepositoryImp) GetAllRentals(pagination pkg.Pagination) (*[]models.Rental, *pkg.Pagination, error) {
	var rentals []models.Rental

	err := r.db.Scopes(pkg.Paginate(&models.Rental{}, &pagination, r.db)).Find(&rentals).Error
	if err != nil {
		return nil, nil, err
	}

	return &rentals, &pagination, nil
}

// GetRentalByUserID retrieves all rentals associated with a specific user ID.
//
// Parameters:
// - id: the ID of the user.
//
// Returns:
// - *[]models.Rental: a pointer to a slice of models.Rental containing the rentals associated with the user.
// - error: an error if there was a problem retrieving the rentals.
func (r *rentalRepositoryImp) GetRentalByUserID(id string) (*[]models.Rental, error) {
	var rentals []models.Rental

	err := r.db.Find(&rentals).Where("user_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &rentals, nil
}

// GetRentalByID retrieves a rental from the database by its ID.
//
// Parameters:
// - id: the ID of the rental to retrieve.
// Returns:
// - *models.Rental: a pointer to the rental model if found, or nil if not found.
// - error: an error if there was a problem retrieving the rental, or nil if successful.
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

// UpdateBikeStatus updates the status of a bike in the database.
//
// Parameters:
// - bikeID: the ID of the bike to update.
// - status: the new status to set for the bike.
//
// Returns:
// - error: an error if there was a problem updating the bike's status.
func (r *rentalRepositoryImp) UpdateBikeStatus(bikeID string, status models.BikeStatusEnum) error {
	return r.db.Model(&models.Bike{}).Where("id = ?", bikeID).Update("status", status).Error
}

// UpdateRental updates a rental in the repository.
//
// It takes a pointer to a models.Rental struct as a parameter, which represents the rental to be updated.
// The function returns an error if there was an issue updating the rental.
func (r *rentalRepositoryImp) UpdateRental(rental *models.Rental) error {
	return r.db.Save(rental).Error
}
