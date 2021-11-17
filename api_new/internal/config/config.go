package config

import (
	"github.com/nhannt315/real_estate_api/internal/openapi"
	"github.com/nhannt315/real_estate_api/internal/services/jwt"
	"github.com/nhannt315/real_estate_api/pkg/db"
	"github.com/nhannt315/real_estate_api/pkg/logs"
	"github.com/nhannt315/real_estate_api/pkg/rollbar"
)

type Config struct {
	AppLocation   string          `yaml:"app_location"`
	Rollbar       *rollbar.Config `yaml:"rollbar"`
	Logger        *logs.Config    `yaml:"logger"`
	DBConfig      *db.Config      `yaml:"db"`
	OpenAPIConfig *openapi.Config `yaml:"openapi"`
	JWTConfig     *jwt.Config     `yaml:"jwt"`
}
