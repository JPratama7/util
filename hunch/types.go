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
type IndexedValue struct {
	Index int
	Value any
}

// IndexedExecutableOutput stores both output and error values from a Excetable.
type IndexedExecutableOutput struct {
	Value IndexedValue
	Err   error
}
