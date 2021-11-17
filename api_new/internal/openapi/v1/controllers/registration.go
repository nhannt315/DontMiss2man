package controllers

import (
	"github.com/labstack/echo/v4"
	openapiv1 "github.com/nhannt315/real_estate_api/internal/openapi/v1"
	"github.com/nhannt315/real_estate_api/internal/repository"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

type RegistrationController struct {
	logger logs.Logger
	reg    repository.Registry
}

func NewRegistrationController(ictx *InitializeContext) *RegistrationController {
	return &RegistrationController{
		logger: ictx.Logger,
		reg:    ictx.Reg,
	}
}

func (c *RegistrationController) Register(ectx echo.Context) error {
	return ectx.JSON(200, &openapiv1.AuthenticationResponse{
		Token: "fdsa",
		Email: "test@gmail.com",
	})
}
