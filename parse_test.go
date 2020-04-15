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
