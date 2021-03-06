package timezone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUTCSecond(t *testing.T) {
	utcSecond := UTCSecond()
	t.Logf("utcSecond %v", utcSecond)
	assert.NotEmpty(t, utcSecond)
}

func TestUTCMicro(t *testing.T) {
	utcMicro := UTCMicro()
	t.Logf("utcMicro %v", utcMicro)
	assert.NotEmpty(t, utcMicro)
}

func TestSecond(t *testing.T) {
	second := Second()
	t.Logf("second %v", second)
	assert.NotEmpty(t, second)
}

func TestMicro(t *testing.T) {
	micro := Micro()
	t.Logf("micro %v", micro)
	assert.NotEmpty(t, micro)
}

func TestTimestampSecond(t *testing.T) {
	timestampSecond := TimestampSecond()
	t.Logf("timestampSecond %v", timestampSecond)
	assert.NotZero(t, timestampSecond)
}

func TestTimestampUnixNano(t *testing.T) {
	timestampUnixNano := TimestampUnixNano()
	t.Logf("timestampUnixNano %v", timestampUnixNano)
	assert.NotZero(t, timestampUnixNano)
}
