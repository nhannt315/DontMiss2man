package datetime

import "database/sql/driver"

type NullDate struct {
	Date  *Date
	Valid bool
}

func NewNullDate(date *Date) NullDate {
	return NullDate{Date: date, Valid: true}
}

func NewNullDatePtr(date *Date) *NullDate {
	return &NullDate{Date: date, Valid: true}
}

func (d *NullDate) Scan(value interface{}) error {
	if value == nil {
		d.Valid = false
		return nil
	}

	d.Valid = true
	return d.Date.Scan(value)
}

func (d NullDate) Value() (driver.Value, error) {
	if !d.Valid {
		return nil, nil
	}

	return d.Date.Value()
}

func (d *NullDate) UnmarshalJSON(data []byte) error {
	if string(data) == nullString {
		d.Valid = false
		return nil
	}

	d.Valid = true
	return d.UnmarshalJSON(data)
}

func (d *NullDate) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return []byte(nullString), nil
	}

	return d.MarshalJSON()
}
