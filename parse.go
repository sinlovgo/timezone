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
