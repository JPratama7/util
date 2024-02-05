package types

import (
	"github.com/stretchr/testify/assert"
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

func TestIsIterableReturnsTrueForSlice(t *testing.T) {
	result := IsIterable([]int{1, 2, 3})
	assert.True(t, result)
}

func TestIsIterableReturnsTrueForArray(t *testing.T) {
	result := IsIterable([3]int{1, 2, 3})
	assert.True(t, result)
}

func TestIsIterableReturnsTrueForMap(t *testing.T) {
	result := IsIterable(map[string]int{"one": 1, "two": 2})
	assert.True(t, result)
}

func TestIsIterableReturnsFalseForNonIterable(t *testing.T) {
	result := IsIterable(5)
	assert.False(t, result)
}

func TestIsIterableReturnsFalseForNil(t *testing.T) {
	result := IsIterable(nil)
	assert.False(t, result)
}
