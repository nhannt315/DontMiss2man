package errors

import (
	"fmt"
	"io"
	"strings"

	"github.com/nhannt315/real_estate_api/pkg/errors"
)

// New エラー作成
func New(errorType *ErrorType, detailMessage string, details ...APIErrorDetails) error {
	return newError(nil, errorType, detailMessage, details...)
}

// Errorf エラー作成
func Errorf(errorType *ErrorType, detailFormat string, args ...interface{}) error {
	return newError(nil, errorType, fmt.Sprintf(detailFormat, args...))
}

// Wrap 引数のエラーをWrapして新しいエラーを作成
func Wrap(err error, errorType *ErrorType, detailMessage string) error {
	return newError(err, errorType, detailMessage)
}

// Wrapf 引数のエラーをWrapして新しいエラーを作成
func Wrapf(err error, errorType *ErrorType, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return newError(err, errorType, msg)
}

// Internal インターナルエラーを作成
func Internal(err error) error {
	return newError(err, ErrorTypeInternal, "the server encountered an internal server error")
}

// NewErrorDetail 詳細エラーを作成
func NewErrorDetail(resource, name, reason string) APIErrorDetails {
	return &apiErrorDetail{
		resource: resource,
		name:     name,
		reason:   reason,
	}
}

// NewError APIエラーを作成
func newError(err error, errorType *ErrorType, detailMessage string, details ...APIErrorDetails) error {
	apiError := &apiError{
		errorType:     errorType,
		detailMessage: detailMessage,
		details:       details,
	}

	if err == nil {
		err = errors.New(apiError.errorMessage())
	} else {
		err = errors.Wrap(err, apiError.errorMessage())
	}

	apiError.err = err

	return apiError
}

// APIError APIエラー
type APIError interface {
	error
	// エラー種別
	ErrorType() *ErrorType
	DetailMessage() string
	Details() []APIErrorDetails
}

type apiError struct {
	errorType     *ErrorType
	detailMessage string

	details []APIErrorDetails
	err     error
}

func (a *apiError) errorMessage() string {
	msg := fmt.Sprintf("status: {http: %d}, type: %s, detailMesage: %s",
		a.errorType.HTTPStatusCode(),
		a.errorType.Type(), a.DetailMessage())
	if len(a.details) > 0 {

		var sb strings.Builder
		sb.WriteString(msg)
		sb.WriteString(" details: [")
		for i, d := range a.details {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(d.String())
		}
		sb.WriteString("]")

		msg = sb.String()
	}

	return msg
}

func (a *apiError) ErrorType() *ErrorType {
	return a.errorType
}

func (a *apiError) DetailMessage() string {
	return a.detailMessage
}

func (a *apiError) Details() []APIErrorDetails {
	return a.details
}

func (a *apiError) Error() string {
	return a.err.Error()
}

// Unwrap implements error.Unwrap
func (a *apiError) Unwrap() error {
	if a == nil {
		return nil
	}
	return errors.Unwrap(a.err)
}

// nolint: errcheck
func (a *apiError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", a.err)
			return
		}
		fmt.Fprintf(s, "%v", a.err)

	default:
		io.WriteString(s, a.Error())
	}
}

// APIErrorDetails 詳細エラー
type APIErrorDetails interface {
	fmt.Stringer
	Resource() string
	Name() string
	Reason() string
}

type apiErrorDetail struct {
	resource string
	name     string
	reason   string
}

func (a *apiErrorDetail) String() string {
	return fmt.Sprintf("{resource: %s, name: %s, reason: %s}", a.resource, a.name, a.reason)
}

func (a *apiErrorDetail) Resource() string {
	return a.resource
}

func (a *apiErrorDetail) Name() string {
	return a.name
}

func (a *apiErrorDetail) Reason() string {
	return a.reason
}
