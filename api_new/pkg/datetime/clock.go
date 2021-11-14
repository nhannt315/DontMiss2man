package datetime

import (
	"time"

	"github.com/nhannt315/real_estate_api/pkg/errors"
)

type Clock interface {
	Now() *Time
	NewTime(year int, month time.Month, day int, hour, min, sec, nsec int) *Time
	ParseTime(layout, value string) (*Time, error)
}

type clock struct{}

func NewClock() Clock {
	return &clock{}
}

func (c *clock) Now() *Time {
	timeValue := time.Now().In(getAppLocation())
	return &Time{value: &timeValue}
}

func (c *clock) NewTime(year int, month time.Month, day int, hour, min, sec, nsec int) *Time {
	timeValue := time.Date(year, month, day, hour, min, sec, nsec, getAppLocation())
	return &Time{value: &timeValue}
}

func (c *clock) ParseTime(layout, value string) (*Time, error) {
	timeValue, err := time.ParseInLocation(layout, value, getAppLocation())
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse time")
	}

	return &Time{value: &timeValue}, nil
}
