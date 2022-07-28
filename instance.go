package goretry

import (
	"fmt"
	"io"
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

	// Logger defines log output. You can use os.Stdout, file or any writer stream.
	Logger io.Writer
}

// std is standard instance to use in goretry without specific custom Instance
var std = Instance{
	MaxStopRetries:      NoLimit,
	MaxStopTotalWaiting: NoDuration,
	CeilingSleep:        NoDuration,
	Logger:              nil,
}

func (i *Instance) log(format string, a ...any) {
	if i.Logger == nil {
		return
	}
	i.Logger.Write([]byte(fmt.Sprintf(format+"\n", a...)))
}

func (i *Instance) sleep(duration time.Duration) time.Duration {
	if i.CeilingSleep != NoDuration && duration > i.CeilingSleep {
		duration = i.CeilingSleep
	}
	time.Sleep(duration)
	i.log("sleep %v", duration)
	return duration
}
