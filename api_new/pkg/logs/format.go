package logs

import (
	"strings"

	"github.com/nhannt315/real_estate_api/pkg/errors"
)

// Format is logging format.
type Format int

const (
	// UnknownFormat ログレベル: 不明
	UnknownFormat Format = iota
	// ConsoleFormat コンソール出力ログを出力する。
	ConsoleFormat // console
	// JSONFormat はjsonタイプでログを出力する。
	JSONFormat // json
)

// MarshalYAML ログレベルをYAML用にマーシャルする。
// 返される文字列は、 末尾のLevelが取り除かれ、全て小文字となる。
func (f Format) MarshalYAML() (interface{}, error) {
	return f.String(), nil
}

// UnmarshalYAML  ログレベルをYAML用にアンマーシャルする。
// 大文字/小文字、末尾のLevel有無に関係なく一致するログレベルが設定される。
func (f *Format) UnmarshalYAML(unmarshal func(interface{}) error) error {
	fmtStr := ""
	if err := unmarshal(&fmtStr); err != nil {
		return err
	}
	fmtStr = strings.ToLower(fmtStr)

	for i := 0; i < (len(_Format_index) - 1); i++ {
		s := _Format_name[_Format_index[i]:_Format_index[i+1]]
		if s == fmtStr {
			*f = Format(i)
			return nil
		}
	}
	return errors.Errorf("invalid log level. %s", fmtStr)
}
