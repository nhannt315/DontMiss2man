package datetime

import "time"

type Calendar interface {
	CurrentDate() *Date
	NewDate(year int, month time.Month, day int) *Date
	ParseDate(s string, layout string) (*Date, error)
	NewNullDate(d *Date) *NullDate
}

type calendar struct{}

func NewCalendar() Calendar {
	return &calendar{}
}

// NewCurrentDate returns new Date today
func (c *calendar) CurrentDate() *Date {
	value := time.Now().In(getAppLocation())
	return &Date{value: &value}
}

// NewDate returns new Date from given year, month and day.
// Month and day start from 1.
func (c *calendar) NewDate(year int, month time.Month, day int) *Date {
	newValue := time.Date(year, month, day, 0, 0, 0, 0, getAppLocation())
	return &Date{
		value: &newValue,
	}
}

func (c *calendar) ParseDate(s string, layout string) (*Date, error) {
	if layout == "" {
		layout = defaultLayout
	}

	t, err := time.ParseInLocation(layout, s, getAppLocation())
	if err != nil {
		return nil, err
	}

	return &Date{value: &t}, nil
}

func (c *calendar) NewNullDate(d *Date) *NullDate {
	if d == nil {
		return &NullDate{Valid: false}
	}

	return &NullDate{Date: *d, Valid: true}
}
