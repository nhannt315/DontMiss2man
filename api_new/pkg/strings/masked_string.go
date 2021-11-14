package strings

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// MaskedStringText MaskedStringの表示時のテキスト
var MaskedStringText = "******"

// MaskedString マスクが必要なstring
//	文字列として表示時に、実際のstringではなく、マスクされた固定値が使われる。
type MaskedString struct {
	planeString string
}

// MaskedStringWithString stringされたマスクされた文字列の生成
func MaskedStringWithString(s string) MaskedString {
	return MaskedString{s}
}

// PlaneString マスクされていない値
func (m MaskedString) PlaneString() string {
	return m.planeString
}

// String マスクされた値
func (m MaskedString) String() string {
	return MaskedStringText
}

// Format マスクされた値の出力
func (m MaskedString) Format(f fmt.State, c rune) {
	f.Write([]byte(MaskedStringText))
}

// MarshalJSON マスクされた値の出力/Marshal(JSON)
func (m MaskedString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", MaskedStringText)), nil
}

// UnmarshalJSON マスクされた値へのパース/Unmarshal(JSON)
func (m *MaskedString) UnmarshalJSON(data []byte) error {

	var s string
	err := json.Unmarshal(data, &s)
	//nolint:wrapcheck
	if err != nil {
		return err
	}
	m.planeString = s
	return nil
}

// MarshalYAML マスクされた値の出力/Marshal(YAML)
func (m MaskedString) MarshalYAML() (interface{}, error) {
	return MaskedStringText, nil
}

// UnmarshalYAML マスクされた値へのパース/Unmarshal(YAML)
func (m *MaskedString) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}
	m.planeString = s
	return nil
}

// Value driver.Valueの返却(sql packageでMasked Stringを利用する目的)
func (m MaskedString) Value() (driver.Value, error) {
	return driver.Value(m.PlaneString()), nil
}
