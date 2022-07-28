package goretry_test

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/vnteamopen/goretry"
)

func TestFibonacci(t *testing.T) {
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
				if (*counter) >= 6 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 6,
			expectedLog: `do action()
sleep 10ms
do action()
sleep 10ms
do action()
sleep 20ms
do action()
sleep 30ms
do action()
sleep 50ms
do action()
`,
		},
		{
			name: "MaxStopRetries",
			instance: goretry.Instance{
				MaxStopRetries: 4,
			},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 6 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 4,
			expectedLog: `do action()
sleep 10ms
do action()
sleep 10ms
do action()
sleep 20ms
do action()
`,
		},
		{
			name: "MaxStopTotalWaiting",
			instance: goretry.Instance{
				MaxStopTotalWaiting: time.Duration(30 * time.Millisecond),
			},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 6 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 4,
			expectedLog: `do action()
sleep 10ms
do action()
sleep 10ms
do action()
sleep 20ms
do action()
`,
		},
		{
			name: "MaxWaiting",
			instance: goretry.Instance{
				CeilingSleep: time.Duration(12 * time.Millisecond),
			},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 6 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 6,
			expectedLog: `do action()
sleep 10ms
do action()
sleep 10ms
do action()
sleep 12ms
do action()
sleep 12ms
do action()
sleep 12ms
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

			instance.Fibonacci(10*time.Millisecond, func() error {
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
