package delay

import "time"

// Sleeper - a generic interface for wrapping a sleep function
// So that sleeping can be mocked in tests
type Sleeper interface {
	Sleep(time.Duration)
}

type realSleeper struct{}

func (rc *realSleeper) Sleep(d time.Duration) {
	time.Sleep(d)
}

// NewRealSleeper - returns a new sleeper that uses the real time.Sleep function
func NewRealSleeper() Sleeper {
	return &realSleeper{}
}
