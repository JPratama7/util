package paseto

import (
	"testing"
	"time"
)

var priv, pub = GenerateKey()

func TestEncode(t *testing.T) {
	_, err := Encode("id", priv)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestEncodeWithInvalidPrivateKey(t *testing.T) {
	_, err := Encode("id", "")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestEncodeWithStruct(t *testing.T) {
	data := struct {
		Name string
		Age  int
	}{"John", 30}
	_, err := EncodeWithStruct("id", &data, priv)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestEncodeWithStructDuration(t *testing.T) {
	data := struct {
		Name string
		Age  int
	}{"John", 30}
	_, err := EncodeWithStructDuration("id", &data, priv, 1*time.Hour)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestEncodeforHours(t *testing.T) {
	_, err := EncodeforHours("id", priv, 1)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestEncodeforMinutes(t *testing.T) {
	_, err := EncodeforMinutes("id", priv, 60)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestEncodeforSeconds(t *testing.T) {
	_, err := EncodeforSeconds("id", priv, 3600)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestDecode(t *testing.T) {
	token, _ := Encode("id", priv)
	_, err := Decode[any](pub, token)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestDecodeWithStruct(t *testing.T) {
	type d struct {
		Name string
		Age  int
	}
	data := d{"John", 30}
	token, _ := EncodeWithStruct("id", &data, priv)
	_, err := DecodeWithStruct[d](pub, token)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestDecodeGetId(t *testing.T) {
	token, err := Encode("id", priv)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	id := DecodeGetId(pub, token)
	if id != "id" {
		t.Errorf("Expected 'id', got %v", id)
	}
}
