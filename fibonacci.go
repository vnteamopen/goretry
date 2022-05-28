package goretry

import "time"

/*Do is basic retry function. If action return error, it will retry after constant backoff duration.
If backoff = 0, no waiting duration between action retries, same with NoBackoff().*/
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
