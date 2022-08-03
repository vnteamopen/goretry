package goretry

import "time"

/* Quadratic performs retry function with backoff time calculated by quadratic function.
It will wait for baseTime to do the first retry, and then increase the waiting time by time = attempt ^ 2 * baseTime */
func Quadratic(baseTime time.Duration, action func() error) {
	std.Quadratic(baseTime, action)
}

/* Quadratic performs retry function with backoff time calculated by quadratic function.
It will wait for baseTime to do the first retry, and then increase the waiting time by time = attempt ^ 2 * baseTime */
func (i *Instance) Quadratic(baseTime time.Duration, action func() error) {
	i.Polynomial(baseTime, 2, action)
}
