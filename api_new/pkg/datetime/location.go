package datetime

import (
	"strings"
	"time"
)

// LocationUTC is constant for utc location.
const LocationUTC = "UTC"

// LoadLocation returns location by given location string.
func LoadLocation(locStr string) (*time.Location, error) {
	if strings.ToUpper(locStr) == LocationUTC {
		return time.UTC, nil
	}

	return time.LoadLocation(locStr)
}
