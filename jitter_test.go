package goretry

import (
	"testing"
	"time"
)

func Test_calculateJitter(t *testing.T) {
	for duration := time.Duration(0); duration <= 10; duration++ {
		for min := duration; min >= 0; min-- {
			for i := 0; i < 100; i++ {
				got := calculateJitter(time.Duration(duration*time.Second), time.Duration(min*time.Second))
				if got < time.Duration(min*time.Second) || got > time.Duration(duration*time.Second) {
					t.Errorf("[index=%d] calculateJitter() Min: %d, Max: %d, got: %d", i, min, duration, got)
				}
			}
		}
	}
}
