package test

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

func CreateEmptyDB(db *sql.DB, logger logs.Logger, dbNamePrefix string) (string, error) {
	const retry = 10
	for i := 0; i < retry; i++ {
		name, err := newDB(db, dbNamePrefix)
		if err == nil {
			return name, nil
		}
		logger.Infof(context.Background(), "Error when creating empty database: %+v", err)
	}

	return "", errors.New("Failed to create empty database")
}

func newDB(db *sql.DB, dbNamePrefix string) (string, error) {

	// ランダムな DB 名を決める
	// DB 名をランダムに決めているだけでせいぜい1パッケージごとに1DBなので比較的弱い乱数生成器でも問題ない
	testDBName := fmt.Sprintf("%s_%d", dbNamePrefix, time.Now().UnixNano())
	if _, err := db.Exec("CREATE DATABASE " + testDBName); err != nil {
		return "", errors.Wrap(err, "create DB failed")
	}
	return testDBName, nil
}
