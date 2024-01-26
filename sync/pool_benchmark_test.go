package sync

import (
	"testing"
)

func BenchmarkNewPoolInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPool(func() int { return 1 })
	}
}

func BenchmarkGetInt(b *testing.B) {
	pool := NewPool(func() int { return 1 })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkPutInt(b *testing.B) {
	pool := NewPool(func() int { return 1 })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Put(1)
	}
}

func BenchmarkNewPoolSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPool(func() []string { return make([]string, 0) })
	}
}

func BenchmarkGetSlice(b *testing.B) {
	pool := NewPool(func() []string { return make([]string, 0) })
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		pool.Get()
	}
}

func BenchmarkPutSlice(b *testing.B) {
	pool := NewPool(func() []string { return make([]string, 0) })
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		data := pool.Get()

		data = append(data, "hello")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")
		data = append(data, "world")

		pool.Put(data)
	}
}
