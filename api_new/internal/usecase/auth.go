package usecase

import (
	"context"

	apperrors "github.com/nhannt315/real_estate_api/internal/errors"
	"github.com/nhannt315/real_estate_api/internal/model"
	"github.com/nhannt315/real_estate_api/internal/services/jwt"
	"github.com/nhannt315/real_estate_api/pkg/errors"

	"github.com/nhannt315/real_estate_api/internal/repository"
	"github.com/nhannt315/real_estate_api/pkg/password"
)

type Auth interface {
	LoginUser(ctx context.Context, email, password string) (*model.User, error)
	RegisterUser(ctx context.Context, email, password, passwordConfirmation string) (*model.User, error)
}

type auth struct {
	userRepo     repository.User
	jwtGenerator jwt.Generator
}

func NewAuthUseCase(userRepo repository.User, jwtGenerator jwt.Generator) Auth {
	return &auth{
		userRepo:     userRepo,
		jwtGenerator: jwtGenerator,
	}
}

func (u *auth) LoginUser(ctx context.Context, email, pwd string) (user *model.User, err error) {
	defer func() {
		if err != nil {
			err = apperrors.Wrap(err, apperrors.ErrorTypeInvalidCredential, "Login failed")
		}
	}()

	findUser, err := u.userRepo.WithContext(ctx).FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if ok := password.CheckPasswordHash(pwd, findUser.PasswordDigest); !ok {
		return nil, errors.New("Authentication failed")
	}

	token, err := u.jwtGenerator.GenerateJWTString(findUser.ID)
	if err != nil {
		return nil, errors.Wrap(err, "Generate jwt token")
	}
	findUser.AccessToken = token

	return findUser, nil
}

func (u *auth) RegisterUser(ctx context.Context, email, pwd, pwdConfirmation string) (user *model.User, err error) {
	if pwd != pwdConfirmation {
		return nil, apperrors.New(apperrors.ErrorTypeBadRequest, "Password not match")
	}

	hashedPassword, err := password.HashPassword(pwd)
	if err != nil {
		return nil, apperrors.Wrap(err, apperrors.ErrorTypeInternal, "Hash password")
	}
	user = &model.User{
		Email:          email,
		PasswordDigest: hashedPassword,
	}

	err = u.userRepo.WithContext(ctx).Create(user)
	if err != nil {
		return nil, apperrors.Wrap(err, apperrors.ErrorTypeBadRequest, err.Error())
	}

	accessToken, err := u.jwtGenerator.GenerateJWTString(user.ID)
	if err != nil {
		return nil, apperrors.Wrap(err, apperrors.ErrorTypeInternal, "Generate jwt token")
	}

	user.AccessToken = accessToken
	return user, nil
}
