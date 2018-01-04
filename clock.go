// Package clock contains interface and set of implementations for getting current time. For production usage
// time must be mostly real time. For testing purposes it's often much easier to use mocked time so it will
// return required time each time.
package clock

import "time"

// Clock represents interface which can be used to get current time
type Clock interface {
	// Now returns current time. Actual returned time depends on clock implementation
	Now() time.Time
}
