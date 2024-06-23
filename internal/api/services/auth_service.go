package services

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/constants"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/helpers"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/repositories"
	"github.com/vinniciusgomes/ebike-rental-service/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	user := new(models.User)

	id, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to create user"})
		return
	}

	if err := c.BindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to create user"})
		return
	}

	user.ID = id
	user.Role = models.UserRoleDefault
	user.Status = models.UserStatusInactive
	user.Verified = false

	if err = helpers.ValidateModel(user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to create user"})
		return
	}

	user.Password = string(hash)

	if err = s.repo.CreateUser(user); err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
			c.JSON(http.StatusConflict, gin.H{"message": "user with email already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to create user"})
		return
	}

	// TODO: send email with token and save token in db

	c.Status(http.StatusCreated)
}

// TODO: add verify user
func (s *AuthService) VerifyUser(c *gin.Context) {
	// Get token and email from body

	// Verify token and email

	// Validate token

	// Update user for validate: true in db

	// Return success
	c.JSON(http.StatusOK, gin.H{"message": "verify user"})
}

// Login handles the login functionality for the AuthService.
//
// Parameters:
// - c: a pointer to the gin.Context object for handling HTTP request and response.
//
// Returns: void
func (s *AuthService) Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid email or password"})
		return
	}

	if err := helpers.ValidateModel(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	user, err := s.repo.GetUserByEmail(body.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to log in"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to log in"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(constants.AuthCookieName, tokenString, 3600*24*7, "", "", false, true)

	c.Status(http.StatusOK)
}

// TODO: add refresh token
func (s *AuthService) RefreshToken(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Logout handles the logout functionality for the AuthService.
//
// Parameters:
// - c: a pointer to the gin.Context object for handling HTTP request and response.
//
// Returns: void
func (s *AuthService) Logout(c *gin.Context) {
	c.SetCookie(constants.AuthCookieName, "", -1, "", "", false, true)
	c.Status(http.StatusOK)
}

// ForgotPassword handles the forgot password functionality for the AuthService.
//
// Parameters:
// - c: a pointer to the gin.Context object for handling HTTP request and response.
//
// Returns: void
func (s *AuthService) ForgotPassword(c *gin.Context) {
	var body struct {
		Email string `json:"email" validate:"required,email"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid email"})
		return
	}

	if err := helpers.ValidateModel(body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	user, err := s.repo.GetUserByEmail(body.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "email sent successfully",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to reset password"})
		return
	}

	tokenString, err := helpers.GenerateSecureToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to reset password"})
		return
	}

	token := models.ValidationToken{
		Token:     tokenString,
		Type:      models.ForgotPasswordToken, // Usando o enum correto
		Valid:     true,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	if err := s.repo.CreateValidationToken(&token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to reset password"})
		return
	}

	resetURL := fmt.Sprintf("%s/reset-password/%s", os.Getenv("WEB_CLIENT_URL"), tokenString)
	err = pkg.SendEmail([]string{user.Email}, "Reset Password", fmt.Sprintf("Click the link to reset your password: <a href='%s' target='_blank'>Reset Password</a>", resetURL))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to reset password"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "email sent",
	})
}

// ResetPassword handles the password reset functionality for the AuthService.
//
// Parameters:
// - c: a pointer to the gin.Context object for handling HTTP request and response.
// Returns: void
func (s *AuthService) ResetPassword(c *gin.Context) {
	tokenString := c.Param("token")

	var body struct {
		Password string `json:"password" gorm:"not null;size:100" validate:"required,min=1,max=100"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid password"})
		return
	}

	if err := helpers.ValidateModel(body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	token, err := s.repo.GetValidationToken(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid or expired token"})
		return
	}

	if !token.Valid || time.Now().After(token.ExpiresAt) {
		token.Valid = false
		if err := s.repo.DeleteValidationToken(token.Token); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to reset password"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid or expired token"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to reset password"})
		return
	}

	token.Valid = false
	if err := s.repo.DeleteValidationToken(token.Token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to reset password"})
		return
	}

	if err := s.repo.UpdatePassword(token.UserID, string(hashedPassword)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to reset password"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "password reset successful",
	})
}
