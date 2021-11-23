package jwt_test

import (
	"context"
	"github.com/nhannt315/real_estate_api/internal/services/jwt"
	"github.com/nhannt315/real_estate_api/internal/test"
	"github.com/nhannt315/real_estate_api/pkg/jwk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_verifier_VerifyJWT(t *testing.T) {
	config := test.NewTestConfig()
	ctx := context.Background()

	tests := []struct {
		name    string
		token   string
		want    uint64
		wantErr bool
	}{
		{
			name: "Verify jwt token",
			token: "eyJhbGciOiJSUzI1NiIsImtpZCI6InJlYWwtZXN0YXRlIiwidHlwIjoiSldUIn0.eyJhdWQiOlsicmVhbC1lc3RhdGUiXSwiZXhwIjoxNjM3NDk3NjI1LCJpYXQiOjE2Mzc0OTQwMjUsImlzcyI6InJlYWwtZXN0YXRlIiwic3ViIjoiMSJ9.R1ViEdSmNEcrEh9lpRYAVivfjP4xvfpSwdx2XNgerpcBTxBOJBL_RBD84S_QG_4yvsQwNnTjMuA3biy_i3bYWNc2S_u3uL_jdoYLHf_GKZQR6pKwuKOj2T9rg95UIujTdNttiGsEJ4ehiMzLDFD7JIfWMGj2C5So48RrELN-JAI",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jwkHelper, err := jwk.NewHelper(config.JWTConfig.KeyID, config.JWTConfig.PrivateKey)
			if err != nil {
				t.Error(err)
				return
			}
			verifier := jwt.NewVerifier(jwkHelper)
			got, err := verifier.VerifyJWT(ctx, tt.token)
			if err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
