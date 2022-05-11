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
	to := time.Duration(0)
	goretry.NoBackoff(func() error {
		s1 := time.Now()
		defer func() {
			to += time.Since(s1)
		}()
		time.Sleep(100 * time.Millisecond)
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

	difference := duration - to
	expectedDifference := time.Microsecond * 20
	if difference > expectedDifference {
		t.Errorf("NoBackoff() expected delay: %d, actual: %d", difference, expectedDifference)
	}
}
