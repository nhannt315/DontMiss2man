package controllers

import (
	"github.com/nhannt315/real_estate_api/internal/config"
	"github.com/nhannt315/real_estate_api/internal/openapi"
	openapiv1 "github.com/nhannt315/real_estate_api/internal/openapi/v1"
	"github.com/nhannt315/real_estate_api/internal/repository"
	"github.com/nhannt315/real_estate_api/internal/services/jwt"
	"github.com/nhannt315/real_estate_api/pkg/datetime"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

type handler struct {
	*RegistrationController
	*AuthenticationController
}

type InitializeContext struct {
	Logger          logs.Logger
	AppConf         *config.Config
	Reg             repository.Registry
	JWTGenerator    jwt.Generator
	DateTimeManager datetime.Manager
}

func newHandler(ictx *InitializeContext) *handler {
	return &handler{
		RegistrationController:   NewRegistrationController(ictx),
		AuthenticationController: NewAuthenticationController(ictx),
	}
}

// RegisterHandler  OpenAPIリクエストハンドラを登録する
func RegisterHandler(ictx *InitializeContext, server *openapi.Server) {
	g := server.Group(openapiv1.BasePath)

	handler := newHandler(ictx)
	openapiv1.RegisterHandlers(g, handler)
}
