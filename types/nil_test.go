package types

import "testing"

func TestNilStringReturnsNil(t *testing.T) {
	result := NilString()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestNilBoolReturnsNil(t *testing.T) {
	result := NilBool()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestNilIntReturnsNil(t *testing.T) {
	result := NilInt()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestNilStructReturnsNil(t *testing.T) {
	result := NilStruct()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
