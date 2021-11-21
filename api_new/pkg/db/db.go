package db

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/logs"
	"github.com/nhannt315/real_estate_api/pkg/retry"

	gormMySQLDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const (
	tlsConfigKey  = "custom"
	tlsSkipVerify = "skip-verify"
	tlsTrue       = "true"
	tlsFalse      = "false"
)

// NewDBConn initializes DB connection.
func NewDBConn(dbConfig *Config) (conn *sql.DB, err error) {
	if dbConfig.CACertPath != "" {
		if dbConfig.TLS != tlsTrue {
			return nil, errors.Errorf("CA Cert Path isn't enabled in this mode %s", dbConfig.TLS)
		}
		if err := registerTLSConfig(dbConfig.CACertPath); err != nil {
			return nil, err
		}
	}

	var dsn string
	switch dbConfig.Driver {
	case driverMySQL:
		dsn, err = buildMySQLConnectionString(dbConfig)
		if err != nil {
			return nil, err
		}
		conn, err = sql.Open(driverMySQL, dsn)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.Errorf("not support db driver %s", dbConfig.Driver)
	}

	conn.SetMaxOpenConns(int(dbConfig.MaxOpenConn))
	conn.SetMaxIdleConns(int(dbConfig.MaxIdleConn))
	conn.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)

	return conn, nil
}

// NewDB initializes DB.
func NewDB(ctx context.Context, dbConfig *Config, logger logs.Logger, conn *sql.DB, logLevel logs.Level) (db *gorm.DB, err error) {

	retryer := retry.NewFixed(
		retry.WithInterval(dbConfig.RetryInterval),
		retry.WithNotify(func(ctx context.Context, retryCount uint64, err error, interval time.Duration) {
			if retryCount == 1 {
				logger.Warnf(ctx, "retrying to connect DB: count=%d interval=%s err={%v} settings=%s",
					retryCount, interval, err, dbConfig.Info())
			} else {
				logger.Infof(ctx, "retrying to connect DB: count=%d interval=%s err={%v} settings=%s",
					retryCount, interval, err, dbConfig.Info())
			}
		}),
		retry.WithMaxRetryCount(dbConfig.MaxRetryCount),
	)

	err = retryer.Execute(ctx, func(ctx context.Context) (stopRetry bool, err error) {
		// Connect to database
		db, err = gorm.Open(gormMySQLDriver.New(gormMySQLDriver.Config{
			Conn: conn,
		}), &gorm.Config{
			Logger: gormLogger.New(logger,
				gormLogger.Config{
					LogLevel: logLevelToGormLogLevel(logLevel), // 開発用にSQLログを出力
				}),
		})

		if err != nil {
			return false, err
		}

		return true, nil
	})

	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect to db: %s", dbConfig.Info())
	}

	return db, nil
}

func logLevelToGormLogLevel(logLevel logs.Level) gormLogger.LogLevel {
	switch logLevel {
	case logs.DebugLevel:
		return gormLogger.Info
	case logs.InfoLevel:
		return gormLogger.Info
	case logs.WarnLevel:
		return gormLogger.Warn
	case logs.ErrorLevel:
		return gormLogger.Error
	default:
		return gormLogger.Info
	}
}

// buildMySQLConnectionString builds mysql connection string.
func buildMySQLConnectionString(dbConfig *Config) (string, error) {
	mysqlCfg := mysql.NewConfig()

	mysqlCfg.Net = "tcp"
	port := int(dbConfig.Port)

	// Hostにport番号を含めた場合は portの設定を無視する
	if strings.Contains(dbConfig.Host, ":") {
		mysqlCfg.Addr = dbConfig.Host
	} else {
		mysqlCfg.Addr = fmt.Sprintf("%s:%d", dbConfig.Host, port)
	}

	mysqlCfg.DBName = dbConfig.Schema
	mysqlCfg.User = dbConfig.Username.PlaneString()
	mysqlCfg.Passwd = dbConfig.Password.PlaneString()

	mysqlCfg.ParseTime = dbConfig.ParseTime // goの場合は基本的にtrue必須
	switch dbConfig.TLS {
	case tlsTrue:
		if dbConfig.CACertPath != "" {
			mysqlCfg.TLSConfig = tlsConfigKey
			break
		}
		mysqlCfg.TLSConfig = dbConfig.TLS
	case tlsFalse, tlsSkipVerify:
		mysqlCfg.TLSConfig = dbConfig.TLS
	case "":
		// noop
	default:
		return "", errors.Errorf("unknown value for TLS: %v", dbConfig.TLS)
	}
	// Set location
	loc, err := time.LoadLocation(dbConfig.Location)
	if err != nil {
		return "", err
	}
	mysqlCfg.Loc = loc

	if dbConfig.MaxAllowedPacket != 0 {
		mysqlCfg.MaxAllowedPacket = dbConfig.MaxAllowedPacket
	}
	mysqlCfg.InterpolateParams = dbConfig.InterpolateParams
	if dbConfig.Collation != "" {
		mysqlCfg.Collation = dbConfig.Collation
	}
	ret := mysqlCfg.FormatDSN()
	return ret, nil
}

// NewTestDB テスト用のDBインスタンスをつくる
func NewTestDB(logger logs.Logger, conn *sql.DB, logLevel logs.Level) (*gorm.DB, error) {
	return gorm.Open(gormMySQLDriver.New(gormMySQLDriver.Config{
		Conn: conn,
	}), &gorm.Config{
		Logger: gormLogger.New(logger,
			gormLogger.Config{
				LogLevel: logLevelToGormLogLevel(logLevel), // 開発用にSQLログを出力
			}),
	})
}

// nolint: gosec
func registerTLSConfig(caCertPath string) error {
	caCertPool := x509.NewCertPool()
	cert, err := os.ReadFile(caCertPath)
	if err != nil {
		return err
	}

	if ok := caCertPool.AppendCertsFromPEM(cert); !ok {
		return errors.New("failed to append pem to ssl cert pool")
	}
	tlsConfig := tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: false,
	}

	return mysql.RegisterTLSConfig(tlsConfigKey, &tlsConfig)
}
