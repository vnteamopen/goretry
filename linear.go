package goretry

import "time"

/* Linear performs retry function with linearly incremental backoff time.
It will wait for firstWaiting to do the first retry, and then increase the waiting time by incrementStep for each next retry*/
func Linear(firstWaiting, incrementStep time.Duration, action func() error) {
	lastWaiting := firstWaiting
	for {
		if err := action(); err == nil {
			return
		}
		lastWaiting = lastWaiting + incrementStep
		time.Sleep(lastWaiting)
	}
}
