package clock

import "time"

type realClock struct {
}

func (realClock) Now() time.Time {
	return time.Now()
}

// NewReal returns Clock which will return actual real time
func NewReal() Clock {
	return &realClock{}
}
