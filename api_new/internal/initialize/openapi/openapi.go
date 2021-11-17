package openapi

import (
	"context"
	"net/http"

	"github.com/nhannt315/real_estate_api/internal/config"
	"github.com/nhannt315/real_estate_api/internal/openapi"
	oapilog "github.com/nhannt315/real_estate_api/internal/openapi/log"
	oapi_middlewares "github.com/nhannt315/real_estate_api/internal/openapi/middlewares"
	openapiv1_server "github.com/nhannt315/real_estate_api/internal/openapi/v1/controllers"
	"github.com/nhannt315/real_estate_api/internal/repository"
	"github.com/nhannt315/real_estate_api/internal/services/jwt"
	"github.com/nhannt315/real_estate_api/pkg/datetime"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/goroutine"
	"github.com/nhannt315/real_estate_api/pkg/jwk"
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
	appConfig *config.Config,
	reg repository.Registry,
) (*Task, error) {

	oapiLogger := oapilog.NewLogger(logger)

	apiServer := openapi.NewServer(oapiLogger)
	apiServer.RegisterMiddleware(
		oapi_middlewares.NewRecover(logger),
		oapi_middlewares.NewContext(),
		oapi_middlewares.NewLogger(oapiLogger),
	)

	dateTimeManager := datetime.NewManager()
	jwkHelper, err := jwk.NewHelper(appConfig.JWTConfig.KeyID, appConfig.JWTConfig.PrivateKey)
	if err != nil {
		return nil, err
	}

	jwtGenerator, err := jwt.NewGenerator(appConfig.JWTConfig, jwkHelper, dateTimeManager)
	if err != nil {
		return nil, err
	}

	ictx := &openapiv1_server.InitializeContext{
		Logger:          logger,
		Reg:             reg,
		AppConf:         appConfig,
		JWTGenerator:    jwtGenerator,
		DateTimeManager: dateTimeManager,
	}

	openapiv1_server.RegisterHandler(ictx, apiServer)

	return &Task{server: apiServer, logger: logger, apiServerConfig: appConfig.OpenAPIConfig}, nil
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
