package convert

import "testing"

func TestToPtrReturnsPointerToInt(t *testing.T) {
	value := 123
	ptr := ToPtr(value)
	if *ptr != value {
		t.Errorf("Expected %v, got %v", value, *ptr)
	}
}

func TestToPtrReturnsPointerToString(t *testing.T) {
	value := "hello"
	ptr := ToPtr(value)
	if *ptr != value {
		t.Errorf("Expected %v, got %v", value, *ptr)
	}
}

func TestToPtrReturnsPointerToBool(t *testing.T) {
	value := true
	ptr := ToPtr(value)
	if *ptr != value {
		t.Errorf("Expected %v, got %v", value, *ptr)
	}
}

func TestToPtrReturnsPointerToFloat(t *testing.T) {
	value := 1.23
	ptr := ToPtr(value)
	if *ptr != value {
		t.Errorf("Expected %v, got %v", value, *ptr)
	}
}
