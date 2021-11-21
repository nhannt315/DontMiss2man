package datetime

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

const defaultTimeLayout = time.RFC3339

// Time 時間
type Time struct {
	value time.Time
}

func (t *Time) IsValid() bool {
	return !t.value.IsZero()
}

func (t *Time) Scan(value interface{}) error {
	var wrapper sql.NullString
	if err := wrapper.Scan(value); err != nil {
		return err
	}

	timeValue, err := time.Parse(defaultTimeLayout, wrapper.String)
	if err != nil {
		return err
	}

	t.value = timeValue.In(getAppLocation())

	return nil
}

func (t Time) Value() (driver.Value, error) {
	if !t.IsValid() {
		return nil, nil
	}

	timeToSave := t.value.In(getDBLocation())

	return timeToSave.Format(defaultTimeLayout), nil
}

// ToTime converts Date to time.Time in UTC location.
func (t *Time) ToTime() time.Time {
	if !t.IsValid() {
		return time.Time{}
	}
	return t.value
}

// ToTimePtr time.Timeのポインタを返す
func (t *Time) ToTimePtr() *time.Time {
	return &t.value
}

// Format returns format of Date
func (t *Time) Format(layout string) string {
	return t.ToTime().Format(layout)
}

func (t *Time) FormatISO8601() string {
	if !t.IsValid() {
		return ""
	}
	return t.Format(defaultTimeLayout)
}

func (t *Time) String() string {
	return t.Format(defaultLayout)
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
	t.value = value
	return nil
}
