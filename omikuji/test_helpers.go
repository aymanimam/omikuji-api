package omikuji

import (
	"testing"
)

// Contains Is the given array has the given omikuji
func Contains(arr []Omikuji, o Omikuji) bool {
	for _, e := range arr {
		if e.Text == o.Text {
			return true
		}
	}
	return false
}

// AssertPanicFunc Function under test
type AssertPanicFunc func()

// AssertPanic assert that the given function should panic
func AssertPanic(t *testing.T, failMessage string, function AssertPanicFunc) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf(failMessage)
			}
		}()
		// This function should cause a panic
		function()
	}()
}
