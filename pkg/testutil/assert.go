package testutil

import "testing"

func Equal[V comparable](t *testing.T, expected, actual V) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected %v, but got: %v", expected, actual)
	}
}

func NotEqual[V comparable](t *testing.T, expected, actual V) {
	t.Helper()

	if expected == actual {
		t.Errorf("expected %v should not be equal to actual: %v", expected, actual)
	}
}

func NilError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("error should be nil, but is: %v", err.Error())
	}
}

func NotNilError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("error should not be nil")
	}
}
