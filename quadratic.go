package goretry

import "time"

/* Quadratic performs retry function with backoff time calculated by quadratic function.
It will wait for baseTime to do the first retry, and then increase the waiting time by time = attempt ^ 2 * baseTime */
func Quadratic(baseTime time.Duration, action func() error) {
	attempt := 0
	var backoff time.Duration
	for {
		if err := action(); err == nil {
			return
		}
		attempt += 1
		backoff = baseTime * time.Duration(attempt*attempt)
		time.Sleep(backoff)
	}
}
