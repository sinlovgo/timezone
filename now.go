package timezone

import "time"

// format layout as 2006-01-02 15:04:05
func UTCSecond() string {
	return time.Now().In(time.FixedZone(ZoneUTC, 0)).Format(LayoutSecond)
}

// format layout as 2006-01-02 15:04:05.999999
func UTCMicro() string {
	return time.Now().In(time.FixedZone(ZoneUTC, 0)).Format(LayoutMicro)
}

// format layout as 2006-01-02 15:04:05
// get location time format by SetZoneByName() or SetZoneFix()
func Second() string {
	return time.Now().In(zoneLocal).Format(LayoutSecond)
}

// format layout as 2006-01-02 15:04:05.999999
// get location time format by SetZoneByName() or SetZoneFix()
func Micro() string {
	return time.Now().In(zoneLocal).Format(LayoutMicro)
}