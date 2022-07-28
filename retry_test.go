package goretry_test

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/vnteamopen/goretry"
)

func TestDo(t *testing.T) {
	testCases := []struct {
		name            string
		instance        goretry.Instance
		action          func(counter *int64) error
		expectedCounter int64
		expectedLog     string
	}{
		{
			name:     "default",
			instance: goretry.Instance{},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 3 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 3,
			expectedLog: `do action()
sleep 10ms
do action()
sleep 10ms
do action()
`,
		},
		{
			name: "MaxStopRetries",
			instance: goretry.Instance{
				MaxStopRetries: 2,
			},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 3 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 2,
			expectedLog: `do action()
sleep 10ms
do action()
`,
		},
		{
			name: "MaxStopTotalWaiting",
			instance: goretry.Instance{
				MaxStopTotalWaiting: time.Duration(10 * time.Millisecond),
			},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 3 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 2,
			expectedLog: `do action()
sleep 10ms
do action()
`,
		},
		{
			name: "MaxWaiting",
			instance: goretry.Instance{
				CeilingSleep: time.Duration(5 * time.Millisecond),
			},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 3 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 3,
			expectedLog: `do action()
sleep 5ms
do action()
sleep 5ms
do action()
`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var counter int64
			var buffer bytes.Buffer
			instance := goretry.Instance{
				MaxStopRetries:      testCase.instance.MaxStopRetries,
				MaxStopTotalWaiting: testCase.instance.MaxStopTotalWaiting,
				CeilingSleep:        testCase.instance.CeilingSleep,
				Logger:              &buffer,
			}

			instance.Do(10*time.Millisecond, func() error {
				return testCase.action(&counter)
			})

			if counter != testCase.expectedCounter {
				t.Errorf("Do() expected counting: %d, actual: %d", testCase.expectedCounter, counter)
			}
			if buffer.String() != testCase.expectedLog {
				t.Errorf("Expected: %v, got: %v", testCase.expectedLog, buffer.String())
			}
		})
	}
}

func TestNoBackoff(t *testing.T) {
	testCases := []struct {
		name            string
		instance        goretry.Instance
		action          func(counter *int64) error
		expectedCounter int64
		expectedLog     string
	}{
		{
			name:     "default",
			instance: goretry.Instance{},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 3 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 3,
			expectedLog: `do action()
sleep 0s
do action()
sleep 0s
do action()
`,
		},
		{
			name: "MaxStopRetries",
			instance: goretry.Instance{
				MaxStopRetries: 2,
			},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 3 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 2,
			expectedLog: `do action()
sleep 0s
do action()
`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var counter int64
			var buffer bytes.Buffer
			instance := goretry.Instance{
				MaxStopRetries: testCase.instance.MaxStopRetries,
				Logger:         &buffer,
			}

			instance.NoBackoff(func() error {
				return testCase.action(&counter)
			})

			if counter != testCase.expectedCounter {
				t.Errorf("Do() expected counting: %d, actual: %d", testCase.expectedCounter, counter)
			}
			if buffer.String() != testCase.expectedLog {
				t.Errorf("Expected: %v, got: %v", testCase.expectedLog, buffer.String())
			}
		})
	}
}
