package requestid

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ContextKey コンテキストのキー
type ContextKey string

const (
	requestKey ContextKey = "request_id"
)

// GetRequestID Context中のrequestID取得
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	requestID, ok := ctx.Value(requestKey).(string)
	if !ok || requestID == "" {
		return ""
	}
	return requestID
}

// SetRequestID RequestIDをcontextにsetする
func SetRequestID(ctx context.Context, req *http.Request) context.Context {
	// リクエストID
	requestID := req.Header.Get(echo.HeaderXRequestID)
	if requestID == "" {
		return ctx
	}
	return context.WithValue(ctx, requestKey, requestID)
}
