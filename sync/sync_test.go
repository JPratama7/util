package sync

import "testing"

func TestNewPoolReturnsNewInstance(t *testing.T) {
	pool := NewPool(func() int { return 1 })
	if pool == nil {
		t.Errorf("Expected pool to be not nil")
	}
}

func TestGetReturnsValueFromPool(t *testing.T) {
	pool := NewPool(func() int { return 1 })
	value := pool.Get()
	if value != 1 {
		t.Errorf("Expected 1, got %v", value)
	}
}

func TestPutAddsValueToPool(t *testing.T) {
	pool := NewPool(func() int { return 0 })
	pool.Put(2)
	value := pool.Get()
	if value != 2 {
		t.Errorf("Expected 2, got %v", value)
	}
}

func TestGetReturnsNewValueIfPoolIsEmpty(t *testing.T) {
	pool := NewPool(func() int { return 1 })
	pool.Put(2)
	pool.Get() // This will empty the pool
	value := pool.Get()
	if value != 1 {
		t.Errorf("Expected 1, got %v", value)
	}
}

func TestGetRetrievesValueFromPool(t *testing.T) {
	pool := NewPool(func() int { return 1 })
	pool.Put(2)
	value := pool.Get()
	if value != 2 {
		t.Errorf("Expected 2, got %v", value)
	}
}

func TestGetRetrievesNewValueIfPoolIsEmpty(t *testing.T) {
	pool := NewPool(func() int { return 1 })
	pool.Put(2)
	pool.Get() // This will empty the pool
	value := pool.Get()
	if value != 1 {
		t.Errorf("Expected 1, got %v", value)
	}
}

func TestGetRetrievesSameValueIfPoolIsNotEmpty(t *testing.T) {
	pool := NewPool(func() int { return 1 })
	pool.Put(2)
	value1 := pool.Get()
	pool.Put(value1)
	value2 := pool.Get()
	if value1 != value2 {
		t.Errorf("Expected %v, got %v", value1, value2)
	}
}
