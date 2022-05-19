package goretry

import "time"

/*Do is basic retry function. If action return error, it will retry after constant backoff duration.
If backoff = 0, no waiting duration between action retries, same with NoBackoff().*/
func Do(backoff time.Duration, action func() error) {
	for {
		if err := action(); err == nil {
			return
		}
		time.Sleep(backoff)
	}
}

// NoBackoff is a shortcut to call Do(0, action). It will retry immediately and don't wait at all.
func NoBackoff(action func() error) {
	Do(time.Duration(0), action)
}

/* Linear performs retry function with linearly incremental backoff time.
It will wait for firstWaiting to do the first retry, and then increase the waiting time by incrementStep for each next retry*/
func Linear(firstWaiting, incrementStep time.Duration, action func() error) {
	var attempt, backoffTime int64
	for {
		if err := action(); err == nil {
			return
		}

		backoffTime = int64(firstWaiting) + attempt*int64(incrementStep)
		time.Sleep(time.Duration(backoffTime))

		attempt += 1
	}

}
