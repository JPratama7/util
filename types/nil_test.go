package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilStringReturnsNil(t *testing.T) {
	result := NilString()
	assert.Nil(t, result, "Expected nil, got %v", result)

}

func TestNilBoolReturnsNil(t *testing.T) {
	result := NilBool()
	assert.Nil(t, result, "Expected nil, got %v", result)
}

func TestNilIntReturnsNil(t *testing.T) {
	result := NilInt()
	assert.Nil(t, result, "Expected nil, got %v", result)

}

func TestNilStructReturnsNil(t *testing.T) {
	result := NilStruct()
	assert.Nil(t, result, "Expected nil, got %v", result)
}
