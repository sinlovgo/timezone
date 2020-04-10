package timezone

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup
	SetZoneUTC()
	os.Exit(m.Run())
	// teardown
}
