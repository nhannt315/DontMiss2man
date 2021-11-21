package test

import (
	"github.com/nhannt315/real_estate_api/internal/config"
	pkgconfig "github.com/nhannt315/real_estate_api/pkg/config"
)

func NewTestConfig() *config.Config {
	conf := &config.Config{}
	if err := pkgconfig.Load(conf, RootDir()+"/test/config/test.yml"); err != nil {
		panic(err)
	}

	return conf
}
