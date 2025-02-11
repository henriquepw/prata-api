package assert

import (
	"reflect"
	"testing"
	"time"
)

func LenEqual[V any](t *testing.T, slice []V, expectedLen int) {
	t.Helper()

	actualLen := len(slice)

	if expectedLen != actualLen {
		t.Errorf("expected that the slice had size %v, but actual size is: %v", expectedLen, actualLen)
	}
}

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

func Nil(t *testing.T, value any) {
	t.Helper()

	if err, ok := value.(error); ok {
		if err != nil {
			t.Errorf("error should be nil, but is: %v", err.Error())
		}
	}

	if value != nil && reflect.ValueOf(value).IsValid() && !reflect.ValueOf(value).IsNil() {
		t.Errorf("Expected the value to be nil, but is: %v", value)
	}
}

func NotNil(t *testing.T, value any) {
	t.Helper()

	if err, ok := value.(error); ok {
		if err == nil {
			t.Errorf("error should not be nil, but is")
		}
	}

	if value == nil {
		t.Errorf("Expected the value not to be nil, but is")
	}
}

func False(t *testing.T, v bool) {
	t.Helper()

	if v {
		t.Errorf("Expected value to be %t but it's %t", false, v)
	}
}

func True(t *testing.T, v bool) {
	t.Helper()

	if !v {
		t.Errorf("Expected value to be %t but it's %t", true, v)
	}
}

func EmptyString[T ~string](t *testing.T, s T) {
	t.Helper()

	if s != "" {
		t.Errorf("Expected value to be empty, but was %v", s)
	}
}

func NotEmptyString[T ~string](t *testing.T, s T) {
	t.Helper()

	if s == "" {
		t.Errorf("Expected value to not be empty, but was")
	}
}

func TimeIsZero(t *testing.T, ti time.Time) {
	t.Helper()

	if !ti.IsZero() {
		t.Errorf("Expected the time value to be zero, but was %v", ti)
	}
}

func TimeIsNotZero(t *testing.T, ti time.Time) {
	t.Helper()

	if ti.IsZero() {
		t.Errorf("Expected the time value to be zero, but was %v", ti)
	}
}
