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
	ErrorTypeValidationFailed  = &ErrorType{http.StatusBadRequest, "validation_failed", "Validation Failed"}
	ErrorTypeInvalidCredential = &ErrorType{http.StatusUnauthorized, "invalid_credential", "Invalid Credential"}
	ErrorTypeDataNotFound      = &ErrorType{http.StatusNotFound, "data_not_found", "Data Not Found"}
	ErrorTypeInternal          = &ErrorType{http.StatusInternalServerError, "internal", "Internal"}
)

var allErrorTypes = []*ErrorType{
	ErrorTypeValidationFailed,
	ErrorTypeInvalidCredential,
	ErrorTypeDataNotFound,
	ErrorTypeInternal,
}
