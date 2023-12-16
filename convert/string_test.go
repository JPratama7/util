package convert

import "testing"

func TestUnsafeStringConvertsBytesToString(t *testing.T) {
	b := []byte("hello")
	expected := "hello"
	result := UnsafeString(b)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSafeStringConvertsBytesToString(t *testing.T) {
	b := []byte("hello")
	expected := "hello"
	result := SafeString(b)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestUnsafeStringConvertsEmptyBytesToString(t *testing.T) {
	b := []byte("")
	expected := ""
	result := UnsafeString(b)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSafeStringConvertsEmptyBytesToString(t *testing.T) {
	b := []byte("")
	expected := ""
	result := SafeString(b)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
