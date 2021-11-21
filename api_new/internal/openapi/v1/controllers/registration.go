package controllers

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	apperrors "github.com/nhannt315/real_estate_api/internal/errors"
	openapiv1 "github.com/nhannt315/real_estate_api/internal/openapi/v1"
	"github.com/nhannt315/real_estate_api/internal/repository"
	"github.com/nhannt315/real_estate_api/internal/services/jwt"
	"github.com/nhannt315/real_estate_api/internal/usecase"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

type RegistrationController struct {
	logger       logs.Logger
	reg          repository.Registry
	jwtGenerator jwt.Generator
}

func NewRegistrationController(ictx *InitializeContext) *RegistrationController {
	return &RegistrationController{
		logger:       ictx.Logger,
		reg:          ictx.Reg,
		jwtGenerator: ictx.JWTGenerator,
	}
}

func (c *RegistrationController) Register(ectx echo.Context) error {
	req := ectx.Request()
	ctx := ectx.Request().Context()

	if req.ContentLength == 0 {
		return apperrors.New(apperrors.ErrorTypeBadRequest, "missing request body")
	}

	requestBody := new(openapiv1.RegistrationRequest)
	if err := json.NewDecoder(req.Body).Decode(requestBody); err != nil {
		return apperrors.New(apperrors.ErrorTypeInternal, "Error when decoding request body")
	}

	authUsecase := usecase.NewAuthUseCase(c.reg.UserRepository(), c.jwtGenerator)
	user, err := authUsecase.RegisterUser(ctx, requestBody.Email, requestBody.Password, requestBody.PasswordConfirmation)
	if err != nil {
		return err
	}

	return ectx.JSON(200, &openapiv1.AuthenticationResponse{
		Token: user.AccessToken,
		Email: user.Email,
	})
}
