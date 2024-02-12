package generator

import "context"

// IteratorMap is a type that can be used to generate values of type T.
type IteratorMap[K comparable, V any, M ~map[K]V] struct {
	iter    M
	lastKey K
}

func NewIterMap[K comparable, V any, M map[K]V](iter M) *IteratorMap[K, V, M] {
	return &IteratorMap[K, V, M]{
		iter:    iter,
		lastKey: *new(K),
	}
}

// HasNext returns true if there are more values to generate.
func (i *IteratorMap[K, V, M]) HasNext() bool {
	if len(i.iter) == 0 {
		return false
	}

	if i.lastKey == *new(K) {
		return true
	}

	for key := range i.iter {
		if key != i.lastKey {
			return true
		}
	}

	return false
}

// Next returns true if there are more values to generate.
func (i *IteratorMap[K, V, M]) Next() bool {
	if len(i.iter) == 0 {
		return false
	}

	if i.lastKey == *new(K) {
		for key := range i.iter {
			i.lastKey = key
			return true
		}
	}

	delete(i.iter, i.lastKey)
	for key := range i.iter {
		i.lastKey = key
		return true
	}

	return false
}

// Value returns the current value.
func (i *IteratorMap[K, V, M]) Value() (K, V) {
	return i.lastKey, i.iter[i.lastKey]
}

func (i *IteratorMap[K, V, M]) Chan(c context.Context) (chan K, chan V) {
	return ToChannelMap(c, i.iter)
}
