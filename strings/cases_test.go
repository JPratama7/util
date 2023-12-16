package strings

import (
	"testing"
)

func TestLowerCaseConversion(t *testing.T) {
	result := ToLower("HELLO")
	if result != "hello" {
		t.Errorf("Expected hello, got %v", result)
	}
}

func TestUpperCaseConversion(t *testing.T) {
	result := ToUpper("hello")
	if result != "HELLO" {
		t.Errorf("Expected HELLO, got %v", result)
	}
}

func TestLowerCaseConversionWithNumbers(t *testing.T) {
	result := ToLower("HELLO123")
	if result != "hello123" {
		t.Errorf("Expected hello123, got %v", result)
	}
}

func TestUpperCaseConversionWithNumbers(t *testing.T) {
	result := ToUpper("hello123")
	if result != "HELLO123" {
		t.Errorf("Expected HELLO123, got %v", result)
	}
}

func TestLowerCaseConversionWithSpecialCharacters(t *testing.T) {
	result := ToLower("HELLO@#")
	if result != "hello@#" {
		t.Errorf("Expected hello@#, got %v", result)
	}
}

func TestUpperCaseConversionWithSpecialCharacters(t *testing.T) {
	result := ToUpper("hello@#")
	if result != "HELLO@#" {
		t.Errorf("Expected HELLO@#, got %v", result)
	}
}
