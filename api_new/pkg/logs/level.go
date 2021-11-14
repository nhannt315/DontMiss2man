package logs

import (
	"fmt"
	"strings"

	"github.com/nhannt315/real_estate_api/pkg/errors"
)

// Level ログレベル
type Level uint8

var levelSuffix = "level"

const (
	// UnknownLevel ログレベル: 不明
	UnknownLevel Level = iota
	// DebugLevel ログレベル: デバッグ
	DebugLevel
	// InfoLevel ログレベル: 情報
	InfoLevel
	// WarnLevel ログレベル: 警告
	WarnLevel
	// ErrorLevel ログレベル: エラー
	ErrorLevel
)

// MarshalYAML ログレベルをYAML用にマーシャルする。
// 返される文字列は、 末尾のLevelが取り除かれ、全て小文字となる。
func (lv Level) MarshalYAML() (interface{}, error) {
	s := strings.ToLower(lv.String())
	if strings.HasSuffix(s, levelSuffix) {
		s = s[0:(len(s) - len(levelSuffix))]
	}
	return s, nil
}

// UnmarshalYAML  ログレベルをYAML用にアンマーシャルする。
// 大文字/小文字、末尾のLevel有無に関係なく一致するログレベルが設定される。
func (lv *Level) UnmarshalYAML(unmarshal func(interface{}) error) error {
	orgLvStr := ""
	if err := unmarshal(&orgLvStr); err != nil {
		return err
	}
	lvStr := strings.ToLower(orgLvStr)
	if !strings.HasSuffix(lvStr, levelSuffix) {
		lvStr = fmt.Sprintf("%s%s", lvStr, levelSuffix)
	}

	for i := 0; i < (len(_Level_index) - 1); i++ {
		s := _Level_name[_Level_index[i]:_Level_index[i+1]]
		s = strings.ToLower(s)
		if s == lvStr {
			*lv = Level(i)
			return nil
		}
	}
	return errors.Errorf("invalid log level. %s", orgLvStr)
}

// IsErrorEnabled エラー・ログ出力が有効なログレベルか(エラー以下か)
func (lv Level) IsErrorEnabled() bool {
	return lv <= ErrorLevel
}

// IsWarnEnabled 警告・ログ出力が有効なログレベルか(警告以下か)
func (lv Level) IsWarnEnabled() bool {
	return lv <= WarnLevel
}

// IsInfoEnabled 情報・ログ出力が有効なログレベルか(情報以下か)
func (lv Level) IsInfoEnabled() bool {
	return lv <= InfoLevel
}

// IsDebugEnabled デバッグ・ログ出力が有効なログレベルか(デバッグ以下か)
func (lv Level) IsDebugEnabled() bool {
	return lv <= DebugLevel
}
