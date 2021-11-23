package jwt

import (
	"context"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/jwk"
	"gopkg.in/square/go-jose.v2/jwt"
	"strconv"
)

type Verifier interface {
	VerifyJWT(ctx context.Context, jwtToken string) (uint64, error)
}

type verifier struct {
	jwkHelper jwk.Helper
}

func NewVerifier(jwkHelper jwk.Helper) Verifier {
	return &verifier{
		jwkHelper: jwkHelper,
	}
}

func (v *verifier) VerifyJWT(ctx context.Context, jwtToken string) (uint64, error) {
	token, err := jwt.ParseSigned(jwtToken)
	if err != nil {
		return 0, errors.Wrap(err, "Parse jwt access token")
	}
	claims := jwt.Claims{}
	if err = token.Claims(v.jwkHelper.JSONWebKeySet(), &claims); err != nil {
		return 0, errors.Wrap(err, "Error when verify claims")
	}

	userID, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "Parse user id from string")
	}

	return userID, nil
}
