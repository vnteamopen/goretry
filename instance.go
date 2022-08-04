package goretry

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

const (
	NoLimit    = int64(0)
	NoDuration = time.Duration(0)
)

// Instance defines a goretry instance with their configs.
type Instance struct {
	// MaxStopRetries defines maximum number of retry times. If reach to this number, the retry will be stopped. Default: NoLimit.
	MaxStopRetries int64

	// MaxStopDuration defines maximum total waiting duration of retry times. If total number of waiting duration is reached, the try will be stopped. Default: NoLimit.
	MaxStopTotalWaiting time.Duration

	// CeilingSleep defines max duration waiting duration during increasing. If next increasing is over this value, it keeps this value instead. Default: NoLimit.
	CeilingSleep time.Duration

	// JitterEnabled defines if Jitter is applied when calculating sleep time. Jitter adds or removes different random waiting durations to back off time. Default: false
	JitterEnabled bool

	// JitterFloorSleep is the lower bound of the random function when calculating sleep time with jitter. Default: NoLimit
	JitterFloorSleep time.Duration

	// JitterMinSleep defines the smallest duration of sleep time with jitter. Default: 0
	JitterMinSleep time.Duration

	// Logger defines log output. You can use os.Stdout, file or any writer stream.
	Logger io.Writer
}

// std is standard instance to use in goretry without specific custom Instance
var std = Instance{
	MaxStopRetries:      NoLimit,
	MaxStopTotalWaiting: NoDuration,
	CeilingSleep:        NoDuration,
	JitterEnabled:       false,
	JitterFloorSleep:    0,
	JitterMinSleep:      0,
	Logger:              nil,
}

func (i *Instance) log(format string, a ...any) {
	if i.Logger == nil {
		return
	}
	i.Logger.Write([]byte(fmt.Sprintf(format+"\n", a...)))
}

func (i *Instance) sleep(duration time.Duration) time.Duration {
	if i.JitterEnabled && duration != 0 {
		duration = time.Duration(rand.Int63n(int64(duration)) + int64(i.JitterFloorSleep))
		if duration < i.JitterMinSleep {
			duration = i.JitterMinSleep
		}
	}
	if i.CeilingSleep != NoDuration && duration > i.CeilingSleep {
		duration = i.CeilingSleep
	}
	time.Sleep(duration)
	i.log("sleep %v", duration)
	return duration
}
