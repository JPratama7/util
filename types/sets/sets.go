package sets

import "sync"

type empty = struct{}

type Sets[T comparable] struct {
	m  map[T]empty
	sy sync.RWMutex
}

func NewSets[T comparable]() *Sets[T] {
	return &Sets[T]{m: make(map[T]empty)}
}

func NewFromSlices[T comparable](slices []T) *Sets[T] {
	set := NewSets[T]()
	for _, slice := range slices {
		set.Add(slice)
	}
	return set
}

func NewFromMapKey[T comparable](m map[T]any) *Sets[T] {
	set := NewSets[T]()
	for key := range m {
		set.Add(key)
	}

	return set
}

func NewFromMapValue[T comparable](m map[any]T) *Sets[T] {
	set := NewSets[T]()
	for _, value := range m {
		set.Add(value)
	}
	return set
}

func (s *Sets[T]) Add(value T) {
	s.sy.Lock()
	defer s.sy.Unlock()

	s.m[value] = empty{}
}

func (s *Sets[T]) Remove(value T) {
	s.sy.Lock()
	defer s.sy.Unlock()

	delete(s.m, value)
}

func (s *Sets[T]) Contains(value T) bool {
	s.sy.RLock()
	defer s.sy.RUnlock()

	_, ok := s.m[value]
	return ok
}

func (s *Sets[T]) Size() int {
	s.sy.RLock()
	defer s.sy.RUnlock()

	return len(s.m)
}

func (s *Sets[T]) ToSlice() []T {
	s.sy.RLock()
	defer s.sy.RUnlock()

	slice := make([]T, 0, len(s.m))
	for key := range s.m {
		slice = append(slice, key)
	}
	return slice
}

func (s *Sets[T]) Clear() {
	s.sy.Lock()
	defer s.sy.Unlock()

	s.m = make(map[T]empty)
}

func (s *Sets[T]) Union(other *Sets[T]) *Sets[T] {
	s.sy.RLock()
	defer s.sy.RUnlock()

	other.sy.RLock()
	defer other.sy.RUnlock()

	union := NewSets[T]()
	for key := range s.m {
		union.Add(key)
	}
	for key := range other.m {
		union.Add(key)
	}
	return union
}

func (s *Sets[T]) Intersection(other *Sets[T]) *Sets[T] {
	s.sy.RLock()
	defer s.sy.RUnlock()

	other.sy.RLock()
	defer other.sy.RUnlock()

	intersection := NewSets[T]()
	for key := range s.m {
		if other.Contains(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

func (s *Sets[T]) Difference(other *Sets[T]) *Sets[T] {
	s.sy.RLock()
	defer s.sy.RUnlock()

	other.sy.RLock()
	defer other.sy.RUnlock()

	difference := NewSets[T]()
	for key := range s.m {
		if !other.Contains(key) {
			difference.Add(key)
		}
	}
	return difference
}

func (s *Sets[T]) IsSubset(other *Sets[T]) bool {
	s.sy.RLock()
	defer s.sy.RUnlock()

	other.sy.RLock()
	defer other.sy.RUnlock()

	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (s *Sets[T]) Next() (c chan T) {
	c = make(chan T, len(s.m))

	go func(m map[T]empty, mu *sync.RWMutex) {
		s.sy.RLock()
		defer s.sy.RUnlock()
		for key := range m {
			c <- key
		}
		close(c)
	}(s.m, &s.sy)

	return
}

func (s *Sets[T]) Range(f func(v T) bool) {
	s.sy.RLock()
	defer s.sy.RUnlock()

	for key := range s.m {
		if !f(key) {
			break
		}
	}
}
