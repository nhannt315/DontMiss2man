package datetime

import (
	"database/sql"
	"database/sql/driver"
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
