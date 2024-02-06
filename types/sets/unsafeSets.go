package sets

type UnsafeSets[T comparable] struct {
	m map[T]empty
}

func NewUnsafeSets[T comparable]() *UnsafeSets[T] {
	return &UnsafeSets[T]{m: make(map[T]empty)}
}

func NewUnsafeFromSlices[T comparable](slices []T) *UnsafeSets[T] {
	set := NewUnsafeSets[T]()
	for _, slice := range slices {
		set.Add(slice)
	}
	return set
}

func NewUnsafeFromMapKey[T comparable](m map[T]any) *UnsafeSets[T] {
	set := NewUnsafeSets[T]()
	for key := range m {
		set.Add(key)
	}

	return set
}

func NewUnsafeFromMapValue[T comparable](m map[any]T) *UnsafeSets[T] {
	set := NewUnsafeSets[T]()
	for _, value := range m {
		set.Add(value)
	}
	return set
}

func (s *UnsafeSets[T]) Add(value T) {
	s.m[value] = empty{}
}

func (s *UnsafeSets[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *UnsafeSets[T]) Contains(value T) bool {
	_, ok := s.m[value]
	return ok
}

func (s *UnsafeSets[T]) Size() int {
	return len(s.m)
}

func (s *UnsafeSets[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.m))
	for key := range s.m {
		slice = append(slice, key)
	}
	return slice
}

func (s *UnsafeSets[T]) Clear() {
	s.m = make(map[T]empty)
}

func (s *UnsafeSets[T]) Union(other *UnsafeSets[T]) *UnsafeSets[T] {
	union := NewUnsafeSets[T]()
	for key := range s.m {
		union.Add(key)
	}
	for key := range other.m {
		union.Add(key)
	}
	return union
}

func (s *UnsafeSets[T]) Intersection(other *UnsafeSets[T]) *UnsafeSets[T] {
	intersection := NewUnsafeSets[T]()
	for key := range s.m {
		if other.Contains(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

func (s *UnsafeSets[T]) Difference(other *UnsafeSets[T]) *UnsafeSets[T] {
	difference := NewUnsafeSets[T]()
	for key := range s.m {
		if !other.Contains(key) {
			difference.Add(key)
		}
	}
	return difference
}

func (s *UnsafeSets[T]) IsSubset(other *UnsafeSets[T]) bool {
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (s *UnsafeSets[T]) Next() (c chan T) {
	c = make(chan T, len(s.m))

	go func(m map[T]empty) {
		for key := range m {
			c <- key
		}
		close(c)
	}(s.m)

	return
}

func (s *UnsafeSets[T]) Range(f func(v T) bool) {
	for key := range s.m {
		if !f(key) {
			break
		}
	}
}

func (s *UnsafeSets[T]) Iterator() {
}
