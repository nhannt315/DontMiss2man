package openapi

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/nhannt315/real_estate_api/internal/config"
	"github.com/nhannt315/real_estate_api/internal/openapi"
	oapilog "github.com/nhannt315/real_estate_api/internal/openapi/log"
	oapi_middlewares "github.com/nhannt315/real_estate_api/internal/openapi/middlewares"
	openapiv1 "github.com/nhannt315/real_estate_api/internal/openapi/v1"
	openapiv1_controllers "github.com/nhannt315/real_estate_api/internal/openapi/v1/controllers"
	"github.com/nhannt315/real_estate_api/internal/test"
)

type TestServer struct {
	appConf           *config.Config
	initializeContext *openapiv1_controllers.InitializeContext
	server            *openapi.Server
}

func NewTestServer(ctx context.Context, ictx *openapiv1_controllers.InitializeContext) (*TestServer, error) {
	appConfig := test.NewTestConfig()

	oapiLogger := oapilog.NewLogger(ictx.Logger)
	s := &TestServer{
		appConf:           appConfig,
		initializeContext: ictx,
		server:            openapi.NewServer(oapiLogger),
	}

	s.server.RegisterMiddleware(oapi_middlewares.NewLogger(oapiLogger))
	openapiv1_controllers.RegisterHandler(ictx, s.server)

	return s, nil
}

func (s *TestServer) InvokeRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	s.server.ServeHTTP(rec, req)
	return rec
}

func (s *TestServer) Get(path string, opts ...TestRequestOption) *httptest.ResponseRecorder {
	target := fmt.Sprintf("%s%s", openapiv1.BasePath, path)
	req := httptest.NewRequest(echo.GET, target, bytes.NewReader([]byte{}))
	for _, o := range opts {
		o(req)
	}

	return s.InvokeRequest(req)
}

func (s *TestServer) PostJSON(path, body string, opts ...TestRequestOption) *httptest.ResponseRecorder {
	target := fmt.Sprintf("%s%s", openapiv1.BasePath, path)
	req := httptest.NewRequest(echo.POST, target, bytes.NewReader([]byte(body)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	for _, o := range opts {
		o(req)
	}

	return s.InvokeRequest(req)
}

// TestRequestOption requestへの option指定用func
type TestRequestOption func(req *http.Request)
