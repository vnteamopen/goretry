package goretry

import "time"

/* Fibonacci performs retry function with the waiting duration calculated using Fibonacci sequence.
It will wait for firstWaiting to do the first retry*/
func Fibonacci(startWaiting time.Duration, action func() error) {
	std.Fibonacci(startWaiting, action)
}

/* Fibonacci performs retry function with the waiting duration calculated using Fibonacci sequence.
It will wait for firstWaiting to do the first retry*/
func (i *Instance) Fibonacci(startWaiting time.Duration, action func() error) {
	var count int64
	var totalWaiting time.Duration

	lastWaiting := time.Duration(0)
	secondLastWaiting := startWaiting
	var backoff time.Duration
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

		backoff = i.sleep(lastWaiting + secondLastWaiting)
		totalWaiting += backoff

		secondLastWaiting = lastWaiting
		lastWaiting = backoff
	}
}
