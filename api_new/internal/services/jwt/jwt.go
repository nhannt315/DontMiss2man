package jwt

import (
	"strconv"

	"github.com/nhannt315/real_estate_api/pkg/datetime"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/jwk"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type Generator interface {
	GenerateJWTString(userID uint64) (string, error)
}

type generator struct {
	tokenSigner     jose.Signer
	config          *Config
	dateTimeManager datetime.Manager
}

func NewGenerator(config *Config, jwkHelper *jwk.Helper, dateTimeManager datetime.Manager) (Generator, error) {
	headerOptions := make(map[jose.HeaderKey]interface{})
	headerOptions["typ"] = "JWT"
	headerOptions["kid"] = config.KeyID

	privateKey := jwkHelper.RSAPrivateKey()

	signer, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.SignatureAlgorithm(config.SigningAlg),
		Key:       privateKey,
	}, &jose.SignerOptions{
		ExtraHeaders: headerOptions,
	})

	if err != nil {
		return nil, errors.Wrap(err, "Error when generating JWT signer")
	}

	return &generator{
		config:          config,
		tokenSigner:     signer,
		dateTimeManager: dateTimeManager,
	}, nil
}

func (g *generator) GenerateJWTString(userID uint64) (string, error) {
	iat := g.dateTimeManager.Clock().Now().ToTime()
	exp := g.dateTimeManager.Clock().FromNow(g.config.Expiration).ToTime()

	claims := &jwt.Claims{
		Issuer:   g.config.Issuer,
		Subject:  strconv.FormatUint(userID, 10),
		Audience: jwt.Audience{g.config.Issuer},
		Expiry:   jwt.NewNumericDate(exp),
		IssuedAt: jwt.NewNumericDate(iat),
	}

	result, err := jwt.Signed(g.tokenSigner).Claims(claims).CompactSerialize()
	if err != nil {
		return "", errors.Wrap(err, "Serialize ID Token")
	}

	return result, nil
}
