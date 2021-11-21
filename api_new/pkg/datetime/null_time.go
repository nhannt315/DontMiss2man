package datetime

import (
	"database/sql/driver"
)

type NullTime struct {
	Time  *Time
	Valid bool
}

func NewNullTime(t *Time) NullTime {
	return NullTime{Time: t, Valid: true}
}

func NewNullTimePtr(t *Time) *NullTime {
	return &NullTime{Time: t, Valid: true}
}

func (t *NullTime) Scan(value interface{}) error {
	if value == nil {
		t.Valid = false
		return nil
	}

	t.Valid = true
	return t.Time.Scan(value)
}

func (t NullTime) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}

	return t.Time.Value()
}

func (t *NullTime) UnmarshalJSON(data []byte) error {
	if string(data) == nullString {
		t.Valid = false
		return nil
	}

	t.Valid = true
	return t.Time.UnmarshalJSON(data)
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte(nullString), nil
	}

	return t.Time.MarshalJSON()
}
