package services

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/domain/models"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/domain/repositories"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/helpers"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repositories.AuthRepository
}

// NewAuthService creates a new instance of the AuthService struct.
//
// Parameters:
// - repo: an instance of the AuthRepository interface representing the authentication repository.
//
// Returns:
// - *AuthService: a pointer to the newly created AuthService struct.
func NewAuthService(repo repositories.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser creates a new user based on the information provided in the request body.
//
// Parameters:
// - c: a pointer to the gin.Context object for handling HTTP request and response.
// Return type: void
func (s *AuthService) CreateUser(c *gin.Context) {
	user := new(models.UserModel)

	id, err := uuid.NewRandom()
	if err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	if err := c.BindJSON(user); err != nil {
		helpers.HandleError(c, err, http.StatusBadRequest)
		return
	}

	user.ID = id
	user.Role = models.UserRoleDefault
	user.Status = models.UserStatusInactive
	user.Verified = false

	if err = helpers.ValidateModel(user); err != nil {
		helpers.HandleError(c, err, http.StatusUnprocessableEntity)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	user.Password = string(hash)

	if err = s.repo.CreateUser(user); err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_user_models_email\" (SQLSTATE 23505)" {
			helpers.HandleError(c, errors.New("user with email already exists"), http.StatusConflict)
			return
		}

		helpers.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}
