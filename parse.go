package timezone

import "time"

// time string to timestamp
//	format -> 2006-01-02 15:04:05.999
//	timeStr -> for parse time string
// when error will return 0 and err
func ParseTimestamp(format, timeStr string) (int64, error) {
	parseInLocation, err := time.ParseInLocation(format, timeStr, zoneLocal)
	if err != nil {
		return 0, err
	}
	return parseInLocation.Unix(), nil
}

// time string must as 2006-01-02 15:04:05 UTC
func ParseTimestampSecond(timeStr string) (int64, error) {
	return ParseTimestamp(LayoutSecond, timeStr)
}

// time string must as 2006-01-02 15:04:05.999999 UTC
func ParesTimestampMicro(timeStr string) (int64, error) {
	return ParseTimestamp(LayoutMicro, timeStr)
}

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
	return ParseLocation(LayoutISO8601TimeSecond, LayoutSecond, timeStr, ZoneUTC, ZoneUTC)
}

func ParseLocationISO8601TimeSecond(timeStr, fromLocation, toLocation string) (string, error) {
	return ParseLocation(LayoutISO8601TimeSecond, LayoutSecond, timeStr, fromLocation, toLocation)
}

// use format 2006-01-02 15:04:05 as ParseLocation()
func ParseLocationSecond(timeStr, fromLocation, toLocation string) (string, error) {
	return ParseLocation(LayoutSecond, LayoutSecond, timeStr, fromLocation, toLocation)
}

// use format 2006-01-02 15:04:05.999999 as ParseLocation()
func ParseLocationMicro(timeStr, fromLocation, toLocation string) (string, error) {
	return ParseLocation(LayoutMicro, LayoutMicro, timeStr, fromLocation, toLocation)
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
	return ParseLocationFix(LayoutISO8601TimeSecond, LayoutSecond, timeStr, ZoneUTC, 0, ZoneUTC, 0)
}

func ParseLocationFixISO8601TimeSecond(timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	return ParseLocationFix(LayoutISO8601TimeSecond, LayoutSecond, timeStr, fromLocation, fromHour, toLocation, toHour)
}

// use format 2006-01-02 15:04:05 as ParseLocation()
func ParseLocationFixSecond(timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	return ParseLocationFix(LayoutSecond, LayoutSecond, timeStr, fromLocation, fromHour, toLocation, toHour)
}

// use format 2006-01-02 15:04:05.999999 as ParseLocation()
func ParseLocationFixMicro(timeStr, fromLocation string, fromHour int, toLocation string, toHour int) (string, error) {
	return ParseLocationFix(LayoutMicro, LayoutMicro, timeStr, fromLocation, fromHour, toLocation, toHour)
}
