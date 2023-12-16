package convert

import (
	"reflect"
	"testing"
	"time"
)

func TestCopyStringReturnsCopy(t *testing.T) {
	s := "hello"
	result := CopyString(s)
	if result != s {
		t.Errorf("Expected %v, got %v", s, result)
	}
}

func TestCopyBytesReturnsCopy(t *testing.T) {
	b := []byte("hello")
	result := CopyBytes(b)
	if !reflect.DeepEqual(result, b) {
		t.Errorf("Expected %v, got %v", b, result)
	}
}

func TestToStringConvertsInt(t *testing.T) {
	result := ToString(123)
	if result != "123" {
		t.Errorf("Expected 123, got %v", result)
	}
}

func TestToStringConvertsString(t *testing.T) {
	result := ToString("hello")
	if result != "hello" {
		t.Errorf("Expected hello, got %v", result)
	}
}

func TestToStringConvertsBool(t *testing.T) {
	result := ToString(true)
	if result != "true" {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestToStringConvertsFloat(t *testing.T) {
	result := ToString(1.23)
	if result != "1.23" {
		t.Errorf("Expected 1.23, got %v", result)
	}
}

func TestToStringConvertsTime(t *testing.T) {
	timeNow := time.Now()
	result := ToString(timeNow)
	if result != timeNow.Format("2006-01-02 15:04:05") {
		t.Errorf("Expected %v, got %v", timeNow.Format("2006-01-02 15:04:05"), result)
	}
}

func TestToStringReturnsEmptyForUnknown(t *testing.T) {
	result := ToString(struct{}{})
	if result != "" {
		t.Errorf("Expected empty string, got %v", result)
	}
}
