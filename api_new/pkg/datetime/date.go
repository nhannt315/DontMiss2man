package datetime

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/nhannt315/real_estate_api/pkg/errors"
)

const defaultLayout = "2006-01-02"

// Date 日付
//	業務ロジックとして便利な func をサポートする。
//	DBカラムとして利用できる。( database/sql/Scanner, database/sql/driver/Valuer の 実装 )
type Date struct {
	value time.Time
}

// IsValid Check if date is valid
func (m *Date) IsValid() bool {
	return !m.value.IsZero()
}

// Year returns year of Date
func (m *Date) Year() int {
	return m.value.Year()
}

// Month returns month of Date
func (m *Date) Month() time.Month {
	return m.value.Month()
}

// Day returns Day of Date
func (m *Date) Day() int {
	return m.value.Day()
}

// Date returns Year, Month, Day
func (m *Date) Date() (int, time.Month, int) {
	return m.value.Date()
}

// Scan implements sql.Scanner interface
// to scan value from db to Date.
func (m *Date) Scan(value interface{}) error {
	var wrapper sql.NullString
	if err := wrapper.Scan(value); err != nil {
		return err
	}
	var year, month, day int

	if wrapper.Valid {
		// wrapper.String in format of time.RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		if len(wrapper.String) < 10 {
			return errors.Errorf("invalid string %s, want length >= 10", wrapper.String)
		}
		dateStr := strings.Split(wrapper.String[0:10], "-")
		if len(dateStr) != 3 {
			return errors.Errorf("invalid date %s, want yyyy-mm-dd", wrapper.String)
		}

		i, err := strconv.Atoi(dateStr[0])
		if err != nil {
			return errors.Wrapf(err, "cannot convert string %s to year in int", wrapper.String)
		}
		year = i

		i, err = strconv.Atoi(dateStr[1])
		if err != nil {
			return errors.Wrapf(err, "cannot convert string %s to month in int", wrapper.String)
		}
		month = i

		i, err = strconv.Atoi(dateStr[2])
		if err != nil {
			return errors.Wrapf(err, "cannot convert string %s to day in int", wrapper.String)
		}
		day = i
	}

	m.value = time.Date(year, time.Month(month), day, 0, 0, 0, 0, getDBLocation()).In(getAppLocation())

	return nil
}

// Value implements sql/types/Valuer interface
// to convert date to sql value in string format.
func (m *Date) Value() (driver.Value, error) {
	if !m.IsValid() {
		return nil, nil
	}

	timeToSave := m.value.In(getDBLocation())

	return fmt.Sprintf("%04d-%02d-%02d", timeToSave.Year(), timeToSave.Month(), timeToSave.Day()), nil
}

// String to string
func (m *Date) String() string {
	return m.FormatToDate()
}

// Format returns format of Date
func (m *Date) Format(layout string) string {
	return m.ToTime().Format(layout)
}

// FormatToDate converts time's format to Date yyyy-mm-dd layout
func (m *Date) FormatToDate() string {
	return m.Format(defaultLayout)
}

// ToTime converts Date to time.Time in UTC location.
func (m *Date) ToTime() time.Time {
	if !m.IsValid() {
		return time.Time{}
	}
	return m.value
}

// ToTimePtr time.Timeのポインタを返す
func (m *Date) ToTimePtr() *time.Time {
	return &m.value
}

// Compare 2つのDateを比較する
// ( 1:対象 > 指定, 0:等しい, -1:対象 < 指定 )
func (m *Date) Compare(t *Date) int {
	if !m.IsValid() {
		if !t.IsValid() {
			return 0
		}
		return -1 // m < t
	}
	if !t.IsValid() {
		return 1 // m > t
	}

	mT, tT := m.ToTime(), t.ToTime()
	if mT.After(tT) {
		return 1 // m > t
	}
	if mT.Before(tT) {
		return -1 // m < t
	}
	return 0 // m == t
}

// RangeDate Calculate range between two dates and return in days
func RangeDate(src *Date, des *Date) int {
	if !src.IsValid() || !des.IsValid() {
		return 0
	}

	differences := src.ToTime().Sub(des.ToTime()).Hours() / 24
	if differences == 0 {
		return 0
	}

	// SrcDateとTargetDateの時刻はいつも00:00:00になっているので、日数を計算するときに1を足さないといけない
	// 例: 2021-01-01:00-00-00 から 2021-01-03:00-00-00までのDurationを計算すると -> 48時間（2日間）になっちゃうのだが、求める結果は3日間
	differences = math.Abs(differences) + 1

	return int(differences)
}

// MarshalJSON json出力時の変換/Marshal(JSON)
func (m Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", m.String())), nil
}

// UnmarshalJSON 入力されたjsonのパース/Unmarshal(JSON)
func (m *Date) UnmarshalJSON(data []byte) error {

	var d time.Time
	err := json.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	m.value = d
	return nil
}
