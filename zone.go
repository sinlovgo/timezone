package timezone

import (
	"sync"
	"time"
)

const (
	// Universal Time Coordinated
	ZoneUTC string = "UTC"
	// Greenwich Mean Time some as UT0, so equal as UTC
	ZoneGMT string = "GMT"
	// Central European Time.
	//	Winter time is UTC +1, and summer European summer time is UTC +2
	ZoneCET string = "CET"
	// Central European Summer Time.
	// warning do not use this in timezone.SetZoneByName
	//	use as Most European countries and some North African countries.
	//	Winter time is UTC +1, and summer European summer time is UTC +2
	ZoneCEST string = "CEST"
	//	Central Standard Time (USA) UTC -6
	//	Central Standard Time (Australia) UTC +9:30
	//	China Standard Time UTC +8
	//	Cuba Standard Time UTC -4
	ZoneCST string = "CST"

	// zone location Asia/Shanghai more see https://en.wikipedia.org/wiki/Tz_database
	ZoneAsiaShanghai string = "Asia/Shanghai"
)

var zoneLocal = time.Local

// mutex
var lock sync.Mutex

// get zone Location
func GetZoneLocal() *time.Location {
	return zoneLocal
}

func SetZoneUTC() {
	lock.Lock()
	zoneLocal = time.FixedZone("UTC", 0)
	lock.Unlock()
}

//	locationName -> UTC CST Asia/Shanghai of others
//	hour : for hour set like 8 Eastern Eight
func SetZoneFix(locationName string, hour int) {
	offset := 3600 * hour
	lock.Lock()
	zoneLocal = time.FixedZone(locationName, offset)
	lock.Unlock()
}

//	locationName -> UTC Asia/Shanghai of others
func SetZoneByName(locationName string) error {
	location, err := time.LoadLocation(locationName)
	if err != nil {
		return err
	}
	lock.Lock()
	zoneLocal = location
	lock.Unlock()
	return nil
}

// present on all systems, especially non-Unix systems
//	most of time use SetZoneFix(locationName string, hour int)
func SetZoneLocation() {
	lock.Lock()
	location, _ := time.LoadLocation("Local")
	zoneLocal = location
	lock.Unlock()
}
