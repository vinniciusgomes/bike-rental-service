package repositories

import (
	"github.com/google/uuid"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdatePassword(userID uuid.UUID, newPassword string) error
	CreateValidationToken(token *models.ValidationToken) error
	GetValidationToken(token string) (*models.ValidationToken, error)
	DeleteValidationToken(token string) error
	ValidateUser(id string) error
}

type authRepositoryImp struct {
	db *gorm.DB
}

// NewAuthRepository creates a new instance of the AuthRepository interface.
//
// Parameters:
// - db: a pointer to a gorm.DB object representing the database connection.
//
// Returns:
// - AuthRepository: an implementation of the AuthRepository interface.
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImp{
		db: db,
	}
}

// CreateUser creates a new user in the database.
//
// Parameters:
// - user: a pointer to a User object representing the user to be created.
//
// Returns:
// - error: an error if there was a problem creating the user, or nil if the user was created successfully.
func (r *authRepositoryImp) CreateUser(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {

		return err
	}

	return nil
}

// GetUserByEmail retrieves a user from the database based on their email.
//
// Parameters:
// - email: the email of the user to retrieve.
//
// Returns:
// - *models.User: a pointer to the retrieved user model.
// - error: an error if the user retrieval fails.
func (r *authRepositoryImp) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID retrieves a user from the database based on their ID.
//
// Parameters:
// - id: the ID of the user to retrieve.
//
// Returns:
// - *models.User
func (r *authRepositoryImp) GetUserByID(id string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateValidationToken creates a validation token in the database.
//
// Parameters:
// - token: a pointer to a ValidationToken object representing the token to be created.
// Returns:
// - error: an error if there was a problem creating the token, or nil if the token was created successfully.
func (r *authRepositoryImp) CreateValidationToken(token *models.ValidationToken) error {
	if err := r.db.Create(token).Error; err != nil {
		return err
	}

	return nil
}

// GetValidationToken retrieves a validation token from the database based on the provided token.
//
// Parameters:
// - token: a string representing the token to search for.
// Returns:
// - *models.ValidationToken: a pointer to the retrieved validation token model.
// - error: an error if the token retrieval fails.
func (r *authRepositoryImp) GetValidationToken(token string) (*models.ValidationToken, error) {
	var ValidationToken models.ValidationToken

	if err := r.db.Where("token = ?", token).First(&ValidationToken).Error; err != nil {
		return nil, err
	}

	return &ValidationToken, nil
}

// UpdatePassword updates the password for a user in the database.
//
// Parameters:
// - userID: the UUID of the user whose password is being updated.
// - newPassword: the new password to set for the user.
//
// Returns:
// - error: an error if there was a problem updating the password, or nil if the password was updated successfully.
func (r *authRepositoryImp) UpdatePassword(userID uuid.UUID, newPassword string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("password", newPassword).Error
}

// DeleteValidationToken deletes a validation token from the database based on the provided token.
//
// Parameters:
// - token: a string representing the token to delete.
//
// Returns:
// - error: an error if there was a problem deleting the token, or nil if the token was deleted successfully.
func (r *authRepositoryImp) DeleteValidationToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&models.ValidationToken{}).Error
}

// ValidateUser updates the "validated" field of a user in the database.
//
// Parameters:
// - id: a string representing the ID of the user to validate.
//
// Returns:
// - error: an error if there was a problem updating the user, or nil if the user was updated successfully.
func (r *authRepositoryImp) ValidateUser(id string) error {
	return r.db.Where("id = ?", id).Model(&models.User{}).Update("verified", true).Error
}
