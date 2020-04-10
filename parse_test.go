package timezone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTimestamp(t *testing.T) {
	timestamp, err := ParseTimestamp("2006-01-02 15:04:05", "2018-07-11 15:07:51")
	if err != nil {
		t.Errorf("ParseTimestamp err %v", err)
	}
	t.Logf("timestamp %v", timestamp)
	assert.NotZero(t, timestamp)
}

func TestParseTimestampSecond(t *testing.T) {
	timestampSecond, err := ParseTimestampSecond("2018-07-11 15:07:51")
	if err != nil {
		t.Errorf("ParseTimestampSecond err %v", err)
	}
	t.Logf("timestampSecond %v", timestampSecond)
	assert.NotZero(t, timestampSecond)

	errorTimeStr := "2018:07:11 15:07:51"
	str2TimestampSecondErr, err := ParseTimestampSecond(errorTimeStr)
	if err == nil {
		t.Errorf("ParseTimestampSecond not found error at %v", errorTimeStr)
	}
	assert.Zero(t, str2TimestampSecondErr)
}

func TestParesTimestampMicro(t *testing.T) {
	timestampMicro, err := ParesTimestampMicro("2018-07-11 15:07:51.456123")
	if err != nil {
		t.Errorf("ParesTimestampMicro err %v", err)
	}
	t.Logf("timestampMicro -> %v", timestampMicro)
	assert.NotZero(t, timestampMicro)

	errorTimeStr := "2018:07:11 15:07:51.456123"
	str2TimestampMicroErr, err := ParesTimestampMicro(errorTimeStr)
	if err == nil {
		t.Errorf("ParesTimestampMicro not found error at %v", errorTimeStr)
	}
	assert.Zero(t, str2TimestampMicroErr)
}

func TestParseLocation(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	parseLocation, err := ParseLocation("2006-01-02T15:04:05Z", "2006-01-02 15:04:05", testTime, "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("ParseLocation err: %v", err)
	}
	t.Logf("parseLocation -> %v", parseLocation)
	assert.NotEmpty(t, parseLocation)
}

func TestParseLocationISO8601TimeSecond(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	locationISO8601TimeSecond, err := ParseLocationISO8601TimeSecond(testTime, "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("ParseLocationISO8601TimeSecond err: %v", err)
	}
	t.Logf("ParseLocationISO8601TimeSecond -> %v", locationISO8601TimeSecond)
	assert.NotEmpty(t, locationISO8601TimeSecond)
}

func TestParseLocationISO8601TimeSecondUTC(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	locationISO8601TimeSecondUTC, err := ParseLocationISO8601TimeSecondUTC(testTime)
	if err != nil {
		t.Errorf("ParseLocationISO8601TimeSecondUTC err: %v", err)
	}
	t.Logf("ParseLocationISO8601TimeSecondUTC -> %v", locationISO8601TimeSecondUTC)
	assert.NotEmpty(t, locationISO8601TimeSecondUTC)
}

func TestParseLocationSecond(t *testing.T) {
	parseLocationSecond, err := ParseLocationSecond("2020-02-26 10:08:58", "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("ParseLocationSecond err: %v", err)
	}
	t.Logf("parseLocationSecond -> %v", parseLocationSecond)
	assert.NotEmpty(t, parseLocationSecond)

	errorTimeString := "2020:02#26 10:08:58"
	locationSecondErr, err := ParseLocationSecond(errorTimeString, "UTC", "Asia/Shanghai")
	if err == nil {
		t.Errorf("ParseLocationSecond not return error")
	}
	assert.Empty(t, locationSecondErr)
}

func TestParseLocationMicro(t *testing.T) {
	parseLocationMicro, err := ParseLocationMicro("2020-02-26 10:08:58", "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("ParseLocationMicro err: %v", parseLocationMicro)
	}
	t.Logf("parseLocationMicro -> %v", parseLocationMicro)
	assert.NotEmpty(t, parseLocationMicro)

	errorTimeString := "2020:02#26 10:08:58"
	locationMicroErr, err := ParseLocationMicro(errorTimeString, "UTC", "Asia/Shanghai")
	if err == nil {
		t.Errorf("ParseLocationMicro not return error")
	}
	assert.Empty(t, locationMicroErr)
}

func TestParseLocationFix(t *testing.T) {
	testTime := "2020-02-25T15:52:18Z"
	parseLocation, err := ParseLocationFix("2006-01-02T15:04:05Z", "2006-01-02 15:04:05", testTime, "UTC", 0, "Asia/Shanghai", 8)
	if err != nil {
		t.Errorf("ParseLocationFix err: %v", err)
	}
	t.Logf("ParseLocationFix -> %v", parseLocation)
	assert.Equal(t, "2020-02-25 23:52:18", parseLocation)
}

func TestParseLocationFixISO8601TimeSecond(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	locationISO8601TimeSecond, err := ParseLocationFixISO8601TimeSecond(testTime, "UTC", 0, "Asia/Shanghai", 8)
	if err != nil {
		t.Errorf("ParseLocationFixISO8601TimeSecond err: %v", err)
	}
	t.Logf("ParseLocationFixISO8601TimeSecond -> %v", locationISO8601TimeSecond)
	assert.NotEmpty(t, locationISO8601TimeSecond)
}

func TestParseLocationFixISO8601TimeSecondUTC(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	locationISO8601TimeSecondUTC, err := ParseLocationFixISO8601TimeSecondUTC(testTime)
	if err != nil {
		t.Errorf("ParseLocationFixISO8601TimeSecondUTC err: %v", err)
	}
	t.Logf("ParseLocationFixISO8601TimeSecondUTC -> %v", locationISO8601TimeSecondUTC)
	assert.NotEmpty(t, locationISO8601TimeSecondUTC)
}

func TestParseLocationFixSecond(t *testing.T) {
	parseLocationSecond, err := ParseLocationFixSecond("2020-02-26 10:08:58", "UTC", 0, "Asia/Shanghai", 8)
	if err != nil {
		t.Errorf("ParseLocationFixSecond err: %v", err)
	}
	t.Logf("ParseLocationFixSecond -> %v", parseLocationSecond)
	assert.NotEmpty(t, parseLocationSecond)

	errorTimeString := "2020:02#26 10:08:58"
	locationSecondErr, err := ParseLocationFixSecond(errorTimeString, "UTC", 0, "Asia/Shanghai", 8)
	if err == nil {
		t.Errorf("ParseLocationFixSecond not return error")
	}
	assert.Empty(t, locationSecondErr)
}

func TestParseLocationFixMicro(t *testing.T) {
	parseLocationMicro, err := ParseLocationFixMicro("2020-02-26 10:08:58", "UTC", 0, "Asia/Shanghai", 8)
	if err != nil {
		t.Errorf("ParseLocationFixMicro err: %v", parseLocationMicro)
	}
	t.Logf("ParseLocationFixMicro -> %v", parseLocationMicro)
	assert.NotEmpty(t, parseLocationMicro)

	errorTimeString := "2020:02#26 10:08:58"
	locationMicroErr, err := ParseLocationFixMicro(errorTimeString, "UTC", 0, "Asia/Shanghai", 8)
	if err == nil {
		t.Errorf("ParseLocationFixMicro not return error")
	}
	t.Logf("ParseLocationFixMicro err -> %v", locationMicroErr)
	assert.NotEmpty(t, err)
}
