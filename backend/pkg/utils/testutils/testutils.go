package testutils

import (
	"reflect"
	"testing"
)

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected respone code %d. Got %d\n", expected, actual)
	}
}

func CheckEqual[K any](t *testing.T, expected, actual K) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v. Got %v\n", expected, actual)
	}
}
