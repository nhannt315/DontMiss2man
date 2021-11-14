package db

import (
	"context"
	"database/sql"

	"github.com/nhannt315/real_estate_api/pkg/logs"

	"github.com/nhannt315/real_estate_api/pkg/errors"

	"github.com/nhannt315/real_estate_api/internal/config"
	"github.com/nhannt315/real_estate_api/pkg/db"
	"gorm.io/gorm"
)

// Task implements db's initialize and other function
type Task struct {
	dbConfig *db.Config
	gormDB   *gorm.DB
	sqlDB    *sql.DB
}

// InitializeAndStart initialize db connection and start
func InitializeAndStart(ctx context.Context, appConfig *config.Config, logger logs.Logger) (*Task, error) {
	sqlDB, err := db.NewDBConn(appConfig.DBConfig)
	if err != nil {
		return nil, err
	}

	gormDB, err := db.NewDB(ctx, appConfig.DBConfig, logger, sqlDB, appConfig.Logger.Level)

	if err != nil {
		return nil, err
	}
	logger.Info(ctx, "Connected to database successfully")

	return &Task{
		sqlDB:    sqlDB,
		gormDB:   gormDB,
		dbConfig: appConfig.DBConfig,
	}, nil
}

// Name return name
func (t *Task) Name() string {
	return "database"
}

// Shutdown shutdown function
func (t *Task) Shutdown(ctx context.Context) error {
	err := t.sqlDB.Close()
	if err != nil {
		return errors.Wrap(err, "Close sql DB connection")
	}

	return nil
}

// GormDB return gorm database
func (t *Task) GormDB() *gorm.DB {
	return t.gormDB
}
