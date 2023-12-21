package sync

import "sync"

type Pool[T any] struct {
	m *sync.Pool
}

func NewPool[T any](f func() T) *Pool[T] {
	return &Pool[T]{
		m: &sync.Pool{
			New: func() any { return f() },
		},
	}
}

func (p *Pool[T]) Get() T {
	return p.m.Get().(T)
}

func (p *Pool[T]) Put(x T) {
	p.m.Put(x)
}
