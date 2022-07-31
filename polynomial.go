package goretry

import (
	"time"

	"github.com/vnteamopen/goretry/common"
)

/* Polynomial performs retry function with backoff time calculated by Polynomial function.
It will wait for baseTime to do the first retry, and then increase the waiting time by time = attempt ^ degree * baseTime */
func Polynomial(baseTime time.Duration, degree int, action func() error) {
	std.Polynomial(baseTime, degree, action)
}

/* Polynomial performs retry function with backoff time calculated by Polynomial function.
It will wait for baseTime to do the first retry, and then increase the waiting time by time = attempt ^ degree * baseTime */
func (i *Instance) Polynomial(baseTime time.Duration, degree int, action func() error) {
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
		backoff = baseTime * time.Duration(common.IntPow((count+1), degree))
		totalWaiting += backoff
	}
}
