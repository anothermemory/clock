package clock_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/anothermemory/clock"
	"github.com/stretchr/testify/assert"
)

var dummyTime = time.Date(2017, 11, 24, 17, 0, 0, 0, time.Local)
var dummySecondTime = time.Date(2017, 11, 26, 17, 0, 0, 0, time.Local)

func TestNewMock(t *testing.T) {
	c := clock.NewMock(dummyTime)
	assert.Equal(t, dummyTime, c.Now())
}

func TestMockClock_Now_SameValue(t *testing.T) {
	c := clock.NewMock(dummyTime)
	assert.Equal(t, c.Now(), c.Now())
}

func TestMockClock_Now_MultipleValues(t *testing.T) {
	c := clock.NewMock(dummyTime, dummySecondTime)
	assert.Equal(t, dummyTime, c.Now())
	assert.Equal(t, dummySecondTime, c.Now())
	assert.Equal(t, dummyTime, c.Now())
	assert.Equal(t, dummySecondTime, c.Now())
}

func ExampleNewMock() {
	c := clock.NewMock(time.Date(2017, 11, 24, 17, 0, 0, 0, time.UTC))
	fmt.Println(c.Now())
	fmt.Println(c.Now())
	// Output:
	// 2017-11-24 17:00:00 +0000 UTC
	// 2017-11-24 17:00:00 +0000 UTC
}
