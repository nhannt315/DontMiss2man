package errors

import "net/http"

type ErrorType struct {
	httpStatusCode int
	errType        string
	title          string
}

func (t *ErrorType) HTTPStatusCode() int {
	return t.httpStatusCode
}

func (t *ErrorType) Type() string {
	return t.errType
}

func (t *ErrorType) Title() string {
	return t.title
}

var (
	// ErrorTypeValidationFailed is the error type for validation errors.
	ErrorTypeValidationFailed = &ErrorType{http.StatusBadRequest, "validation_failed", "Validation Failed"}
	// ErrorTypeInvalidCredential is the error type for invalid credentials.
	ErrorTypeInvalidCredential = &ErrorType{http.StatusUnauthorized, "invalid_credential", "Invalid Credential"}
	// ErrorTypeDataNotFound is the error type for data not found.
	ErrorTypeDataNotFound = &ErrorType{http.StatusNotFound, "data_not_found", "Data Not Found"}
	// ErrorTypeInternal is the error type for internal errors.
	ErrorTypeInternal = &ErrorType{http.StatusInternalServerError, "internal", "Internal"}
)
