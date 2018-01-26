package run

import (
	"time"
)

var waiter = NewWaiter(time.Second, time.Minute)

func NewWaiter(tick time.Duration, duration time.Duration) *Waiter {
	return &Waiter{
		tick:     tick,
		duration: duration,
	}
}

type Waiter struct {
	tick     time.Duration
	duration time.Duration
}

func (w *Waiter) UntilError(f func() (interface{}, error)) (interface{}, error) {
	ticker := time.NewTicker(w.tick)

	deadline := time.Now().Add(w.duration)

	var err error
	var res interface{}

	for now := range ticker.C {

		if now.After(deadline) {
			break
		}

		res, err = f()

		if err == nil {
			return res, nil
		}
	}

	ticker.Stop()

	return nil, err
}

func UntilError(f func() (interface{}, error)) (interface{}, error) {
	return waiter.UntilError(f)
}
