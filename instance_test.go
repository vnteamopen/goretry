package goretry

import (
	"bytes"
	"testing"
	"time"
)

func TestInstance_logWithoutLogger(t *testing.T) {
	withoutLogInstance := Instance{
		Logger: nil,
	}
	withoutLogInstance.log("Test %s", "no logger") // without error or panic

}

func TestInstance_logWithLogger(t *testing.T) {
	var buffer bytes.Buffer
	instanceWithLogger := Instance{
		Logger: &buffer,
	}
	instanceWithLogger.log("Test %s", "logger")
	expected := "Test logger\n"
	if buffer.String() != expected {
		t.Errorf("Expected: %v, got: %v", expected, buffer.String())
	}
}

func TestInstance_sleepWithMaxValue(t *testing.T) {
	var buffer bytes.Buffer
	instanceWithLogger := Instance{
		CeilingSleep: time.Duration(100 * time.Millisecond),
		Logger:       &buffer,
	}
	instanceWithLogger.sleep(200 * time.Millisecond)
	expected := "sleep 100ms\n"
	if buffer.String() != expected {
		t.Errorf("Expected: %v, got: %v", expected, buffer.String())
	}
}

func TestInstance_sleepWithoutMaxValue(t *testing.T) {
	var buffer bytes.Buffer
	instanceWithLogger := Instance{
		Logger: &buffer,
	}
	instanceWithLogger.sleep(200 * time.Millisecond)
	expected := "sleep 200ms\n"
	if buffer.String() != expected {
		t.Errorf("Expected: %v, got: %v", expected, buffer.String())
	}
}

func TestInstance_sleepWithJitter(t *testing.T) {
	var buffer bytes.Buffer
	instanceWithLogger := Instance{
		Logger:           &buffer,
		JitterEnabled:    true,
		JitterFloorSleep: 10 * time.Millisecond,
	}
	instanceWithLogger.sleep(200 * time.Millisecond)
	notExpected := "sleep 200ms\n"
	if buffer.String() == notExpected {
		t.Errorf("Not expected but got: %v", buffer.String())
	}
}
