package test

import (
	"runtime/debug"
	"testing"
)

// AssertEquals asserts if the two objects are the same.
func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		debug.PrintStack()
		t.Errorf("Expected: %s, got: %s\n", expected, actual)
	}
}

// Fail fails the test.
func Fail(t *testing.T, message ...interface{}) {
	debug.PrintStack()
	t.Fatal(message)
}
