package clawrer

import "time"

// Clawrer _
type Clawrer struct {
}

type (
	// Option is clawrer options
	Option struct {
		Interval  time.Duration
		MaxRepeat int
	}
)
