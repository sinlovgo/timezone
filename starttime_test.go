package timezone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartDaySecond(t *testing.T) {
	startDaySecond := StartDaySecond(-1)
	t.Logf("startDaySecond: %v", startDaySecond)
	assert.NotEmpty(t, startDaySecond)
}

func TestStartDayMicro(t *testing.T) {
	startDayMicro := StartDayMicro(-2)
	t.Logf("startDayMicro: %v", startDayMicro)
	assert.NotEmpty(t, startDayMicro)
}

func TestStartMonthSecond(t *testing.T) {
	startMonthSecond := StartMonthSecond(-1)
	t.Logf("startMonthSecond: %v", startMonthSecond)
	assert.NotEmpty(t, startMonthSecond)
}

func TestStartMonthMicro(t *testing.T) {
	startMonthMicro := StartMonthMicro(-2)
	t.Logf("startMonthMicro: %v", startMonthMicro)
	assert.NotEmpty(t, startMonthMicro)
}

func TestStartYearSecond(t *testing.T) {
	startYearSecond := StartYearSecond(-1)
	t.Logf("startYearSecond: %v", startYearSecond)
	assert.NotEmpty(t, startYearSecond)
}

func TestStartYearMicro(t *testing.T) {
	startYearMicro := StartYearMicro(-2)
	t.Logf("startYearMicro: %v", startYearMicro)
	assert.NotEmpty(t, startYearMicro)
}
