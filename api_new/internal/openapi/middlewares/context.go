package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/nhannt315/real_estate_api/pkg/goroutine"
	"github.com/nhannt315/real_estate_api/pkg/requestid"
)

// NewContext コンテキストにUno共通で利用する値をセットする
func NewContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			ctx := ectx.Request().Context()
			goroutineID := goroutine.IDValue(ctx)
			if goroutineID == 0 {
				ctx = goroutine.SetGoroutineID(ctx)
			}
			// リクエストID
			req := ectx.Request()
			ctx = requestid.SetRequestID(ctx, req)

			newReq := ectx.Request().WithContext(ctx)
			ectx.SetRequest(newReq)
			return next(ectx)
		}
	}
}
