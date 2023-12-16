package convert

import (
	"reflect"
	"testing"
)

func TestUnsafeBytesConversion(t *testing.T) {
	s := "hello"
	expected := []byte(s)
	result := UnsafeBytes(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSafeBytesConversion(t *testing.T) {
	s := "hello"
	expected := []byte(s)
	result := SafeBytes(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestUnsafeBytesConversionWithEmptyString(t *testing.T) {
	s := ""
	expected := []byte(s)
	result := UnsafeBytes(s)
	if len(s) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSafeBytesConversionWithEmptyString(t *testing.T) {
	s := ""
	expected := []byte(s)
	result := SafeBytes(s)
	if len(s) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
