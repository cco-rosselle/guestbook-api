// a simple copy/paste from capex

package models

import ( 
	"net/http"
)

// APIError is the generic error struct used for returning
// error status responses on requests
type APIError struct {
	Message string `json:"message,omitempty"`
	Status int `json:"status"`
	Title string `json:"title"`
}

func (e *APIError) Error() string {
	return e.Message
}

// BadRequestError is an API error for 400s
func BadRequestError(m string) *APIError {
	return &APIError{
		Message: m,
		Status:  http.StatusBadRequest,
		Title:   "Bad Request",
	}
}

// Conflict is an API error for 409s
func Conflict(m string) *APIError {
	return &APIError{
		Message: m,
		Status:  http.StatusConflict,
		Title:   "Conflict error",
	}
}

// InternalServerError is an API error for 500s
func InternalServerError(m string) *APIError {
	return &APIError{
		Message: m,
		Status:  http.StatusInternalServerError,
		Title:   "Internal Server Error",
	}
}

// ResourceNotFoundError is an API error for 404s
func ResourceNotFoundError(m string) *APIError {
	return &APIError{
		Message: m,
		Status:  http.StatusNotFound,
		Title:   "Resource Not Found",
	}
}

// Teapot is an API error for 418s
func Teapot(m string) *APIError {
	return &APIError{
		Message: m,
		Status:  http.StatusTeapot,
		Title:   "I'm a teapot",
	}
}

// UnauthorizedError is an API error for 401s
func UnauthorizedError(m string) *APIError {
	return &APIError{
		Message: m,
		Status:  http.StatusUnauthorized,
		Title:   "Unauthorized",
	}
}