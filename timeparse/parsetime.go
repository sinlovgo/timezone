package timeparse

import (
	"github.com/sinlovgo/timezone"
	"time"
)

// Parse location from location to next
//	format -> 2006-01-02 15:04:05.999999
//	timeStr -> for parse time string
//	fromLocation -> parse from like UTC
//	toLocation   -> parse to   like Asia/Shanghai
func ParseLocation(formFormat, toFormat, timeStr, fromLocation, toLocation string) (string, error) {
	fromLoc, err := time.LoadLocation(fromLocation)
	if err != nil {
		return "", err
	}
	toLoc, err := time.LoadLocation(toLocation)
	if err != nil {
		return "", err
	}
	parse, err := time.ParseInLocation(formFormat, timeStr, fromLoc)
	if err != nil {
		return "", err
	}
	return parse.In(toLoc).Format(toFormat), nil
}

func ParseLocationISO8601TimeSecondUTC(timeStr string) (string, error) {
	return ParseLocation(timezone.LayoutISO8601TimeSecond, timezone.LayoutSecond, timeStr, timezone.ZoneUTC, timezone.ZoneUTC)
}

func ParseLocationISO8601TimeSecond(timeStr, fromLocation, toLocation string) (string, error) {
	return ParseLocation(timezone.LayoutISO8601TimeSecond, timezone.LayoutSecond, timeStr, fromLocation, toLocation)
}

// use format 2006-01-02 15:04:05 as ParseLocation()
func ParseLocationSecond(timeStr, fromLocation, toLocation string) (string, error) {
	return ParseLocation(timezone.LayoutSecond, timezone.LayoutSecond, timeStr, fromLocation, toLocation)
}

// use format 2006-01-02 15:04:05.999999 as ParseLocation()
func ParseLocationMicro(timeStr, fromLocation, toLocation string) (string, error) {
	return ParseLocation(timezone.LayoutMicro, timezone.LayoutMicro, timeStr, fromLocation, toLocation)
}

// Parse location from location to next
//	format -> 2006-01-02 15:04:05.999999
//	timeStr -> for parse time string
//	fromLocation -> parse from like UTC
//	fromHour     -> parse from time offset +8 -3
//	toLocation   -> parse to   like Asia/Shanghai
func ParseLocationFix(formFormat, toFormat, timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	fromLoc := time.FixedZone(fromLocation, fromHour*3600)
	toLoc := time.FixedZone(toLocation, toHour*3600)
	parse, err := time.ParseInLocation(formFormat, timeStr, fromLoc)
	if err != nil {
		return "", err
	}
	return parse.In(toLoc).Format(toFormat), nil
}

func ParseLocationFixISO8601TimeSecondUTC(timeStr string) (string, error) {
	return ParseLocationFix(timezone.LayoutISO8601TimeSecond, timezone.LayoutSecond, timeStr, timezone.ZoneUTC, 0, timezone.ZoneUTC, 0)
}

func ParseLocationFixISO8601TimeSecond(timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	return ParseLocationFix(timezone.LayoutISO8601TimeSecond, timezone.LayoutSecond, timeStr, fromLocation, fromHour, toLocation, toHour)
}

// use format 2006-01-02 15:04:05 as ParseLocation()
func ParseLocationFixSecond(timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	return ParseLocationFix(timezone.LayoutSecond, timezone.LayoutSecond, timeStr, fromLocation, fromHour, toLocation, toHour)
}

// use format 2006-01-02 15:04:05.999999 as ParseLocation()
func ParseLocationFixMicro(timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	return ParseLocationFix(timezone.LayoutMicro, timezone.LayoutMicro, timeStr, fromLocation, fromHour, toLocation, toHour)
}
