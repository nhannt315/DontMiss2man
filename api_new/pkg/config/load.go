package config

import (
	"github.com/nhannt315/real_estate_api/pkg/errors"

	"github.com/jinzhu/configor"
)

// Load 設定ファイルと環境変数をロードし、structにマッピングします。
// また、confのフィールドにmapがある場合、${map_key}_${"env" タグの値} をキーとして、map valueにos.Envの値をセットします。
func Load(conf interface{}, filePath string) error {
	if err := loadFromFile(conf, filePath); err != nil {
		return errors.Wrapf(err, "cannot read conf file: %s", filePath)
	}
	return loadFromEnv(conf)
}

func loadFromFile(conf interface{}, filePath string) error {
	if filePath == "" {
		return configor.Load(conf)
	}

	return configor.Load(conf, filePath)
}
