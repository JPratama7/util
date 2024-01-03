package paseto

import (
	"testing"
	"time"
)

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	if len(privateKey) == 0 || len(publicKey) == 0 {
		t.Errorf("Expected non-empty keys, got privateKey: %v, publicKey: %v", privateKey, publicKey)
	}
}

func TestNewPASETO(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	paseto := NewPASETO(publicKey, privateKey)
	if paseto.Private != privateKey || paseto.Public != publicKey || paseto.Duration != 2*time.Hour {
		t.Errorf("Expected PASETO with privateKey: %v, publicKey: %v, duration: %v, got %v", privateKey, publicKey, 2*time.Hour, paseto)
	}
}

func TestNewPASETOWithCustomDuration(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	paseto := NewPASETO(publicKey, privateKey, 1*time.Hour)
	if paseto.Private != privateKey || paseto.Public != publicKey || paseto.Duration != 1*time.Hour {
		t.Errorf("Expected PASETO with privateKey: %v, publicKey: %v, duration: %v, got %v", privateKey, publicKey, 1*time.Hour, paseto)
	}
}
