package goretry

import "time"

/* Fibonacci performs retry function with the waiting duration calculated using Fibonacci sequence.
It will wait for firstWaiting to do the first retry*/
func Fibonacci(firstWaiting time.Duration, action func() error) {
	var backoff time.Duration
	lastWaiting := firstWaiting
	secondLastWaiting := time.Duration(0)
	for {
		if err := action(); err == nil {
			return
		}
		backoff = lastWaiting + secondLastWaiting
		secondLastWaiting = lastWaiting
		lastWaiting = backoff
		time.Sleep(backoff)
	}
}
