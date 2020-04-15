package timezone

import "time"

// get location time by duration
//	duration like time.Duration(yourTime) * time.Second
//	layout 2006-01-02 15:04:05
func CalcDuration(duration time.Duration, layout string) string {
	return time.Now().
		Add(duration).
		In(zoneLocal).
		Format(layout)
}

// get time duration, location UTC
//	duration like time.Duration(yourTime) * time.Second
//	layout 2006-01-02 15:04:05
func CalcDurationUTC(duration time.Duration, layout string) string {
	return time.Now().
		Add(duration).
		In(time.FixedZone(ZoneUTC, 0)).
		Format(layout)
}

// get time duration, location UTC format layout as "2006-01-02 15:04:05"
//	duration like time.Duration(yourTime) * time.Second
func CalcDurationSecondUTC(duration time.Duration) string {
	return CalcDurationUTC(duration, LayoutSecond)
}

// get time duration, location UTC format layout as "2006-01-02 15:04:05.999999"
//	duration like time.Duration(yourTime) * time.Second
func CalcDurationMicroUTC(duration time.Duration) string {
	return CalcDurationUTC(duration, LayoutSecond)
}

// get location time
//	year -> year count
//	month -> month of year
//	day -> 1 tomorrow -1 yesterday
//	layout -> 2006-01-02 15:04:05
func CalcDate(year, month, day int, layout string) string {
	return time.Now().
		AddDate(year, month, day).
		In(zoneLocal).
		Format(layout)
}

// get location time
// format layout as "2006-01-02 15:04:05"
//	year -> year count
//	month -> month of year
//	day -> 1 tomorrow -1 yesterday
func CalcDateSecond(year, month, day int) string {
	return CalcDate(year, month, day, LayoutSecond)
}

// get location time
// format layout as "2006-01-02 15:04:05.999999"
//	year -> year count
//	month -> month of year
//	day -> 1 tomorrow -1 yesterday
func CalcDateMicro(year, month, day int) string {
	return CalcDate(year, month, day, LayoutMicro)
}

// get location time day
//	day -> 1 tomorrow -1 yesterday
//	layout -> 2006-01-02 15:04:05
func CalcDay(day int, layout string) string {
	return CalcDate(0, 0, day, layout)
}

// get location time day
// format layout as "2006-01-02 15:04:05"
//	day -> 1 tomorrow -1 yesterday
func CalcDaySecond(day int) string {
	return CalcDate(0, 0, day, LayoutSecond)
}

// get location time day
// format layout as "2006-01-02 15:04:05.999999"
//	day -> 1 tomorrow -1 yesterday
func CalcDayMicro(day int) string {
	return CalcDate(0, 0, day, LayoutMicro)
}

// get location time month
//	month -> month of year
//	layout -> 2006-01-02 15:04:05
func CalcMonth(month int, layout string) string {
	return CalcDate(0, month, 0, layout)
}

// get location time month
// format layout as "2006-01-02 15:04:05"
//	month -> month of year
func CalcMonthSecond(month int) string {
	return CalcDate(0, month, 0, LayoutSecond)
}

// get location time month
// format layout as "2006-01-02 15:04:05.999999"
//	month -> month of year
func CalcMonthMicro(month int) string {
	return CalcDate(0, month, 0, LayoutMicro)
}

// get location time year
//	year -> year count
//	layout -> 2006-01-02 15:04:05
func CalcYear(year int, layout string) string {
	return CalcDate(year, 0, 0, layout)
}

// get location time year
// format layout as "2006-01-02 15:04:05"
//	year -> year count
func CalcYearSecond(year int) string {
	return CalcDate(year, 0, 0, LayoutSecond)
}

// get location time year
// format layout as "2006-01-02 15:04:05.999999"
//	year -> year count
func CalcYearMicro(year int) string {
	return CalcDate(year, 0, 0, LayoutMicro)
}