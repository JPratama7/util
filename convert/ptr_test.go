package convert

import "testing"

func TestToPtrReturnsPointerToInt(t *testing.T) {
	value := 123
	ptr := ToPointer(value)
	if *ptr != value {
		t.Errorf("Expected %v, got %v", value, *ptr)
	}
}

func TestToPtrReturnsPointerToString(t *testing.T) {
	value := "hello"
	ptr := ToPointer(value)
	if *ptr != value {
		t.Errorf("Expected %v, got %v", value, *ptr)
	}
}

func TestToPtrReturnsPointerToBool(t *testing.T) {
	value := true
	ptr := ToPointer(value)
	if *ptr != value {
		t.Errorf("Expected %v, got %v", value, *ptr)
	}
}

func TestToPtrReturnsPointerToFloat(t *testing.T) {
	value := 1.23
	ptr := ToPointer(value)
	if *ptr != value {
		t.Errorf("Expected %v, got %v", value, *ptr)
	}
}
