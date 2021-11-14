package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/nhannt315/real_estate_api/internal/openapi/log"
	"github.com/nhannt315/real_estate_api/pkg/errors"
)

// NewLogger ログ出力ミドルウェア
func NewLogger(l *log.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			if err := l.LogRequest(ectx); err != nil { // 処理継続不能なエラーなのでnextは呼ばない
				return errors.Wrap(err, "fail to read request body")
			}

			err := next(ectx)

			if err == nil {
				l.LogResponse(ectx, nil)
			}

			// エラー時のレスポンスの内容はerrorHandlerで処理するまで確定しないので、ログもerrorHandlerに任せる
			return err
		}
	}
}
