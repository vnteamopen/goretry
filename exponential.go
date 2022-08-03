package goretry

import "time"

/* Exponential performs retry function with backoff time calculated by exponential function.
It will wait for baseTime to do the first retry, and then increase the waiting time by time = 2 ^ attempt * baseTime */
func Exponential(baseTime time.Duration, action func() error) {
	std.Exponential(baseTime, action)
}

/* Exponential performs retry function with backoff time calculated by exponential function.
It will wait for baseTime to do the first retry, and then increase the waiting time by time = 2 ^ attempt * baseTime */
func (i *Instance) Exponential(baseTime time.Duration, action func() error) {
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

		i.sleep(backoff)
		backoff = baseTime * time.Duration(intPow(2, (count+1)))
		totalWaiting += backoff
	}
}
