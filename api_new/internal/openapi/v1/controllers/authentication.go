package controllers

import (
	"github.com/labstack/echo/v4"
	openapiv1 "github.com/nhannt315/real_estate_api/internal/openapi/v1"
	"github.com/nhannt315/real_estate_api/internal/repository"
	"github.com/nhannt315/real_estate_api/internal/services/jwt"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

type AuthenticationController struct {
	logger       logs.Logger
	jwtGenerator jwt.Generator
	reg          repository.Registry
}

func NewAuthenticationController(ictx *InitializeContext) *AuthenticationController {
	return &AuthenticationController{
		logger:       ictx.Logger,
		reg:          ictx.Reg,
		jwtGenerator: ictx.JWTGenerator,
	}
}

func (c *AuthenticationController) Login(ectx echo.Context) error {
	token, err := c.jwtGenerator.GenerateJWTString(1)
	if err != nil {
		return err
	}
	return ectx.JSON(200, &openapiv1.AuthenticationResponse{
		Token: token,
		Email: "test@gmail.com",
	})
}
