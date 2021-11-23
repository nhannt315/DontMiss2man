package jwk

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"

	"github.com/nhannt315/real_estate_api/pkg/errors"
	"gopkg.in/square/go-jose.v2"
)

type Helper struct {
	keyID, rawPrivateKey string
	rsaPrivateKey        *rsa.PrivateKey
	jsonWebKeySet        *jose.JSONWebKeySet
}

func NewHelper(keyID string, encodedPrivateKey string) (*Helper, error) {
	// Decode base64 private key
	privateKeyPem, err := base64.StdEncoding.DecodeString(encodedPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	// Convert to rsa private key from pem format
	block, _ := pem.Decode(privateKeyPem)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	jsonWebKey := jose.JSONWebKey{
		Key:       &privateKey.PublicKey,
		KeyID:     keyID,
		Algorithm: string(jose.RS256),
		Use:       "sig",
	}
	jsonWebKeySet := jose.JSONWebKeySet{Keys: []jose.JSONWebKey{jsonWebKey}}

	return &Helper{
		keyID:         keyID,
		rawPrivateKey: string(privateKeyPem),
		rsaPrivateKey: privateKey,
		jsonWebKeySet: &jsonWebKeySet,
	}, nil
}

func (h *Helper) JSONWebKeySetJSON() ([]byte, error) {
	return json.Marshal(h.jsonWebKeySet)
}

func (h *Helper) RSAPrivateKey() *rsa.PrivateKey {
	return h.rsaPrivateKey
}

func (h *Helper) JSONWebKeySet() *jose.JSONWebKeySet {
	return h.jsonWebKeySet
}
