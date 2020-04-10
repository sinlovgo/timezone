package timezone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareEQ(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareLocationEQ, err := CompareEQ(timeFrom, timeTo, "2006-01-02 15:04:05")
	if err != nil {
		t.Errorf("CompareEQ err %v", err)
	}
	t.Logf("CompareEQ %v", compareLocationEQ)
	assert.False(t, compareLocationEQ)

	timeFrom2 := "2018-07-11 15:07:51"
	timeTo2 := "2018-07-11 15:07:51.000000"
	locationEQ, err := CompareEQ(timeFrom2, timeTo2, "2006-01-02 15:04:05")
	if err != nil {
		t.Errorf("CompareEQ err %v", err)
	}
	assert.True(t, locationEQ)
}

func TestCompareMicroEQ(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareUTCSecondEQ, err := CompareSecondEQ(timeFrom, timeTo)
	if err != nil {
		t.Errorf("CompareMicroEQ err %v", err)
	}
	t.Logf("CompareMicroEQ %v", compareUTCSecondEQ)
	assert.False(t, compareUTCSecondEQ)

	timeFrom2 := "2018-07-11 15:07:51"
	timeTo2 := "2018-07-11 15:07:51.000000"
	utcSecondEQ, err := CompareMicroEQ(timeFrom2, timeTo2)
	if err != nil {
		t.Errorf("CompareMicroEQ err %v", err)
	}
	assert.True(t, utcSecondEQ)
}

func TestCompareLT(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareLocation, err := CompareLT(timeFrom, timeTo, LayoutSecond)
	if err != nil {
		t.Errorf("CompareLT err %v", err)
	}
	t.Logf("CompareLT %v", compareLocation)
	assert.True(t, compareLocation)

	timeFrom2 := "2018-07-11 15:07:51"
	timeTo2 := "2018-07-11 15:07:51.000100"
	location, err := CompareLT(timeFrom2, timeTo2, LayoutMicro)
	if err != nil {
		t.Errorf("CompareLT err %v", err)
	}
	assert.True(t, location)

	compareLTErr, err := CompareLT(timeFrom2, timeTo2, "error format")
	if err != nil {
		t.Logf("CompareLT err %v", compareLTErr)
	}
	assert.NotEmpty(t, err)
}

func TestCompareSecondLT(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compare, err := CompareSecondLT(timeFrom, timeTo)
	if err != nil {
		t.Errorf("CompareSecondLT err %v", err)
	}
	t.Logf("CompareSecondLT %v", compare)
	assert.True(t, compare)
}

func TestCompareMicroLT(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51.000000"
	timeTo := "2018-07-11 15:07:51.000100"
	compare, err := CompareMicroLT(timeFrom, timeTo)
	if err != nil {
		t.Errorf("CompareMicroLT err %v", err)
	}
	t.Logf("CompareMicroLT %v", compare)
	assert.True(t, compare)
}

func TestCompareGT(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareLocation, err := CompareGT(timeFrom, timeTo, LayoutSecond)
	if err != nil {
		t.Errorf("CompareGT err %v", err)
	}
	t.Logf("CompareGT %v", compareLocation)
	assert.False(t, compareLocation)

	timeFrom2 := "2018-07-11 15:07:51"
	timeTo2 := "2018-07-11 15:07:51.000100"
	location, err := CompareGT(timeFrom2, timeTo2, LayoutMicro)
	if err != nil {
		t.Errorf("CompareGT err %v", err)
	}
	assert.False(t, location)

	compareLTErr, err := CompareGT(timeFrom2, timeTo2, "error format")
	if err != nil {
		t.Logf("CompareGT err %v", compareLTErr)
	}
	assert.NotEmpty(t, err)
}

func TestCompareSecondGT(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compare, err := CompareSecondGT(timeFrom, timeTo)
	if err != nil {
		t.Errorf("CompareSecondGT err %v", err)
	}
	t.Logf("CompareSecondGT %v", compare)
	assert.False(t, compare)
}

func TestCompareMicroGT(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51.000000"
	timeTo := "2018-07-11 15:07:51.000100"
	compare, err := CompareMicroGT(timeFrom, timeTo)
	if err != nil {
		t.Errorf("CompareMicroGT err %v", err)
	}
	t.Logf("CompareMicroGT %v", compare)
	assert.False(t, compare)
}
