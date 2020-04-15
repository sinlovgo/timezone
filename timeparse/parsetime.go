package timeparse

import (
	"github.com/sinlovgo/timezone"
	"time"
)

// Parse location from location to next
//	formLayout -> 2006-01-02 15:04:05.999999
//	toLayout -> 2006-01-02 15:04:05
//	timeStr -> for parse time string
//	fromLocation -> parse from like UTC
//	toLocation   -> parse to   like Asia/Shanghai
func Location(formLayout, toLayout, timeStr, fromLocation, toLocation string) (string, error) {
	fromLoc, err := time.LoadLocation(fromLocation)
	if err != nil {
		return "", err
	}
	toLoc, err := time.LoadLocation(toLocation)
	if err != nil {
		return "", err
	}
	parse, err := time.ParseInLocation(formLayout, timeStr, fromLoc)
	if err != nil {
		return "", err
	}
	return parse.In(toLoc).Format(toLayout), nil
}

func LocationISO8601TimeSecondUTC(timeStr string) (string, error) {
	return Location(timezone.LayoutISO8601TimeSecond, timezone.LayoutSecond, timeStr, timezone.ZoneUTC, timezone.ZoneUTC)
}

func LocationISO8601TimeSecond(timeStr, fromLocation, toLocation string) (string, error) {
	return Location(timezone.LayoutISO8601TimeSecond, timezone.LayoutSecond, timeStr, fromLocation, toLocation)
}

// use format 2006-01-02 15:04:05 as Location()
func LocationSecond(timeStr, fromLocation, toLocation string) (string, error) {
	return Location(timezone.LayoutSecond, timezone.LayoutSecond, timeStr, fromLocation, toLocation)
}

// use format 2006-01-02 15:04:05.999999 as Location()
func LocationMicro(timeStr, fromLocation, toLocation string) (string, error) {
	return Location(timezone.LayoutMicro, timezone.LayoutMicro, timeStr, fromLocation, toLocation)
}

// Parse location from location to next
//	formLayout -> 2006-01-02 15:04:05.999999
//	toLayout -> 2006-01-02 15:04:05
//	timeStr -> for parse time string
//	fromLocation -> parse from like UTC
//	fromHour     -> parse from time offset +8 -3
//	toLocation   -> parse to   like Asia/Shanghai
func LocationFix(formLayout, toLayout, timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	fromLoc := time.FixedZone(fromLocation, fromHour*3600)
	toLoc := time.FixedZone(toLocation, toHour*3600)
	parse, err := time.ParseInLocation(formLayout, timeStr, fromLoc)
	if err != nil {
		return "", err
	}
	return parse.In(toLoc).Format(toLayout), nil
}

func LocationFixISO8601TimeSecondUTC(timeStr string) (string, error) {
	return LocationFix(timezone.LayoutISO8601TimeSecond, timezone.LayoutSecond, timeStr, timezone.ZoneUTC, 0, timezone.ZoneUTC, 0)
}

func LocationFixISO8601TimeSecond(timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	return LocationFix(timezone.LayoutISO8601TimeSecond, timezone.LayoutSecond, timeStr, fromLocation, fromHour, toLocation, toHour)
}

// use format 2006-01-02 15:04:05 as Location()
func LocationFixSecond(timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	return LocationFix(timezone.LayoutSecond, timezone.LayoutSecond, timeStr, fromLocation, fromHour, toLocation, toHour)
}

// use format 2006-01-02 15:04:05.999999 as Location()
func LocationFixMicro(timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	return LocationFix(timezone.LayoutMicro, timezone.LayoutMicro, timeStr, fromLocation, fromHour, toLocation, toHour)
}
