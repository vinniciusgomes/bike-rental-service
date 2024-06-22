package repositories

import (
	"fmt"

	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(filters map[string]interface{}) (*[]models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type userRepositoryImp struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of the UserRepository interface.
//
// Parameters:
// - db: a pointer to a gorm.DB object representing the database connection.
//
// Returns:
// - UserRepository: an implementation of the UserRepository interface.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImp{db: db}
}

// GetUserByID retrieves a user from the database based on their ID.
//
// Parameters:
// - id: the ID of the user to retrieve.
//
// Returns:
// - *models.User
func (r *userRepositoryImp) GetUserByID(id string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAllUsers retrieves all users from the database based on the provided filters.
//
// Parameters:
//   - filters: a map of filters to apply to the query. The keys can be "email" or "name" to perform a case-insensitive
//     search, or any other key to perform an exact match. The values are the filter values.
//
// Returns:
// - *[]models.User
func (r *userRepositoryImp) GetAllUsers(filters map[string]interface{}) (*[]models.User, error) {
	var users []models.User

	query := r.db.Model(&models.User{})
	for key, value := range filters {
		if key == "email" || key == "name" {
			query = query.Where(fmt.Sprintf("%s ILIKE ?", key), fmt.Sprintf("%%%s%%", value))
		} else {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

// UpdateUser updates a user in the database.
//
// Parameters:
// - user: a pointer to a models.User struct representing the user to be updated.
//
// Returns:
// - error: an error if the update operation fails, or if the user is not found.
func (r *userRepositoryImp) UpdateUser(user *models.User) error {
	result := r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// DeleteUser deletes a user from the database.
//
// Parameters:
// - user: a pointer to a models.User struct representing the user to be deleted.
//
// Returns:
// - error: an error if the delete operation fails.
func (r *userRepositoryImp) DeleteUser(user *models.User) error {
	return r.db.Where("id = ?", user.ID).Delete(&models.User{}).Error
}
