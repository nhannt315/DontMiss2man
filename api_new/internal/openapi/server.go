package openapi

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true // no outputs. ⇨ http server started on [::]:8080
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
