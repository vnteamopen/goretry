package goretry_test

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/vnteamopen/goretry"
)

func TestExponential(t *testing.T) {
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
				if (*counter) >= 4 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 4,
			expectedLog: `do action()
sleep 10ms
do action()
sleep 40ms
do action()
sleep 80ms
do action()
`,
		},
		{
			name: "MaxStopRetries",
			instance: goretry.Instance{
				MaxStopRetries: 3,
			},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 4 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 3,
			expectedLog: `do action()
sleep 10ms
do action()
sleep 40ms
do action()
`,
		},
		{
			name: "MaxStopTotalWaiting",
			instance: goretry.Instance{
				MaxStopTotalWaiting: time.Duration(50 * time.Millisecond),
			},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 4 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 3,
			expectedLog: `do action()
sleep 10ms
do action()
sleep 40ms
do action()
`,
		},
		{
			name: "MaxWaiting",
			instance: goretry.Instance{
				CeilingSleep: time.Duration(42 * time.Millisecond),
			},
			action: func(counter *int64) error {
				(*counter)++
				if (*counter) >= 4 {
					return nil
				}
				return errors.New("fake error")
			},
			expectedCounter: 4,
			expectedLog: `do action()
sleep 10ms
do action()
sleep 40ms
do action()
sleep 42ms
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

			instance.Exponential(10*time.Millisecond, func() error {
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
