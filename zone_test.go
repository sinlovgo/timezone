package timezone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_00_SetZoneByName(t *testing.T) {
	t.Logf("timezone.GetZoneLocal().String() %v", GetZoneLocal().String())
	err := SetZoneByName(ZoneAsiaShanghai)
	if err != nil {
		t.Errorf("SetZoneByName err %v", err)
	}
	t.Logf("now: timezone.GetZoneLocal().String() %v", GetZoneLocal())
	assert.Equal(t, zoneLocal, GetZoneLocal())
	err = SetZoneByName("error")
	t.Logf("must error: %v", err)
	assert.NotEmpty(t, err)
}

func Test_01_SetZoneFix(t *testing.T) {
	t.Logf("timezone.GetZoneLocal().String() %v", GetZoneLocal().String())
	zoneCST := ZoneCST
	SetZoneFix(zoneCST, 8)
	assert.Equal(t, zoneCST, GetZoneLocal().String())
}

func Test_02_SetZoneUTC(t *testing.T) {
	t.Logf("timezone.GetZoneLocal().String() %v", GetZoneLocal().String())
	SetZoneUTC()
	nowZone := GetZoneLocal().String()
	t.Logf("timezone.GetZoneLocal().String() %v", nowZone)
	assert.Equal(t, ZoneUTC, GetZoneLocal().String())
}

func Test_03_SetZoneLocation(t *testing.T) {
	t.Logf("timezone.GetZoneLocal().String() %v", GetZoneLocal().String())
	SetZoneLocation()
	t.Logf("timezone.GetZoneLocal().String() %v", GetZoneLocal().String())
	assert.Equal(t, "Local", GetZoneLocal().String())
}
