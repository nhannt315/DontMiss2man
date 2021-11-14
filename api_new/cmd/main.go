package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nhannt315/real_estate_api/pkg/appinfo"
	pkgconf "github.com/nhannt315/real_estate_api/pkg/config"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/logs/zap"
	"github.com/nhannt315/real_estate_api/pkg/rollbar"

	"github.com/nhannt315/real_estate_api/internal/config"
)

func main() {
	ctx := context.Background()
	if err := doInit(ctx); err != nil {
		fmt.Printf("%+v", err)
		panic(err)
	}
}

func doInit(ctx context.Context) error {
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		configFilePath = "config/local.yml"
	}
	var appConf config.Config
	if err := pkgconf.Load(&appConf, configFilePath); err != nil {
		return errors.Wrap(err, "fail to load config file")
	}

	rb := rollbar.NewClient(appConf.Rollbar, appinfo.SimpleVersion())
	logger := zap.NewLogger(appConf.Logger, rb)

	logger.Infof(ctx, "Real estate api server started at port %d", 4000)

	return nil
}
