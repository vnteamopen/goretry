package goretry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/vnteamopen/goretry"
)

func TestFibonacci(t *testing.T) {
	var counting int64

	start := time.Now()
	goretry.Fibonacci(100*time.Millisecond, func() error {
		counting++
		if counting > 5 {
			return nil
		}
		return errors.New("fake error")
	})
	duration := time.Since(start)

	expectedCounting := int64(6)
	if counting != expectedCounting {
		t.Errorf("Linear() expected counting: %d, actual: %d", expectedCounting, counting)
	}

	expectedDuration := 500 * time.Millisecond
	if duration < expectedDuration {
		t.Errorf("Linear() expected duration: %d, actual: %d", expectedDuration, duration)
	}
}
