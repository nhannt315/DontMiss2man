package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

// NewRecover panicのリカバー用ミドルウェア
func NewRecover(log logs.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) (retErr error) {
			ctx := ectx.Request().Context()
			defer func() {
				panicRet := recover()
				if panicRet == nil {
					return
				}

				if retErr != nil {
					// パニック発生 = internal errorとしておきたいので、handlerのエラーは返さない
					// ログだけだしておく
					log.AddStack(retErr).Warn(ctx, "Recover Panic. original error")
					return
				}

				switch e := panicRet.(type) {
				case error:
					retErr = errors.Wrap(e, "Recover Panic") // stackとりたいのでwrapしておく
				default:
					retErr = errors.Errorf("Recover Panic %v", e)
				}
			}()

			return next(ectx)

		}
	}
}
