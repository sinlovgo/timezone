package timezone

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCalcDuration(t *testing.T) {
	calcDurationSecond := time.Duration(-30) * time.Minute
	calcDuration := CalcDuration(calcDurationSecond, LayoutMicro)
	t.Logf("CalcDuration: %v", calcDuration)
	assert.NotEmpty(t, calcDuration)
}

func TestCalcDurationSecondUTC(t *testing.T) {
	calcDurationSecondUTC := CalcDurationSecondUTC(time.Duration(-30) * time.Minute)
	t.Logf("calcDurationSecondUTC: %v", calcDurationSecondUTC)
	assert.NotEmpty(t, calcDurationSecondUTC)
}

func TestCalcDurationMicroUTC(t *testing.T) {
	calcDurationMicroUTC := CalcDurationMicroUTC(time.Duration(30) * time.Minute)
	t.Logf("calcDurationMicroUTC: %v", calcDurationMicroUTC)
	assert.NotEmpty(t, calcDurationMicroUTC)
}

func TestCalcDate(t *testing.T) {
	calcDate := CalcDate(1, 2, 13, LayoutSecond)
	t.Logf("calcDate: %v", calcDate)
	assert.NotEmpty(t, calcDate)
}

func TestCalcDateSecond(t *testing.T) {
	calcDateSecond := CalcDateSecond(2, 3, 10)
	t.Logf("calcDateSecond: %v", calcDateSecond)
	assert.NotEmpty(t, calcDateSecond)
}

func TestCalcDateMicro(t *testing.T) {
	calcDateMicro := CalcDateMicro(2, 2, 2)
	t.Logf("calcDateMicro: %v", calcDateMicro)
	assert.NotEmpty(t, calcDateMicro)
}

func TestCalcDay(t *testing.T) {
	calcDay := CalcDay(10, LayoutSecond)
	t.Logf("calcDay: %v", calcDay)
	assert.NotEmpty(t, calcDay)
}

func TestCalcDaySecond(t *testing.T) {
	calcDaySecond := CalcDaySecond(1)
	t.Logf("calcDaySecond: %v", calcDaySecond)
	assert.NotEmpty(t, calcDaySecond)
	calcDayMicro := CalcDayMicro(2)
	t.Logf("calcDayMicro: %v", calcDayMicro)
	assert.NotEmpty(t, calcDayMicro)
}

func TestCalcDayMicro(t *testing.T) {
	calcDayMicro := CalcDayMicro(20)
	t.Logf("calcDayMicro: %v", calcDayMicro)
	assert.NotEmpty(t, calcDayMicro)
}

func TestCalcMonth(t *testing.T) {
	calcMonth := CalcMonth(2, LayoutSecond)
	t.Logf("calcMonth: %v", calcMonth)
	assert.NotEmpty(t, calcMonth)
}

func TestCalcMonthSecond(t *testing.T) {
	calcMonthSecond := CalcMonthSecond(-1)
	t.Logf("calcMonthSecond: %v", calcMonthSecond)
	assert.NotEmpty(t, calcMonthSecond)
}

func TestCalcMonthMicro(t *testing.T) {
	calcMonthMicro := CalcMonthMicro(-2)
	t.Logf("calcMonthMicro: %v", calcMonthMicro)
	assert.NotEmpty(t, calcMonthMicro)
}

func TestCalcYear(t *testing.T) {
	calcYear := CalcYear(10, LayoutSecond)
	t.Logf("calcYear: %v", calcYear)
	assert.NotEmpty(t, calcYear)
}

func TestCalcYearSecond(t *testing.T) {
	calcYearSecond := CalcYearSecond(-1)
	t.Logf("calcYearSecond: %v", calcYearSecond)
	assert.NotEmpty(t, calcYearSecond)
}

func TestCalcYearMicro(t *testing.T) {
	calcYearMicro := CalcYearMicro(-2)
	t.Logf("calcYearMicro: %v", calcYearMicro)
	assert.NotEmpty(t, calcYearMicro)
}
