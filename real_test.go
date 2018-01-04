package clock_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/anothermemory/clock"
	"github.com/stretchr/testify/assert"
)

func TestNewReal(t *testing.T) {
	c := clock.NewReal()
	assert.NotNil(t, c.Now())
}

func TestRealClock_Now(t *testing.T) {
	c := clock.NewReal()
	t1 := c.Now()
	time.Sleep(time.Second)
	t2 := c.Now()
	assert.NotEqual(t, t1, t2)
}

func ExampleNewReal() {
	c := clock.NewReal()
	fmt.Println(c.Now())
}
