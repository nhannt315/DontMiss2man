package openapi

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nhannt315/real_estate_api/internal/openapi/log"
)

type Server struct {
	e *echo.Echo
}

func NewServer(l *log.Logger) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true // no outputs. ⇨ http server started on [::]:8080
	e.HTTPErrorHandler = newErrorHandler(e, l)
	return &Server{e: e}
}

func (s *Server) RegisterHandler(method, path string, handler http.Handler) {
	echoHandler := func(echoCtx echo.Context) error {
		handler.ServeHTTP(echoCtx.Response().Writer, echoCtx.Request())
		return nil
	}
	s.e.Add(method, path, echoHandler)
}

func (s *Server) Start(address string) error {
	return s.e.Start(address)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}

func (s *Server) RegisterMiddleware(mw ...echo.MiddlewareFunc) {
	s.e.Use(mw...)
}

func (s *Server) Group(path string, middleware ...echo.MiddlewareFunc) *echo.Group {
	return s.e.Group(path, middleware...)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.e.ServeHTTP(w, r)
}

// RegisterErrorHandler エラーハンドラを登録する
func RegisterErrorHandler(h ErrorHandler) {
	errorHandlers = append(errorHandlers, h)
}

var errorHandlers []ErrorHandler

func newErrorHandler(e *echo.Echo, logger *log.Logger) echo.HTTPErrorHandler {
	return func(err error, ectx echo.Context) {
		defer func() {
			logger.LogResponse(ectx, err)
		}()

		for _, h := range errorHandlers {
			if h(err, ectx) {
				return
			}
		}
		e.DefaultHTTPErrorHandler(err, ectx)
	}
}

// ErrorHandler ectx にあるリクエスト情報を見て、エラー処理を行うかを判定。処理を行った場合 trueをリターンする
type ErrorHandler = func(err error, ectx echo.Context) (processed bool)
