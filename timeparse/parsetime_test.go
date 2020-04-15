package timeparse

import (
	"github.com/sinlovgo/timezone"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocation(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	parseLocation, err := Location(timezone.LayoutISO8601TimeSecond, timezone.LayoutISO8601TimeMicro,
		testTime,
		timezone.ZoneUTC, timezone.ZoneAsiaShanghai)
	if err != nil {
		t.Errorf("Location err: %v", err)
	}
	t.Logf("parseLocation -> %v", parseLocation)
	assert.NotEmpty(t, parseLocation)
}

func TestLocationISO8601TimeSecond(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	locationISO8601TimeSecond, err := LocationISO8601TimeSecond(testTime, "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("LocationISO8601TimeSecond err: %v", err)
	}
	t.Logf("LocationISO8601TimeSecond -> %v", locationISO8601TimeSecond)
	assert.NotEmpty(t, locationISO8601TimeSecond)
}

func TestLocationISO8601TimeSecondUTC(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	locationISO8601TimeSecondUTC, err := LocationISO8601TimeSecondUTC(testTime)
	if err != nil {
		t.Errorf("LocationISO8601TimeSecondUTC err: %v", err)
	}
	t.Logf("LocationISO8601TimeSecondUTC -> %v", locationISO8601TimeSecondUTC)
	assert.NotEmpty(t, locationISO8601TimeSecondUTC)
}

func TestLocationSecond(t *testing.T) {
	parseLocationSecond, err := LocationSecond("2020-02-26 10:08:58", "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("LocationSecond err: %v", err)
	}
	t.Logf("parseLocationSecond -> %v", parseLocationSecond)
	assert.NotEmpty(t, parseLocationSecond)

	errorTimeString := "2020:02#26 10:08:58"
	locationSecondErr, err := LocationSecond(errorTimeString, "UTC", "Asia/Shanghai")
	if err == nil {
		t.Errorf("LocationSecond not return error")
	}
	assert.Empty(t, locationSecondErr)
}

func TestLocationMicro(t *testing.T) {
	parseLocationMicro, err := LocationMicro("2018-01-02 15:04:05", "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("LocationMicro err: %v", parseLocationMicro)
	}
	t.Logf("parseLocationMicro -> %v", parseLocationMicro)
	assert.NotEmpty(t, parseLocationMicro)

	errorTimeString := "2020:02#26 10:08:58"
	locationMicroErr, err := LocationMicro(errorTimeString, "UTC", "Asia/Shanghai")
	if err == nil {
		t.Errorf("LocationMicro not return error")
	}
	assert.Empty(t, locationMicroErr)
}

func TestLocationFix(t *testing.T) {
	testTime := "2020-02-25T15:52:18Z00:00"
	parseLocation, err := LocationFix(timezone.LayoutISO8601TimeUTC, timezone.LayoutSecond, testTime, timezone.ZoneUTC, 0, timezone.ZoneAsiaShanghai, 8)
	if err != nil {
		t.Errorf("LocationFix err: %v", err)
	}
	t.Logf("LocationFix -> %v", parseLocation)
	assert.Equal(t, "2020-02-25 23:52:18", parseLocation)
}

func TestLocationFixISO8601TimeSecond(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	locationISO8601TimeSecond, err := LocationFixISO8601TimeSecond(testTime, "UTC", 0, "Asia/Shanghai", 8)
	if err != nil {
		t.Errorf("LocationFixISO8601TimeSecond err: %v", err)
	}
	t.Logf("LocationFixISO8601TimeSecond -> %v", locationISO8601TimeSecond)
	assert.NotEmpty(t, locationISO8601TimeSecond)
}

func TestLocationFixISO8601TimeSecondUTC(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	locationISO8601TimeSecondUTC, err := LocationFixISO8601TimeSecondUTC(testTime)
	if err != nil {
		t.Errorf("LocationFixISO8601TimeSecondUTC err: %v", err)
	}
	t.Logf("LocationFixISO8601TimeSecondUTC -> %v", locationISO8601TimeSecondUTC)
	assert.NotEmpty(t, locationISO8601TimeSecondUTC)
}

func TestLocationFixSecond(t *testing.T) {
	parseLocationSecond, err := LocationFixSecond("2020-02-26 10:08:58", "UTC", 0, "Asia/Shanghai", 8)
	if err != nil {
		t.Errorf("LocationFixSecond err: %v", err)
	}
	t.Logf("LocationFixSecond -> %v", parseLocationSecond)
	assert.NotEmpty(t, parseLocationSecond)

	errorTimeString := "2020:02#26 10:08:58"
	locationSecondErr, err := LocationFixSecond(errorTimeString, "UTC", 0, "Asia/Shanghai", 8)
	if err == nil {
		t.Errorf("LocationFixSecond not return error")
	}
	assert.Empty(t, locationSecondErr)
}

func TestLocationFixMicro(t *testing.T) {
	parseLocationMicro, err := LocationFixMicro("2020-02-26 10:08:58", "UTC", 0, "Asia/Shanghai", 8)
	if err != nil {
		t.Errorf("LocationFixMicro err: %v", err)
	}
	t.Logf("LocationFixMicro -> %v", parseLocationMicro)
	assert.NotEmpty(t, parseLocationMicro)

	errorTimeString := "2020:02#26 10:08:58"
	locationMicroErr, err := LocationFixMicro(errorTimeString, "UTC", 0, "Asia/Shanghai", 8)
	if err == nil {
		t.Errorf("LocationFixMicro not return error")
	}
	t.Logf("locationMicroErr err -> %v , result %v", err, locationMicroErr)
	assert.NotEmpty(t, err)
}
