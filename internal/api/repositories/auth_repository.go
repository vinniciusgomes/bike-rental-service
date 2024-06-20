package repositories

import (
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
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
