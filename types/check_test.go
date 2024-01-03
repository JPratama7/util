package types

import (
	"testing"
)

func TestIsPointerReturnsTrueForPointer(t *testing.T) {
	result := IsPointer(&struct{}{})
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestIsPointerReturnsFalseForNonPointer(t *testing.T) {
	result := IsPointer(struct{}{})
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestIsPointerReturnsFalseForNil(t *testing.T) {
	result := IsPointer(nil)
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestIsPointerReturnsFalseForPrimitive(t *testing.T) {
	result := IsPointer(5)
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
