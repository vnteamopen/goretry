package goretry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/vnteamopen/goretry"
)

func TestDo(t *testing.T) {
	var counting int64

	start := time.Now()
	goretry.Do(100*time.Millisecond, func() error {
		counting++
		if counting > 5 {
			return nil
		}
		return errors.New("fake error")
	})
	duration := time.Since(start)

	expectedCounting := int64(6)
	if counting != expectedCounting {
		t.Errorf("Do() expected counting: %d, actual: %d", expectedCounting, counting)
	}

	expectedDuration := 500 * time.Millisecond
	if duration < expectedDuration {
		t.Errorf("Do() expected duration: %d, actual: %d", expectedDuration, duration)
	}
}

func TestNoBackoff(t *testing.T) {
	var counting int64

	start := time.Now()
	goretry.NoBackoff(func() error {
		counting++
		if counting > 5 {
			return nil
		}
		return errors.New("fake error")
	})

	duration := time.Since(start)

	expectedCounting := int64(6)
	if counting != expectedCounting {
		t.Errorf("NoBackoff() expected counting: %d, actual: %d", expectedCounting, counting)
	}

	expectedDuration := 100 * time.Millisecond
	if duration > expectedDuration {
		t.Errorf("NoBackoff() expected duration: %d, actual: %d", expectedDuration, duration)
	}
}
