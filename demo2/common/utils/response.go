package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is a standardized API response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse returns a successful API response
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// ErrorResponse returns an error API response
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Message: message,
	})
}

// BadRequestResponse returns a 400 Bad Request response
func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message)
}

// UnauthorizedResponse returns a 401 Unauthorized response
func UnauthorizedResponse(c *gin.Context, message string) {
	if message == "" {
		message = "unauthorized"
	}
	ErrorResponse(c, http.StatusUnauthorized, message)
}

// ForbiddenResponse returns a 403 Forbidden response
func ForbiddenResponse(c *gin.Context, message string) {
	if message == "" {
		message = "access forbidden"
	}
	ErrorResponse(c, http.StatusForbidden, message)
}

// NotFoundResponse returns a 404 Not Found response
func NotFoundResponse(c *gin.Context, message string) {
	if message == "" {
		message = "resource not found"
	}
	ErrorResponse(c, http.StatusNotFound, message)
}

// InternalServerErrorResponse returns a 500 Internal Server Error response
func InternalServerErrorResponse(c *gin.Context, message string) {
	if message == "" {
		message = "internal server error"
	}
	ErrorResponse(c, http.StatusInternalServerError, message)
} 