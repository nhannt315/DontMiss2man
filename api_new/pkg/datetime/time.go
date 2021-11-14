package datetime

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

const iso8601Layout = "2006-01-02T15:04:05Z0700"

// Time 時間
type Time struct {
	value *time.Time
}

func (t *Time) IsValid() bool {
	return t.value != nil
}

func (t *Time) Scan(value interface{}) error {
	var wrapper sql.NullString
	if err := wrapper.Scan(value); err != nil {
		return err
	}

	timeValue, err := time.Parse(iso8601Layout, wrapper.String)
	if err != nil {
		return err
	}

	newTimeValue := timeValue.In(getAppLocation())

	t.value = &newTimeValue

	return nil
}

func (t *Time) Value() (driver.Value, error) {
	if !t.IsValid() {
		return nil, nil
	}

	timeToSave := t.value.In(getDBLocation())

	return timeToSave.Format(iso8601Layout), nil
}

// ToTime converts Date to time.Time in UTC location.
func (t *Time) ToTime() time.Time {
	if !t.IsValid() {
		return time.Time{}
	}
	return *t.value
}

// ToTimePtr time.Timeのポインタを返す
func (t *Time) ToTimePtr() *time.Time {
	return t.value
}

// Format returns format of Date
func (t *Time) Format(layout string) string {
	return t.ToTime().Format(layout)
}

func (t *Time) FormatISO8601() string {
	if !t.IsValid() {
		return ""
	}
	return t.Format(iso8601Layout)
}

func (t *Time) String() string {
	return t.FormatISO8601()
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", t.String())), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var value time.Time
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	t.value = &value
	return nil
}
