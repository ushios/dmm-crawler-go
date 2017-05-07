package clawrer

import "time"

const (
	// APILengthMax is max size of length
	APILengthMax = 100
)

type (
	// Option is clawrer options
	Option struct {
		Interval  time.Duration
		MaxRepeat int
	}
)
