package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want %v", actual, expected)
	}
}

func StringContains(t *testing.T, actual, expectedSubstring string) {
	t.Helper()

	if !string.Contains(actual, expectedSubstring) {
		t.Errorf("got: %q; expected to contain %q", actual, expectedSubstring)
	}
}
