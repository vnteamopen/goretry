package goretry

import "time"

/* Linear performs retry function with linearly incremental backoff time.
It will wait for firstWaiting to do the first retry, and then increase the waiting time by incrementStep for each next retry*/
func Linear(backoff, incrementStep time.Duration, action func() error) {
	std.Linear(backoff, incrementStep, action)
}

/* Linear performs retry function with linearly incremental backoff time.
It will wait for firstWaiting to do the first retry, and then increase the waiting time by incrementStep for each next retry*/
func (i *Instance) Linear(backoff, incrementStep time.Duration, action func() error) {
	var count int64
	var totalWaiting time.Duration

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
		backoff += incrementStep
		totalWaiting += backoff
	}
}
