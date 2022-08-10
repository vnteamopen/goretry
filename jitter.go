package goretry

import (
	"math/rand"
	"time"
)

// getJitter calculate jitter duration based on input duration. It's random value between [min:input)
func calculateJitter(duration, min time.Duration) time.Duration {
	if duration == 0 {
		return duration
	}
	if duration <= min {
		return duration
	}
	return time.Duration(rand.Int63n(int64(duration)-int64(min)) + int64(min))
}
