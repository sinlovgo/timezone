package timezone

import "time"

// get time day start
//	day -> 1 tomorrow -1 yesterday
//	layout -> 2006-01-02 15:04:05
func StartDay(day int, layout string) string {
	timeNow := time.Now().In(zoneLocal).Format(LayoutDay)
	timeStart, _ := time.Parse(LayoutDay, timeNow)
	return timeStart.
		AddDate(0, 0, day).
		In(zoneLocal).
		Format(layout)
}

// get time day start
// format layout as "2006-01-02 15:04:05"
//	day -> 1 tomorrow -1 yesterday
func StartDaySecond(day int) string {
	return StartDay(day, LayoutSecond)
}

// get time day start
// format layout as "2006-01-02 15:04:05.999999"
//	day -> 1 tomorrow -1 yesterday
func StartDayMicro(day int) string {
	return StartDay(day, LayoutMicro)
}

// get time month start
//	month -> month of year
//	layout -> 2006-01-02 15:04:05
func StartMonth(month int, layout string) string {
	timeNow := time.Now().In(zoneLocal).Format(LayoutMonth)
	timeStart, _ := time.Parse(LayoutMonth, timeNow)
	return timeStart.
		AddDate(0, month, 0).
		In(zoneLocal).
		Format(layout)
}

// get time month start
// format layout as "2006-01-02 15:04:05"
//	month -> month of year
func StartMonthSecond(month int) string {
	return StartMonth(month, LayoutSecond)
}

// get time month start
// format layout as "2006-01-02 15:04:05.999999"
//	month -> month of year
func StartMonthMicro(month int) string {
	return StartMonth(month, LayoutMicro)
}

// get time year start
//	year -> 1 next year -1 last year
//	layout -> 2006-01-02 15:04:05
func StartYear(year int, layout string) string {
	timeNow := time.Now().In(zoneLocal).Format(LayoutYear)
	timeStart, _ := time.Parse(LayoutYear, timeNow)
	return timeStart.
		AddDate(year, 0, 0).
		In(zoneLocal).
		Format(layout)
}

// get time month start
// format layout as "2006-01-02 15:04:05"
//	year -> 1 next year -1 last year
func StartYearSecond(year int) string {
	return StartYear(year, LayoutSecond)
}

// get time month start
// format layout as "2006-01-02 15:04:05.999999"
//	year -> 1 next year -1 last year
func StartYearMicro(year int) string {
	return StartYear(year, LayoutMicro)
}