package tester

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, expected, actual any) {
	if reflect.DeepEqual(expected, actual) {
		t.Error("not equal")
	}
}

func NotEqual(t *testing.T, expected, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Error("not equal")
	}
}
