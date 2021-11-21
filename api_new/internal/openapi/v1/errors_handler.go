package openapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nhannt315/real_estate_api/internal/errors"
	apperrors "github.com/nhannt315/real_estate_api/internal/errors"
	"github.com/nhannt315/real_estate_api/internal/openapi"
	pkgerrors "github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

var openAPIErrorInternal = &Error{
	Title:  apperrors.ErrorTypeInternal.Title(),
	Type:   ErrorType(apperrors.ErrorTypeInternal.Type()),
	Detail: "the server encountered an internal server error",
}

// NewErrorHandler エラーハンドラを作成
// nolint:gocognit
func NewErrorHandler(logger logs.Logger) openapi.ErrorHandler {
	return func(err error, ectx echo.Context) (processed bool) {
		if !IsV1Request(ectx.Request()) {
			return false
		}
		// nolint:errorlint
		if _, ok := err.(*echo.HTTPError); ok {
			// request path 不正(404 Not Foundなど)の echo のエラーは、そのまま
			return false
		}

		ctx := ectx.Request().Context()
		statusCode := http.StatusInternalServerError
		e := openAPIErrorInternal

		// nolint:errorlint
		if a, ok := err.(errors.APIError); ok {
			e = newOpenAPIError(a)
			statusCode = a.ErrorType().HTTPStatusCode()
		}
		if ectx.Response().Committed {
			// 既に response が書き込み済みなら二重に書き込まない
			if ectx.Response().Status != statusCode {
				// 念のため、http statusが異なる場合、ログ記録
				logger.Error(ctx, pkgerrors.Wrapf(err, "no write error response. status=%d error=%+v",
					statusCode, e))
			}
			return false
		}
		if err := ectx.JSON(statusCode, e); err != nil {
			logger.Error(ctx, pkgerrors.Wrap(err, "fail to write error response"))
		}

		return true
	}
}

func newOpenAPIError(e errors.APIError) *Error {
	return &Error{
		Detail: e.DetailMessage(),
		Errors: newDetails(e.Details()),
		Title:  e.ErrorType().Title(),
		Type:   ErrorType(e.ErrorType().Type()),
	}
}

func newDetails(details []errors.APIErrorDetails) *[]ErrorDetail {
	if len(details) == 0 {
		return nil
	}

	r := make([]ErrorDetail, 0, len(details))

	for _, d := range details {
		r = append(r, ErrorDetail{
			Name:     d.Name(),
			Reason:   d.Reason(),
			Resource: d.Resource(),
		})
	}
	return &r
}
