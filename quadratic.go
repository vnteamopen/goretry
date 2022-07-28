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
	var count int64
	var totalWaiting time.Duration
	backoff := baseTime

	for {
		i.log("do action()")
		if err := action(); err == nil {
			return
		}
		count++
		if i.MaxStopRetries != NoLimit && count >= i.MaxStopRetries {
			break
		}

		if i.MaxStopTotalWaiting != NoDuration && totalWaiting >= i.MaxStopTotalWaiting {
			break
		}

		backoff = i.sleep(backoff)
		backoff = baseTime * time.Duration((count+1)*(count+1))
		totalWaiting += backoff
	}
}
