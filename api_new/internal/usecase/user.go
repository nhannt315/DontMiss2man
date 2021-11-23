package usecase

import (
	"context"
	"github.com/nhannt315/real_estate_api/internal/model"
	"github.com/nhannt315/real_estate_api/internal/repository"
	"github.com/nhannt315/real_estate_api/internal/services/jwt"
)

type User interface {
	GetUserInfo(ctx context.Context, accessToken string) (*model.User, error)
}

type user struct {
	reg         repository.Registry
	jwtVerifier jwt.Verifier
}

func NewUserUseCase(reg repository.Registry, verifier jwt.Verifier) User {
	return &user{
		reg:         reg,
		jwtVerifier: verifier,
	}
}

func (u *user) GetUserInfo(ctx context.Context, accessToken string) (*model.User, error) {
	// verify access token
	userID, err := u.jwtVerifier.VerifyJWT(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	// Get user and return
	return u.reg.UserRepository().WithContext(ctx).FindByID(userID)
}
