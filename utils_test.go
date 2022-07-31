package goretry_test

import (
	"testing"

	"github.com/vnteamopen/goretry"
)

func TestIntPow(t *testing.T) {
	t.Run("Test power calculation", func(t *testing.T) {
		result := goretry.IntPow(2, 3)
		expected := int64(8)
		if result != expected {
			t.Errorf("Expected: %v, got: %v", expected, result)
		}
	})
}
