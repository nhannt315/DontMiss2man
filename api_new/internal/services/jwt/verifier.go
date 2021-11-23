package jwt

import (
	"context"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/jwk"
	"gopkg.in/square/go-jose.v2/jwt"
)

type Verifier interface {
	VerifyJWT(ctx context.Context, jwtToken string) (string, error)
}

type verifier struct {
	jwkHelper *jwk.Helper
}

func NewVerifier(jwkHelper *jwk.Helper) Verifier {
	return &verifier{
		jwkHelper: jwkHelper,
	}
}

func (v *verifier) VerifyJWT(ctx context.Context, jwtToken string) (string, error) {
	token, err := jwt.ParseSigned(jwtToken)
	if err != nil {
		return "", errors.Wrap(err, "Parse jwt access token")
	}
	claims := jwt.Claims{}
	if err = token.Claims(v.jwkHelper.JSONWebKeySet(), &claims); err != nil {
		return "", errors.Wrap(err, "Error when verify claims")
	}

	return claims.Subject, nil
}
