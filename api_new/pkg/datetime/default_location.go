package datetime

import (
	"sync/atomic"
	"time"

	"github.com/nhannt315/real_estate_api/pkg/errors"
)

type location struct {
	loc        *time.Location
	alreadySet bool
}

// appLoc/dbLoc stores location for this package.
// Use getAppLocation()/getLocLocation() function to get location value
// instead of reference to appLoc/dbLoc directly.
var (
	appLoc atomic.Value
	dbLoc  atomic.Value
)

func init() {
	appLoc.Store(location{
		loc: time.UTC,
	})
	dbLoc.Store(location{
		loc: time.UTC,
	})
}

// SetAppLocation sets the location to specify which location will be used
// when converting DB's Time/Datetime/Timestamp to Date/Timestamp of this package.
// UTC will be used as default.
// Returns error if invoked more than once.
func SetAppLocation(loc *time.Location) error {
	return setLoc(&appLoc, loc)
}

// SetDBLocation sets the location to specify which location will be used
// when converting Date/Timestamp of this package to DB's Time/Datetime/Timestamp.
// UTC will be used as default.
// Returns error if invoked more than once.
func SetDBLocation(loc *time.Location) error {
	return setLoc(&dbLoc, loc)
}

// SetAppLocationForTest sets location to this package
// without checking already set or not.
// Only use this function for testing purpose.
func SetAppLocationForTest(loc *time.Location) {
	appLoc.Store(location{
		loc: loc,
	})
}

// SetDBLocationForTest sets location to this package
// without checking already set or not.
// Only use this function for testing purpose.
func SetDBLocationForTest(loc *time.Location) {
	dbLoc.Store(location{
		loc: loc,
	})
}

func setLoc(v *atomic.Value, loc *time.Location) error {
	if loc == nil {
		return errors.New("nil location was given")
	}
	if l := v.Load().(location); l.alreadySet {
		return errors.New("location is already set")
	}

	v.Store(location{
		loc:        loc,
		alreadySet: true,
	})
	return nil
}

// getAppLocation returns the registered app location.
func getAppLocation() *time.Location {
	return appLoc.Load().(location).loc
}

// getDBLocation returns the registered db location.
// nolint:unused, deadcode
func getDBLocation() *time.Location {
	return dbLoc.Load().(location).loc
}
