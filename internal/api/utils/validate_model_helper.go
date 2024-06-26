package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateModel validates the given struct object using the validator package.
//
// It takes an interface{} as a parameter, which represents the object to be validated.
// The function returns an error if the validation fails, otherwise it returns nil.
// If the validation fails, the error message is constructed based on the validation error tag and the struct field.
// The error message includes the name of the field and the specific validation requirement that was not met.
func ValidateModel(obj interface{}) error {
	// Create a new instance of the validator.
	validate := validator.New()

	// Validate the provided struct.
	err := validate.Struct(obj)

	// If there are no validation errors, return nil.
	if err == nil {
		return nil
	}

	// Convert the validation errors to the type validator.ValidationErrors.
	validationErrors := err.(validator.ValidationErrors)

	// Get the first validation error.
	validationError := validationErrors[0]

	// Convert the field name to lowercase.
	field := strings.ToLower(validationError.StructField())

	// Check the type of validation that failed and return an appropriate error message.
	switch validationError.Tag() {
	case "required":
		// Case when the field is required but not provided.
		return errors.New(field + " is required")
	case "url":
		// Case when the field should be a valid URL but is not.
		return errors.New(field + " is an invalid URL")
	case "max":
		// Case when the field value should be less than or equal to a specific value.
		return errors.New(field + " must be less than or equal to " + validationError.Param())
	case "min":
		// Case when the field value should be greater than or equal to a specific value.
		return errors.New(field + " must be greater than or equal to " + validationError.Param())
	case "email":
		// Case when the field should be a valid email address but is not.
		return errors.New(field + " is an invalid email")
	case "uuid4":
		// Case when the field should be a valid UUID version 4 but is not.
		return errors.New(field + " must be a valid UUIDv4")
	case "oneof":
		// Case when the field value should be one of the specified values.
		return errors.New(field + " must be one of: " + validationError.Param())
	case "alphanum":
		// Case when the field should contain only alphanumeric characters.
		return errors.New(field + " must contain only alphanumeric characters")
	case "len":
		// Case when the field should have a specific length.
		return errors.New(field + " must be exactly " + validationError.Param() + " characters long")
	case "numeric":
		// Case when the field should contain only numeric characters.
		return errors.New(field + " must contain only numeric characters")
	case "startswith":
		// Case when the field should start with a specific substring.
		return errors.New(field + " must start with " + validationError.Param())
	case "endswith":
		// Case when the field should end with a specific substring.
		return errors.New(field + " must end with " + validationError.Param())
	case "datetime":
		// Case when the field should be a valid datetime in a specific format.
		return errors.New(field + " must be a valid datetime in the format " + validationError.Param())
	default:
		// Generic case for any other validation errors.
		return errors.New("Validation error for field: " + field)
	}
}
