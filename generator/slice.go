package generator

import "context"

// IteratorSlice is a type that can be used to generate values of type T.
type IteratorSlice[K int, T any] struct {
	iter  []T
	index int
}

func NewIterSlice[T any, V []T](iter V) *IteratorSlice[int, T] {
	return &IteratorSlice[int, T]{
		iter:  iter,
		index: -1,
	}
}

// Next returns true if there are more values to generate.
func (i *IteratorSlice[K, T]) Next() bool {
	i.index++
	return i.index < len(i.iter)
}

// Value returns the current value.
func (i *IteratorSlice[K, T]) Value() (int, T) {
	if i.index < 0 || i.index >= len(i.iter) {
		return -1, *new(T)
	}
	return i.index, i.iter[i.index]
}

// HasNext returns true if there are more values to generate.
func (i *IteratorSlice[K, T]) HasNext() bool {
	return i.index+1 < len(i.iter)
}

func (i *IteratorSlice[K, T]) Chan(ctx context.Context) (chan int, chan T) {
	return ToChannelSlice(ctx, i.iter)
}
