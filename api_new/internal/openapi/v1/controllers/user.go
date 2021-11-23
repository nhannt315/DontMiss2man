package controllers

import (
	"github.com/labstack/echo/v4"
	apperrors "github.com/nhannt315/real_estate_api/internal/errors"
	"github.com/nhannt315/real_estate_api/internal/openapi/utils"
	openapiv1 "github.com/nhannt315/real_estate_api/internal/openapi/v1"
	"github.com/nhannt315/real_estate_api/internal/repository"
	"github.com/nhannt315/real_estate_api/internal/services/jwt"
	"github.com/nhannt315/real_estate_api/internal/usecase"
	"github.com/nhannt315/real_estate_api/pkg/logs"
	"net/http"
)

type UserController struct {
	logger      logs.Logger
	reg         repository.Registry
	jwtVerifier jwt.Verifier
}

func NewUserController(ictx *InitializeContext) *UserController {
	return &UserController{
		logger:      ictx.Logger,
		reg:         ictx.Reg,
		jwtVerifier: ictx.JWTVerifier,
	}
}

func (c *UserController) GetUserInfo(ectx echo.Context) error {
	accessToken := utils.GetAccessTokenFromHeader(ectx.Request())
	if accessToken == "" {
		return apperrors.New(apperrors.ErrorTypeInvalidCredential, "Access token is required")
	}

	userUseCase := usecase.NewUserUseCase(c.reg, c.jwtVerifier)

	user, err := userUseCase.GetUserInfo(ectx.Request().Context(), accessToken)
	if err != nil {
		return err
	}

	return ectx.JSON(http.StatusOK, &openapiv1.UserInfoResponse{Email: user.Email})
}
