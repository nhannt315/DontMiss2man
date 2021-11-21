package test

import (
	"context"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	"github.com/nhannt315/real_estate_api/internal/config"
	"github.com/nhannt315/real_estate_api/internal/repository"
	"github.com/nhannt315/real_estate_api/pkg/db"
	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/logs"
	"github.com/nhannt315/real_estate_api/pkg/logs/zap"
	"gorm.io/gorm"
)

type Helper struct {
	logger     *zap.TestLogger
	testGormDB *gorm.DB
	testDBOnce sync.Once
	testConf   *config.Config
}

func NewHelper(t *testing.T) *Helper {
	return &Helper{
		logger:     zap.NewTest(t),
		testDBOnce: sync.Once{},
		testConf:   NewTestConfig(),
	}
}

func PjRootDir() (string, error) {
	// nolint: gosec
	//G204: Subprocess launching should be audited (gosec)
	out, err := exec.Command("go", "env", "GOMOD").Output()
	if err != nil {
		return "", err
	}

	gomodPath := string(out)
	return filepath.Dir(gomodPath), nil
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func (h *Helper) DB() *gorm.DB {
	h.testDBOnce.Do(func() {
		if testSQLDB == nil || testSchemaName == "" {
			panic("Test database not initialized")
		}
		testGormDB, err := db.NewTestDB(h.logger, testSQLDB, h.testConf.Logger.Level)
		if err != nil {
			panic(err)
		}

		h.testGormDB = testGormDB
	})

	return h.testGormDB
}

func (h *Helper) Logger() logs.Logger {
	return h.logger
}

func (h *Helper) Registry() repository.Registry {
	return repository.NewRegistry(h.DB())
}

func (h *Helper) ClearDB(ctx context.Context) error {
	if testSQLDB == nil {
		return errors.New("Db connection is not initialized")
	}

	h.DB().Exec("SET FOREIGN_KEY_CHECKS=0")

	h.DB().WithContext(ctx).Exec("TRUNCATE TABLE `users`")
	h.DB().WithContext(ctx).Exec("TRUNCATE TABLE `agents`")
	h.DB().WithContext(ctx).Exec("TRUNCATE TABLE `buildings`")
	h.DB().WithContext(ctx).Exec("TRUNCATE TABLE `rooms`")
	h.DB().WithContext(ctx).Exec("TRUNCATE TABLE `images`")
	h.DB().WithContext(ctx).Exec("TRUNCATE TABLE `user_rooms`")

	h.DB().Exec("SET FOREIGN_KEY_CHECKS=1")
	return nil
}
