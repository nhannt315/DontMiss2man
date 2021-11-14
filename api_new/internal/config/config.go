package config

import (
	"github.com/nhannt315/real_estate_api/pkg/logs"
	"github.com/nhannt315/real_estate_api/pkg/rollbar"
)

type Config struct {
	AppLocation string          `yaml:"app_location"`
	Rollbar     *rollbar.Config `yaml:"rollbar"`
	Logger      *logs.Config    `yaml:"logger"`
}
