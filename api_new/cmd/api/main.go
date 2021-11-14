package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nhannt315/real_estate_api/pkg/appinfo"
	pkgconf "github.com/nhannt315/real_estate_api/pkg/config"
	"github.com/nhannt315/real_estate_api/pkg/datetime"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/goroutine"
	goroutine_interceptor "github.com/nhannt315/real_estate_api/pkg/goroutine/interceptors"
	"github.com/nhannt315/real_estate_api/pkg/logs/zap"
	"github.com/nhannt315/real_estate_api/pkg/rollbar"

	"github.com/nhannt315/real_estate_api/internal/config"
	"github.com/nhannt315/real_estate_api/internal/initialize"

	main_db "github.com/nhannt315/real_estate_api/internal/initialize/db"
	main_oapi "github.com/nhannt315/real_estate_api/internal/initialize/openapi"
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

	shutdownTasks := initialize.NewShutdownTasks(logger)

	defer func() {
		shutdownTasks.ExecuteAll(ctx)
	}()

	// init goroutine interceptors
	goroutine.RegisterInterceptor(goroutine_interceptor.Chain(
		goroutine_interceptor.RecoveryInterceptor(logger),
	))

	// init and start database connection
	dbTask, err := main_db.InitializeAndStart(ctx, &appConf, logger)
	if err != nil {
		return err
	}
	shutdownTasks.Add(dbTask)

	// register app & db location to clarify the app and db layer timezone
	if err = registerAppLocation(appConf.AppLocation); err != nil {
		return err
	}
	if err = registerDBLocation(appConf.DBConfig.Location); err != nil {
		return err
	}

	oapiServerTask, err := main_oapi.Initialize(ctx, logger, appConf.OpenAPIConfig)
	if err != nil {
		return err
	}
	err = oapiServerTask.Start(ctx)
	if err != nil {
		return err
	}
	shutdownTasks.Add(oapiServerTask)

	// handle server stopping
	initialize.WaitForServerStop(ctx, logger)

	// 通常の停止のためshutdownログ出力前に終了処理する
	shutdownTasks.ExecuteAll(ctx)

	logger.Info(ctx, "server shutdown gracefully")

	return nil
}

// registerAppLocation registers location by given name
// to all packages that need registering location
// to themselves as app location.
func registerAppLocation(appLocStr string) error {
	appLoc, err := datetime.LoadLocation(appLocStr)
	if err != nil {
		return errors.Wrapf(err, "cannot load app location %s", appLocStr)
	}
	if err = datetime.SetAppLocation(appLoc); err != nil {
		return errors.Wrap(err, "cannot set app location to times package")
	}

	return nil
}

// registerDBLocation registers location by given name
// to all packages that need registering location
// to themselves as db location.
func registerDBLocation(dbLocStr string) error {
	dbLoc, err := datetime.LoadLocation(dbLocStr)
	if err != nil {
		return errors.Wrapf(err, "cannot load database location %s", dbLocStr)
	}
	if err = datetime.SetDBLocation(dbLoc); err != nil {
		return errors.Wrap(err, "cannot set database location to times package")
	}

	return nil
}
