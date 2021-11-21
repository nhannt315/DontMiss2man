package test

import (
	"context"
	"database/sql"
	"path/filepath"
	"sync"

	"github.com/nhannt315/real_estate_api/internal/config"
	"github.com/nhannt315/real_estate_api/pkg/db"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/logs/zap"
	pkgtest "github.com/nhannt315/real_estate_api/pkg/test"
	"github.com/pressly/goose"
)

var testSQLDB *sql.DB
var testSchemaName string
var testDBMutex sync.Mutex

func InitializeTestDB(config *config.Config) (func(), error) {
	testDBMutex.Lock()
	defer testDBMutex.Unlock()

	if testSQLDB != nil {
		return nil, errors.New("Test database already initialized")
	}

	return initTestDB(config)
}

func initTestDB(config *config.Config) (func(), error) {
	ctx := context.Background()
	logger := zap.NewLogger(config.Logger, nil)
	testConf := config.DBConfig.Clone()
	testConf.Schema = ""

	dbConnForCreation, err := db.NewDBConn(testConf)
	if err != nil {
		return nil, err
	}

	defer func() { dbConnForCreation.Close() }()

	const dbNamePrefix = "real_estate"
	schema, err := pkgtest.CreateEmptyDB(dbConnForCreation, logger, dbNamePrefix)
	if err != nil {
		return nil, err
	}

	config.DBConfig.Schema = schema

	mainDBConn, err := db.NewDBConn(config.DBConfig)
	if err != nil {
		return nil, err
	}

	if err = migrate(mainDBConn); err != nil {
		return nil, err
	}

	testSchemaName = schema
	testSQLDB = mainDBConn

	logger.Infof(ctx, "create %s", testSchemaName)

	return func() {
		logger.Infof(ctx, "drop %s", testSchemaName)
		// nolint: errcheck
		testSQLDB.Exec("DROP DATABASE " + testSchemaName)
		err := testSQLDB.Close()
		if err != nil {
			return
		}
		testSQLDB = nil
		testSchemaName = ""
	}, nil
}

func migrate(db *sql.DB) error {
	pjRoot, err := PjRootDir()
	if err != nil {
		return err
	}

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}
	goose.SetVerbose(false)
	dir := filepath.Join(pjRoot, "deployments", "migration", "migrations")
	if err := goose.Run("up", db, dir); err != nil {
		return errors.Wrap(err, "migrate failed")
	}

	return nil
}
