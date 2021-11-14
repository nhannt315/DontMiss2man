package openapi

import (
	"context"
	"net/http"

	"github.com/nhannt315/real_estate_api/internal/openapi"
	oapilog "github.com/nhannt315/real_estate_api/internal/openapi/log"
	oapi_middlewares "github.com/nhannt315/real_estate_api/internal/openapi/middlewares"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/goroutine"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

type Task struct {
	apiServerConfig *openapi.Config
	server          *openapi.Server
	logger          logs.Logger
}

func Initialize(
	ctx context.Context,
	logger logs.Logger,
	apiServerConfig *openapi.Config,
) (*Task, error) {

	oapiLogger := oapilog.NewLogger(logger)

	apiServer := openapi.NewServer(oapiLogger)
	apiServer.RegisterMiddleware(
		oapi_middlewares.NewRecover(logger),
		oapi_middlewares.NewContext(),
		oapi_middlewares.NewLogger(oapiLogger),
	)

	return &Task{server: apiServer, logger: logger, apiServerConfig: apiServerConfig}, nil
}

func (t *Task) Start(ctx context.Context) (err error) {
	goroutine.GoWithContext(ctx, func(innerCtx context.Context) {
		t.logger.Infof(innerCtx, "Start api server at addr %s", t.apiServerConfig.Address)
		if err := t.server.Start(t.apiServerConfig.Address); err != nil && errors.Is(err, http.ErrServerClosed) {
			t.logger.Error(innerCtx, errors.Wrapf(err, "Start api server at addr %s", t.apiServerConfig.Address))
			return
		}
		t.logger.Infof(innerCtx, "Stop api server at address %s", t.apiServerConfig.Address)
	})

	return nil
}

func (t *Task) Shutdown(ctx context.Context) (err error) {
	return t.server.Stop(ctx)
}

// Name return name
func (t *Task) Name() string {
	return "open api server"
}
