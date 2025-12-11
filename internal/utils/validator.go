package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ValidationError represents a single validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// GetValidationErrors extracts validation errors from gin binding error
func GetValidationErrors(err error) []ValidationError {
	var errors []ValidationError

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, ValidationError{
				Field:   strings.ToLower(e.Field()),
				Message: getErrorMessage(e),
			})
		}
	}

	return errors
}

// getErrorMessage returns user-friendly error message
func getErrorMessage(e validator.FieldError) string {
	field := strings.ToLower(e.Field())

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", field, e.Param())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must not exceed %s characters", field, e.Param())
	case "numeric":
		return fmt.Sprintf("%s must contain only numbers", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

// HandleValidationError handles validation errors in handler
func HandleValidationError(c *gin.Context, err error) {
	errors := GetValidationErrors(err)

	if len(errors) > 0 {
		// Return first error message (most common approach)
		c.JSON(400, gin.H{
			"status":  0,
			"message": errors[0].Message,
			"data":    nil,
		})
		return
	}

	// Fallback for unknown errors
	c.JSON(400, gin.H{
		"status":  0,
		"message": "Invalid request body",
		"data":    nil,
	})
}
