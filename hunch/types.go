package hunch

import (
	"context"
)

// Executable represents a singular logic block.
// It can be used with several functions.
type Executable[T any] func(context.Context) (T, error)

// ExecutableInSequence represents one of a sequence of logic blocks.
type ExecutableInSequence[T any] func(context.Context, T) (T, error)

// IndexedValue stores the output of Executables,
// along with the index of the source Executable for ordering.
type IndexedValue[T any] struct {
	Index int
	Value T
}

// IndexedExecutableOutput stores both output and error values from a Excetable.
type IndexedExecutableOutput[T any] struct {
	Value IndexedValue[T]
	Err   error
}
